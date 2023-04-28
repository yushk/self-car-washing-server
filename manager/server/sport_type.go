package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSportType ...
func (s *Server) CreateSportType(ctx context.Context, req *pb.CreateSportTypeRequest) (data *pb.SportType, err error) {
	data, err = s.db.CreateSportType(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSportType ...
func (s *Server) GetSportType(ctx context.Context, req *pb.GetSportTypeRequest) (data *pb.SportType, err error) {
	data, err = s.db.GetSportType(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateSportType ...
func (s *Server) UpdateSportType(ctx context.Context, req *pb.UpdateSportTypeRequest) (data *pb.SportType, err error) {
	data, err = s.db.UpdateSportType(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteSportType ...
func (s *Server) DeleteSportType(ctx context.Context, req *pb.DeleteSportTypeRequest) (data *pb.SportType, err error) {
	data, err = s.db.DeleteSportType(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSportTypes ...
func (s *Server) GetSportTypes(ctx context.Context, req *pb.GetSportTypesRequest) (reply *pb.GetSportTypesReply, err error) {
	totalCount, users, err := s.db.GetSportTypes(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetSportTypesReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
