package postrgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
	"strconv"
	"time"
)

type CommentsStorage interface {
	CreateComment(ctx context.Context, req *pb.CreateCommentReq) (*pb.CreateCommentRes, error)
	UpdateComment(ctx context.Context, req *pb.UpdateCommentReq) (*pb.UpdateCommentRes, error)
	GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentRes, error)
	DeleteComment(ctx context.Context, req *pb.DeleteCommentReq) (*pb.DeleteCommentRes, error)
	GetAllComments(ctx context.Context, req *pb.GetAllCommentsReq) (*pb.GetAllCommentsRes, error)
}

type CommentsStorageImpl struct {
	db *sql.DB
}

func NewCommentsStorage(db *sql.DB) CommentsStorage {
	return &CommentsStorageImpl{
		db: db,
	}
}
func (s *CommentsStorageImpl) CreateComment(ctx context.Context, req *pb.CreateCommentReq) (*pb.CreateCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()
	query := `insert into comments(id, user_id, movie_id, description, rate) values ($1, $2, $3, $4, $5)`
	comment := pb.CommentModel{}
	err = s.db.QueryRow(query, id, req.UserId, req.MovieId, req.Description, req.Rate).Scan(
		&comment.Id, &comment.UserId, &comment.MovieId, &comment.Description, &comment.Rate, &comment.CreatedAt)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	logs.Info("Successfully create comment!")
	return &pb.CreateCommentRes{Comment: &comment}, nil
}

func (s *CommentsStorageImpl) UpdateComment(ctx context.Context, req *pb.UpdateCommentReq) (*pb.UpdateCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `UPDATE comments SET `
	args := []interface{}{}
	argCounter := 1

	if req.Description != "" {
		query += `description = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Description)
		argCounter++
	}
	if req.Rate != 0 {
		query += `rate = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Rate)
		argCounter++
	}

	if len(args) == 0 {
		logs.Error("Error while updating comment")
		return nil, fmt.Errorf("no fields to update")
	}

	query = query[:len(query)-2] + ` WHERE id = $` + strconv.Itoa(argCounter)
	args = append(args, req.Id)

	comment := pb.CommentModel{}

	err = s.db.QueryRowContext(ctx, query, args...).Scan(
		&comment.Id, &comment.UserId, &comment.MovieId, &comment.Description,
		&comment.Rate, &comment.CreatedAt, &comment.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update comment with id %s: %w", req.Id, err)
	}
	logs.Info("Successfully updated comment!")
	return &pb.UpdateCommentRes{Comment: &comment}, nil

}

func (s *CommentsStorageImpl) GetComment(ctx context.Context, req *pb.GetCommentReq) (*pb.GetCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `select * from comments where id = $1 deleted_at = 0`
	comment := pb.CommentModel{}
	err = s.db.QueryRowContext(ctx, query, req.Id).Scan(
		&comment.Id, &comment.UserId, &comment.MovieId,
		&comment.Description, &comment.Rate, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	logs.Info("Successfully get comment!")
	return &pb.GetCommentRes{Comment: &comment}, nil
}

func (s *CommentsStorageImpl) DeleteComment(ctx context.Context, req *pb.DeleteCommentReq) (*pb.DeleteCommentRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `update comments set deleted_at = $1 where id = $2`
	_, err = s.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	logs.Info("Successfully deleted comment!")
	return &pb.DeleteCommentRes{Success: true}, nil
}

func (s *CommentsStorageImpl) GetAllComments(ctx context.Context, req *pb.GetAllCommentsReq) (*pb.GetAllCommentsRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `SELECT id, user_id, movie_id, description, rate FROM comments WHERE 1=1`
	args := []interface{}{}
	argCounter := 1

	if req.UserId != "" {
		query += ` AND user_id = $` + strconv.Itoa(argCounter)
		args = append(args, req.UserId)
		argCounter++
	}
	if req.MovieId != "" {
		query += ` AND movie_id = $` + strconv.Itoa(argCounter)
		args = append(args, req.MovieId)
		argCounter++
	}
	if req.Description != "" {
		query += ` AND description ILIKE $` + strconv.Itoa(argCounter)
		args = append(args, "%"+req.Description+"%")
		argCounter++
	}
	if req.Rate != 0 {
		query += ` AND rate = $` + strconv.Itoa(argCounter)
		args = append(args, req.Rate)
		argCounter++
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	defer rows.Close()

	var comments []*pb.CommentModel
	for rows.Next() {
		var comment pb.CommentModel
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.MovieId, &comment.Description, &comment.Rate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}

	logs.Info("Successfully get all comments!")
	return &pb.GetAllCommentsRes{Comments: comments}, nil
}
