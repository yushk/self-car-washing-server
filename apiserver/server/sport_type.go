package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateSportType(params monitor.CreateSportTypeParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportType().CreateSportType(ctx, &pb.CreateSportTypeRequest{
		Data: convert.SportTypeV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateSportTypeOK().WithPayload(convert.SportTypePb2V1(reply))
}

func DeleteSportType(params monitor.DeleteSportTypeParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportType().DeleteSportType(ctx, &pb.DeleteSportTypeRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteSportTypeOK().WithPayload(convert.SportTypePb2V1(reply))
}

func GetSportType(params monitor.GetSportTypeParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportType().GetSportType(ctx, &pb.GetSportTypeRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetSportTypeOK().WithPayload(convert.SportTypePb2V1(reply))
}

func UpdateSportType(params monitor.UpdateSportTypeParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportType().UpdateSportType(ctx, &pb.UpdateSportTypeRequest{
		Id:   params.ID,
		Data: convert.SportTypeV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateSportTypeOK().WithPayload(convert.SportTypePb2V1(reply))
}

// ListSportType
func GetSportTypes(params monitor.GetSportTypesParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.SportType().GetSportTypes(ctx, &pb.GetSportTypesRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.SportType{}
	for _, v := range reply.Items {
		items = append(items, convert.SportTypePb2V1(v))
	}

	payload := &v1.SportTypes{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetSportTypesOK().WithPayload(payload)
}
