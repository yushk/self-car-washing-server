package server

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yushk/sport_backend/apiserver/restapi/operations/monitor"
	"github.com/yushk/sport_backend/apiserver/server/convert"
	v1 "github.com/yushk/sport_backend/apiserver/v1"
	"github.com/yushk/sport_backend/pkg/pb"
)

func CreateComment(params monitor.CreateCommentParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Comment().CreateComment(ctx, &pb.CreateCommentRequest{
		Data: convert.CommentV12Pb(params.Body),
	})
	if err != nil {
		logrus.Errorln(err)
		return Error(err)
	}
	return monitor.NewCreateCommentOK().WithPayload(convert.CommentPb2V1(reply))
}

func DeleteComment(params monitor.DeleteCommentParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Comment().DeleteComment(ctx, &pb.DeleteCommentRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewDeleteCommentOK().WithPayload(convert.CommentPb2V1(reply))
}

func GetComment(params monitor.GetCommentParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Comment().GetComment(ctx, &pb.GetCommentRequest{
		Id: params.ID,
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewGetCommentOK().WithPayload(convert.CommentPb2V1(reply))
}

func UpdateComment(params monitor.UpdateCommentParams, principal *v1.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	reply, err := client.Comment().UpdateComment(ctx, &pb.UpdateCommentRequest{
		Id:   params.ID,
		Data: convert.CommentV12Pb(params.Body),
	})
	if err != nil {
		return Error(err)
	}
	return monitor.NewUpdateCommentOK().WithPayload(convert.CommentPb2V1(reply))
}

// ListComment
func GetComments(params monitor.GetCommentsParams, principal *v1.Principal) middleware.Responder {
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
	reply, err := client.Comment().GetComments(ctx, &pb.GetCommentsRequest{
		Limit: limit,
		Skip:  skip,
		Query: query,
	})
	if err != nil {
		return Error(err)
	}

	items := []*v1.Comment{}
	for _, v := range reply.Items {
		items = append(items, convert.CommentPb2V1(v))
	}

	payload := &v1.Comments{
		TotalCount: reply.TotalCount,
		Items:      items,
	}
	return monitor.NewGetCommentsOK().WithPayload(payload)
}
