package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateEvaluateInfo(params monitor.CreateEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.EvaluateInfo().CreateEvaluateInfo(ctx, &pb.CreateEvaluateInfoRequest{
		Data: convert.EvaluateInfoV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateEvaluateInfoOK().WithPayload(convert.EvaluateInfoPb2V1(reply))
}

func DeleteEvaluateInfo(params monitor.DeleteEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.EvaluateInfo().DeleteEvaluateInfo(ctx, &pb.DeleteEvaluateInfoRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteEvaluateInfoOK().WithPayload(convert.EvaluateInfoPb2V1(reply))
}

func UpdateEvaluateInfo(params monitor.UpdateEvaluateInfoParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.EvaluateInfo().UpdateEvaluateInfo(ctx, &pb.UpdateEvaluateInfoRequest{
		Id:   params.ID,
		Data: convert.EvaluateInfoV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateEvaluateInfoOK().WithPayload(convert.EvaluateInfoPb2V1(reply))
}

// ListEvaluateInfo
func GetEvaluateInfos(params monitor.GetEvaluateInfosParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.EvaluateInfo().GetEvaluateInfos(ctx, &pb.GetEvaluateInfosRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.EvaluateInfo{}
	for _, v := range reply.Items {
		items = append(items, convert.EvaluateInfoPb2V1(v))
	}

	payload := &v1.EvaluateInfos{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetEvaluateInfosOK().WithPayload(payload)
}
