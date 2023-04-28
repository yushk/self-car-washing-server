package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateEvaluateInfo ...
func (s *Server) CreateEvaluateInfo(ctx context.Context, req *pb.CreateEvaluateInfoRequest) (data *pb.EvaluateInfo, err error) {
	data, err = s.db.CreateEvaluateInfo(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateEvaluateInfo ...
func (s *Server) UpdateEvaluateInfo(ctx context.Context, req *pb.UpdateEvaluateInfoRequest) (data *pb.EvaluateInfo, err error) {
	data, err = s.db.UpdateEvaluateInfo(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteEvaluateInfo ...
func (s *Server) DeleteEvaluateInfo(ctx context.Context, req *pb.DeleteEvaluateInfoRequest) (data *pb.EvaluateInfo, err error) {
	data, err = s.db.DeleteEvaluateInfo(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetEvaluateInfos ...
func (s *Server) GetEvaluateInfos(ctx context.Context, req *pb.GetEvaluateInfosRequest) (reply *pb.GetEvaluateInfosReply, err error) {
	totalCount, users, err := s.db.GetEvaluateInfos(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetEvaluateInfosReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
