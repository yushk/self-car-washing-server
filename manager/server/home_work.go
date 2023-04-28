package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateHomeWork ...
func (s *Server) CreateHomeWork(ctx context.Context, req *pb.CreateHomeWorkRequest) (data *pb.HomeWork, err error) {
	data, err = s.db.CreateHomeWork(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetHomeWork ...
func (s *Server) GetHomeWork(ctx context.Context, req *pb.GetHomeWorkRequest) (data *pb.HomeWork, err error) {
	data, err = s.db.GetHomeWork(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	data.SportItem ,err= s.db.GetSportItem(ctx,data.SportItemId)

	return
}

// UpdateHomeWork ...
func (s *Server) UpdateHomeWork(ctx context.Context, req *pb.UpdateHomeWorkRequest) (data *pb.HomeWork, err error) {
	logrus.Debugln(req.Data);
	data, err = s.db.UpdateHomeWork(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteHomeWork ...
func (s *Server) DeleteHomeWork(ctx context.Context, req *pb.DeleteHomeWorkRequest) (data *pb.HomeWork, err error) {
	data, err = s.db.DeleteHomeWork(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetHomeWorks ...
func (s *Server) GetHomeWorks(ctx context.Context, req *pb.GetHomeWorksRequest) (reply *pb.GetHomeWorksReply, err error) {
	totalCount, users, err := s.db.GetHomeWorks(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for i:=range users{
		users[i].SportItem,err = s.db.GetSportItem(ctx,users[i].SportItemId)
	}
	reply = &pb.GetHomeWorksReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
