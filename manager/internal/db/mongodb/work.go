package mongodb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoDB) CreateWork(ctx context.Context, data *pb.Work) (*pb.Work, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	if data.Create == 0 {
		data.Create = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	if data.Update == 0 {
		data.Update = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	_, err := m.CWork().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetWork(ctx context.Context, id string) (data *pb.Work, err error) {
	data = &pb.Work{}
	filter := bson.M{"id": id}
	err = m.CWork().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateWork(ctx context.Context, id string, data *pb.Work) (*pb.Work, error) {
	newData := &pb.Work{}
	filter := bson.M{"id": id}
	data.Id = id
	data.Update = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	update := bson.M{"$set": data}
	err := m.CWork().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteWork(ctx context.Context, id string) (data *pb.Work, err error) {
	data = &pb.Work{}
	filter := bson.M{"id": id}
	err = m.CWork().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetWorks(ctx context.Context, limit, skip int64, query, sort string) (count int64, datas []*pb.Work, err error) {
	datas = []*pb.Work{}

	filterSort := bson.M{}
	if sort != "" {
		err := json.Unmarshal([]byte(sort), &filterSort)
		if err != nil {
			return 0, datas, err
		}
	}

	findOption := &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  filterSort,
	}

	filter := bson.M{}
	if query != "" {
		err = json.Unmarshal([]byte(query), &filter)
		if err != nil {
			return
		}
	}

	cursor, err := m.CWork().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Work{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Work Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CWork().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Work Count Documents Error")
		count = int64(0)
	}
	return
}
