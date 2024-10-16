package comment

import (
	"context"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
	"movie-service/internal/storage/postrgres"
)

type CommentsService interface {
	CreateComment(ctx context.Context, req *pb.CreateCommentReq) (*pb.CreateCommentRes, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentReq) (*pb.UpdateCommentRes, error)
	GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentRes, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentReq) (*pb.DeleteCommentRes, error)
	GetAllComments(ctx context.Context, req *pb.GetAllCommentsReq) (*pb.GetAllCommentsRes, error)
}

type CommentsServiceImpl struct {
	comment postrgres.CommentsStorage
	client  pb.CommentsServiceClient
	pb.UnimplementedCommentsServiceServer
}

func NewCommentsService(comm postrgres.CommentsStorage, client pb.CommentsServiceClient) CommentsService {
	return &CommentsServiceImpl{
		comment: comm,
		client:  client,
	}
}
func (s *CommentsServiceImpl) CreateComment(ctx context.Context, req *pb.CreateCommentReq) (*pb.CreateCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.comment.CreateComment(ctx, req)
	if err != nil {
		logs.Error("Error while calling CreateComment")
	}
	logs.Info("Create Comment Success")
	return resp, nil

}

func (s *CommentsServiceImpl) UpdateComment(ctx context.Context, req *pb.UpdateCommentReq) (*pb.UpdateCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.comment.UpdateComment(ctx, req)
	if err != nil {
		logs.Error("Error while calling UpdateComment")
	}
	logs.Info("Update Comment Success")
	return resp, nil
}

func (s *CommentsServiceImpl) GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.comment.GetComment(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetComment")
	}
	logs.Info("Get Comment Success")
	return resp, nil
}

func (s *CommentsServiceImpl) DeleteComment(ctx context.Context, req *pb.DeleteCommentReq) (*pb.DeleteCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.comment.DeleteComment(ctx, req)
	if err != nil {
		logs.Error("Error while calling DeleteComment")
	}
	logs.Info("Delete Comment Success")
	return resp, nil
}

func (s *CommentsServiceImpl) GetAllComments(ctx context.Context, req *pb.GetAllCommentsReq) (*pb.GetAllCommentsRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.comment.GetAllComments(ctx, req)
	if err != nil {
		logs.Error("Error while calling GetAllComments")
	}
	logs.Info("Get All Comments Success")
	return resp, nil
}
