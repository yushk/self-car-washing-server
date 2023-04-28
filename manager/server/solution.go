package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSolution ...
func (s *Server) CreateSolution(ctx context.Context, req *pb.CreateSolutionRequest) (data *pb.Solution, err error) {
	data, err = s.db.CreateSolution(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.db.AddHomeWorkSolution(ctx, data.WorkId, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSolution ...
func (s *Server) GetSolution(ctx context.Context, req *pb.GetSolutionRequest) (data *pb.Solution, err error) {
	data, err = s.db.GetSolution(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateSolution ...
func (s *Server) UpdateSolution(ctx context.Context, req *pb.UpdateSolutionRequest) (data *pb.Solution, err error) {
	data, err = s.db.UpdateSolution(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.db.UpdateHomeWorkSolution(ctx, data.WorkId, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteSolution ...
func (s *Server) DeleteSolution(ctx context.Context, req *pb.DeleteSolutionRequest) (data *pb.Solution, err error) {
	data, err = s.db.DeleteSolution(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetSolutions ...
func (s *Server) GetSolutions(ctx context.Context, req *pb.GetSolutionsRequest) (reply *pb.GetSolutionsReply, err error) {
	totalCount, users, err := s.db.GetSolutions(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetSolutionsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
