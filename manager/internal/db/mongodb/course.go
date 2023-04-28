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

func (m *mongoDB) CreateCourse(ctx context.Context, data *pb.Course) (*pb.Course, error) {
	if data.Id == "" {
		data.Id = xid.New().String()
	}
	if data.Create == 0 {
		data.Create = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	if data.Update == 0 {
		data.Update = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	}
	_, err := m.CCourse().InsertOne(ctx, data)
	if err != nil {
		logrus.Errorln(err)
		return data, err
	}
	return data, nil
}

func (m *mongoDB) GetCourse(ctx context.Context, id string) (data *pb.Course, err error) {
	data = &pb.Course{}
	filter := bson.M{"id": id}
	err = m.CCourse().FindOne(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) UpdateCourse(ctx context.Context, id string, data *pb.Course) (*pb.Course, error) {
	newData := &pb.Course{}
	filter := bson.M{"id": id}
	data.Id = id
	data.Update = time.Now().In(time.FixedZone("CST", 8*3600)).Unix()
	update := bson.M{"$set": data}
	err := m.CCourse().FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(newData)
	if err != nil {
		return newData, err
	}
	return newData, nil
}

func (m *mongoDB) DeleteCourse(ctx context.Context, id string) (data *pb.Course, err error) {
	data = &pb.Course{}
	filter := bson.M{"id": id}
	err = m.CCourse().FindOneAndDelete(ctx, filter).Decode(data)
	return
}

func (m *mongoDB) GetCourses(ctx context.Context, limit, skip int64, query, sort string) (count int64, datas []*pb.Course, err error) {
	datas = []*pb.Course{}

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

	cursor, err := m.CCourse().Find(ctx, filter, findOption)
	if err != nil {
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		data := &pb.Course{}
		err := cursor.Decode(data)
		if err != nil {
			logrus.WithError(err).Errorln("Course Decode Error")
			return 0, datas, err
		}
		datas = append(datas, data)
	}
	count, err = m.CCourse().CountDocuments(ctx, filter)
	if err != nil {
		logrus.WithError(err).Errorln("Course Count Documents Error")
		count = int64(0)
	}
	return
}
