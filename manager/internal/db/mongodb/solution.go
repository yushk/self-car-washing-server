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

func (m *mongoDB) CreateSolution(ctx context.Context, data *pb.Solution) (*pb.Solution, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CSolution().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetSolution(ctx context.Context, id string) (data *pb.Solution, err error) {
	data = &pb.Solution{}
	filter := bson.M{"id": id}
	err = m.CSolution().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateSolution(ctx context.Context, id string, data *pb.Solution) (*pb.Solution, error) {
	newData := &pb.Solution{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CSolution().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteSolution(ctx context.Context, id string) (data *pb.Solution, err error) {
	data = &pb.Solution{}
	filter := bson.M{"id": id}
	err = m.CSolution().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetSolutions(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.Solution, err error) {
	datas = []*pb.Solution{}

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
	cursor, err := m.CSolution().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Solution{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Solution Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CSolution().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Solution Count Documents Error")
		count = int64(0)
	}
	return
}
