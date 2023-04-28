package mongodb

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreateComment(ctx context.Context, data *pb.Comment) (*pb.Comment, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CComment().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetComment(ctx context.Context, id string) (data *pb.Comment, err error) {
	data = &pb.Comment{}
	filter := bson.M{"id": id}
	err = m.CComment().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateComment(ctx context.Context, id string, data *pb.Comment) (*pb.Comment, error) {
	newData := &pb.Comment{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CComment().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteComment(ctx context.Context, id string) (data *pb.Comment, err error) {
	data = &pb.Comment{}
	filter := bson.M{"id": id}
	err = m.CComment().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetComments(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.Comment, err error) {
	datas = []*pb.Comment{}

	filter := bson.M{}
	if query != "" {
		err = json.Unmarshal([]byte(query), &filter)
		if err != nil {
			return
		}
	}

	findOption := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	cursor, err := m.CComment().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Comment{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Comment Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CComment().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Comment Count Documents Error")
		count = int64(0)
	}
	return
}
