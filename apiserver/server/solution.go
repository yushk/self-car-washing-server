package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateSolution(params monitor.CreateSolutionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Solution().CreateSolution(ctx, &pb.CreateSolutionRequest{
		Data: convert.SolutionV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateSolutionOK().WithPayload(convert.SolutionPb2V1(reply))
}

func DeleteSolution(params monitor.DeleteSolutionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Solution().DeleteSolution(ctx, &pb.DeleteSolutionRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteSolutionOK().WithPayload(convert.SolutionPb2V1(reply))
}

func GetSolution(params monitor.GetSolutionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Solution().GetSolution(ctx, &pb.GetSolutionRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetSolutionOK().WithPayload(convert.SolutionPb2V1(reply))
}

func UpdateSolution(params monitor.UpdateSolutionParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Solution().UpdateSolution(ctx, &pb.UpdateSolutionRequest{
		Id:   params.ID,
		Data: convert.SolutionV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateSolutionOK().WithPayload(convert.SolutionPb2V1(reply))
}

// ListSolution
func GetSolutions(params monitor.GetSolutionsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.Solution().GetSolutions(ctx, &pb.GetSolutionsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Solution{}
	for _, v := range reply.Items {
		items = append(items, convert.SolutionPb2V1(v))
	}

	payload := &v1.Solutions{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetSolutionsOK().WithPayload(payload)
}
