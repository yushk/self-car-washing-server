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

func (m *mongoDB) CreateHomeWork(ctx context.Context, data *pb.HomeWork) (*pb.HomeWork, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	data.SolutionsInfoId = []string{}
	data.SolutionsInfo = []*pb.Solution{}
	_, err := m.CHomeWork().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetHomeWork(ctx context.Context, id string) (data *pb.HomeWork, err error) {
	data = &pb.HomeWork{}
	filter := bson.M{"id": id}
	err = m.CHomeWork().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) AddHomeWorkSolution(ctx context.Context, id string, data *pb.Solution) error {
	filter := bson.M{"id": id}
	update := bson.M{"$addToSet": bson.M{"solutions_info_id": data.Id, "solutions_info": bson.M{"$each": []*pb.Solution{data}}}}
	// update := bson.M{"$addToSet": bson.M{"solutionsinfoid": data.Id, "solutionsinfo": bson.M{"$each": []*pb.Solution{data}}}}
	_, err := m.CHomeWork().UpdateOne(context.TODO(), filter, update)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}

func (m *mongoDB) UpdateHomeWorkSolution(ctx context.Context, id string, data *pb.Solution) error {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"solutions_info.$[elem]": data,
			// "solutionsinfo.$[elem]": data,
		},
	}
	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.id": data.Id},
		},
	})
	_, err := m.CHomeWork().UpdateOne(context.TODO(), filter, update, updateOptions)
	if err != nil {
		logrus.Errorln(err)
		return err
	}
	return nil
}

func (m *mongoDB) UpdateHomeWork(ctx context.Context, id string, data *pb.HomeWork) (*pb.HomeWork, error) {
	newData := &pb.HomeWork{}
	filter := bson.M{"id": id}
	data.Id = id
	update := bson.M{"$set": data}
	err := m.CHomeWork().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteHomeWork(ctx context.Context, id string) (data *pb.HomeWork, err error) {
	data = &pb.HomeWork{}
	filter := bson.M{"id": id}
	err = m.CHomeWork().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetHomeWorks(ctx context.Context, limit, skip int64, query string) (count int64, datas []*pb.HomeWork, err error) {
	datas = []*pb.HomeWork{}

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
	cursor, err := m.CHomeWork().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.HomeWork{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("HomeWork Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CHomeWork().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("HomeWork Count Documents Error")
		count = int64(0)
	}
	return
}
