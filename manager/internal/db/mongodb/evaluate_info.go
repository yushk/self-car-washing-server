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

func (m *mongoDB) CreateEvaluateInfo(ctx context.Context, data *pb.EvaluateInfo) (*pb.EvaluateInfo, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CEvaluateInfo().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) UpdateEvaluateInfo(ctx context.Context, id string, data *pb.EvaluateInfo) (*pb.EvaluateInfo, error) {
	newData := &pb.EvaluateInfo{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CEvaluateInfo().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteEvaluateInfo(ctx context.Context, id string) (data *pb.EvaluateInfo, err error) {
	data = &pb.EvaluateInfo{}
	filter := bson.M{"id": id}
	err = m.CEvaluateInfo().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetEvaluateInfos(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.EvaluateInfo, err error) {
	datas = []*pb.EvaluateInfo{}

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
	cursor, err := m.CEvaluateInfo().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.EvaluateInfo{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("EvaluateInfo Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CEvaluateInfo().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("EvaluateInfo Count Documents Error")
		count = int64(0)
	}
	return
}
