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

func (m *mongoDB) CreateSportItem(ctx context.Context, data *pb.SportItem) (*pb.SportItem, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	if data.CreateTime == 0 {
		data.CreateTime = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	if data.UpdateTime == 0 {
		data.UpdateTime = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	_, err := m.CSportItem().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetSportItem(ctx context.Context, id string) (data *pb.SportItem, err error) {
	data = &pb.SportItem{}
	filter := bson.M{"id": id}
	err = m.CSportItem().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateSportItem(ctx context.Context, id string, data *pb.SportItem) (*pb.SportItem, error) {
	newData := &pb.SportItem{}
	filter := bson.M{"id": id}
	data.Id = id
	data.UpdateTime = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	update := bson.M{"$set": data}
	err := m.CSportItem().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteSportItem(ctx context.Context, id string) (data *pb.SportItem, err error) {
	data = &pb.SportItem{}
	filter := bson.M{"id": id}
	err = m.CSportItem().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetSportItems(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.SportItem, err error) {
	datas = []*pb.SportItem{}

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
	cursor, err := m.CSportItem().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.SportItem{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("SportItem Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CSportItem().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("SportItem Count Documents Error")
		count = int64(0)
	}
	return
}
