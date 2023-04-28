package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateCourse(params monitor.CreateCourseParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Course().CreateCourse(ctx, &pb.CreateCourseRequest{
		Data: convert.CourseV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateCourseOK().WithPayload(convert.CoursePb2V1(reply))
}

func DeleteCourse(params monitor.DeleteCourseParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Course().DeleteCourse(ctx, &pb.DeleteCourseRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteCourseOK().WithPayload(convert.CoursePb2V1(reply))
}

func GetCourse(params monitor.GetCourseParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Course().GetCourse(ctx, &pb.GetCourseRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetCourseOK().WithPayload(convert.CoursePb2V1(reply))
}

func UpdateCourse(params monitor.UpdateCourseParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Course().UpdateCourse(ctx, &pb.UpdateCourseRequest{
		Id:   params.ID,
		Data: convert.CourseV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateCourseOK().WithPayload(convert.CoursePb2V1(reply))
}

// ListCourse
func GetCourses(params monitor.GetCoursesParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	limit := int64(0)
	if params.Limit != nil {
		limit = *params.Limit
	}
	skip := int64(0)
	if params.Skip != nil {
		skip = *params.Skip
	}
	query := `{}`
	if params.Query != nil {
		query = *params.Query
	}
	sort := `{}`
	if params.Sort != nil {
		sort = *params.Sort
	}
	reply, err := client.Course().GetCourses(ctx, &pb.GetCoursesRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
		Sort:  sort,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Course{}
	for _, v := range reply.Items {
		items = append(items, convert.CoursePb2V1(v))
	}

	payload := &v1.Courses{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetCoursesOK().WithPayload(payload)
}
