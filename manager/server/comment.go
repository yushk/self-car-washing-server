package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateComment ...
func (s *Server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (data *pb.Comment, err error) {
	data, err = s.db.CreateComment(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetComment ...
func (s *Server) GetComment(ctx context.Context, req *pb.GetCommentRequest) (data *pb.Comment, err error) {
	data, err = s.db.GetComment(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateComment ...
func (s *Server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (data *pb.Comment, err error) {
	data, err = s.db.UpdateComment(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteComment ...
func (s *Server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (data *pb.Comment, err error) {
	data, err = s.db.DeleteComment(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetComments ...
func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (reply *pb.GetCommentsReply, err error) {
	totalCount, users, err := s.db.GetComments(ctx, req.Limit, req.Skip, req.Query)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetCommentsReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
