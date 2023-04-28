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

func (m *mongoDB) CreateSportType(ctx context.Context, data *pb.SportType) (*pb.SportType, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	_, err := m.CSportType().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetSportType(ctx context.Context, id string) (data *pb.SportType, err error) {
	data = &pb.SportType{}
	filter := bson.M{"id": id}
	err = m.CSportType().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateSportType(ctx context.Context, id string, data *pb.SportType) (*pb.SportType, error) {
	newData := &pb.SportType{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CSportType().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteSportType(ctx context.Context, id string) (data *pb.SportType, err error) {
	data = &pb.SportType{}
	filter := bson.M{"id": id}
	err = m.CSportType().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetSportTypes(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.SportType, err error) {
	datas = []*pb.SportType{}

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
	cursor, err := m.CSportType().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.SportType{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("SportType Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CSportType().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("SportType Count Documents Error")
		count = int64(0)
	}
	return
}
