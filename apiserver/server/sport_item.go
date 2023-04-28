package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateSportItem(params monitor.CreateSportItemParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportItem().CreateSportItem(ctx, &pb.CreateSportItemRequest{
		Data: convert.SportItemV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateSportItemOK().WithPayload(convert.SportItemPb2V1(reply))
}

func DeleteSportItem(params monitor.DeleteSportItemParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportItem().DeleteSportItem(ctx, &pb.DeleteSportItemRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteSportItemOK().WithPayload(convert.SportItemPb2V1(reply))
}

func GetSportItem(params monitor.GetSportItemParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportItem().GetSportItem(ctx, &pb.GetSportItemRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetSportItemOK().WithPayload(convert.SportItemPb2V1(reply))
}

func UpdateSportItem(params monitor.UpdateSportItemParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.SportItem().UpdateSportItem(ctx, &pb.UpdateSportItemRequest{
		Id:   params.ID,
		Data: convert.SportItemV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateSportItemOK().WithPayload(convert.SportItemPb2V1(reply))
}

// ListSportItem
func GetSportItems(params monitor.GetSportItemsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.SportItem().GetSportItems(ctx, &pb.GetSportItemsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.SportItem{}
	for _, v := range reply.Items {
		items = append(items, convert.SportItemPb2V1(v))
	}

	payload := &v1.SportItems{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetSportItemsOK().WithPayload(payload)
}
