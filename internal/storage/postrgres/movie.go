package postrgres

import (
	"context"
	"database/sql"
	pb "movie-service/internal/genproto/movieservice"
)

type MovieStorage interface {
	AddMovie(ctx context.Context, req *pb.AddMovieReq) (*pb.AddMovieRes,error)
	GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq ) (*pb.GetMovieByIdRes, error)
	UpdateMovie(ctx context.Context, req *pb.UpdateMovieReq) (*pb.UpdateMovieRes, error)
	DeleteMovie(ctx context.Context, req *pb.DeleteMovieReq) (*pb.DeleteMovieRes, error)
	RemoveMovie(ctx context.Context, req *pb.RemoveMovieReq) (*pb.RemoveMovieRes, error)
	GetAllMovies(ctx context.Context, req *pb.GetAllMoviesReq) (*pb.GetAllMoviesRes, error)
}
	
type MovieStorageImpl struct {
	db *sql.DB
}

func NewMovieStorage(db *sql.DB) MovieStorage {
	return &MovieStorageImpl{db: db}
}

func (s *MovieStorageImpl) AddMovie(ctx context.Context, req *pb.AddMovieReq) (*pb.AddMovieRes, error) {
	return nil, nil
}

func (s *MovieStorageImpl) GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq) (*pb.GetMovieByIdRes, error) {
	return nil, nil
}

func (s *MovieStorageImpl) UpdateMovie(ctx context.Context, req *pb.UpdateMovieReq) (*pb.UpdateMovieRes, error) {
	return nil, nil
}
func (s *MovieStorageImpl) DeleteMovie(ctx context.Context, req *pb.DeleteMovieReq) (*pb.DeleteMovieRes, error) {
	return nil, nil
}
func (s *MovieStorageImpl) RemoveMovie(ctx context.Context, req *pb.RemoveMovieReq) (*pb.RemoveMovieRes, error) {
	return nil, nil
}
func (s *MovieStorageImpl) GetAllMovies(ctx context.Context, req *pb.GetAllMoviesReq) (*pb.GetAllMoviesRes, error) {
	return nil, nil
}