package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSportItem ...
func (s *Server) CreateSportItem(ctx context.Context, req *pb.CreateSportItemRequest) (data *pb.SportItem, err error) {
	data, err = s.db.CreateSportItem(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSportItem ...
func (s *Server) GetSportItem(ctx context.Context, req *pb.GetSportItemRequest) (data *pb.SportItem, err error) {
	data, err = s.db.GetSportItem(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateSportItem ...
func (s *Server) UpdateSportItem(ctx context.Context, req *pb.UpdateSportItemRequest) (data *pb.SportItem, err error) {
	data, err = s.db.UpdateSportItem(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteSportItem ...
func (s *Server) DeleteSportItem(ctx context.Context, req *pb.DeleteSportItemRequest) (data *pb.SportItem, err error) {
	data, err = s.db.DeleteSportItem(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSportItems ...
func (s *Server) GetSportItems(ctx context.Context, req *pb.GetSportItemsRequest) (reply *pb.GetSportItemsReply, err error) {
	totalCount, users, err := s.db.GetSportItems(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetSportItemsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
