package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateHomeWork(params monitor.CreateHomeWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.HomeWork().CreateHomeWork(ctx, &pb.CreateHomeWorkRequest{
		Data: convert.HomeWorkV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateHomeWorkOK().WithPayload(convert.HomeWorkPb2V1(reply))
}

func DeleteHomeWork(params monitor.DeleteHomeWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.HomeWork().DeleteHomeWork(ctx, &pb.DeleteHomeWorkRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteHomeWorkOK().WithPayload(convert.HomeWorkPb2V1(reply))
}

func GetHomeWork(params monitor.GetHomeWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.HomeWork().GetHomeWork(ctx, &pb.GetHomeWorkRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetHomeWorkOK().WithPayload(convert.HomeWorkPb2V1(reply))
}

func UpdateHomeWork(params monitor.UpdateHomeWorkParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	reply, err := client.HomeWork().UpdateHomeWork(ctx, &pb.UpdateHomeWorkRequest{
		Id:   params.ID,
		Data: convert.HomeWorkV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateHomeWorkOK().WithPayload(convert.HomeWorkPb2V1(reply))
}

// ListHomeWork
func GetHomeWorks(params monitor.GetHomeWorksParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.HomeWork().GetHomeWorks(ctx, &pb.GetHomeWorksRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.HomeWork{}
	for _, v := range reply.Items {
		items = append(items, convert.HomeWorkPb2V1(v))
	}

	payload := &v1.HomeWorks{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetHomeWorksOK().WithPayload(payload)
}
