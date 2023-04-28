package server

import (
	"context"

	"github.com/yushk/sport_backend/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCourse ...
func (s *Server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (data *pb.Course, err error) {
	data, err = s.db.CreateCourse(ctx, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetCourse ...
func (s *Server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (data *pb.Course, err error) {
	data, err = s.db.GetCourse(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// UpdateCourse ...
func (s *Server) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (data *pb.Course, err error) {
	data, err = s.db.UpdateCourse(ctx, req.Id, req.Data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// DeleteCourse ...
func (s *Server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (data *pb.Course, err error) {
	data, err = s.db.DeleteCourse(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return
}

// GetCourses ...
func (s *Server) GetCourses(ctx context.Context, req *pb.GetCoursesRequest) (reply *pb.GetCoursesReply, err error) {
	totalCount, users, err := s.db.GetCourses(ctx, req.Limit, req.Skip, req.Query, req.Sort)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	reply = &pb.GetCoursesReply{
		TotalCount: totalCount,
		Items:      users,
	}
	return
}
