package postrgres

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
)

type MovieStorage interface {
	AddMovie(ctx context.Context, req *pb.AddMovieReq) (*pb.AddMovieRes, error)
	GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq) (*pb.GetMovieByIdRes, error)
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
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	id := uuid.NewString()

	query := `insert into movies (moviename,agelimit,season,beckround_image_url,movie_url,
                  studio, bio , genres) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = s.db.ExecContext(ctx, query, id, req.MovieName, req.AgeLimit, req.Season, req.BackgroundImageUrl,
		req.MovieUrl, req.Studio, req.Bio, req.Genres)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info("Successfully added movie")
	return &pb.AddMovieRes{Message: "Successfully added movie"}, nil
}

func (s *MovieStorageImpl) GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq) (*pb.GetMovieByIdRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `select id, moviename,agelimit, season, beckround_image_url, beckround_image_url,studio,bio, genres from movies where id = $1`
	row := s.db.QueryRowContext(ctx, query, req.Id)
	movie := &pb.GetMovieByIdRes{}
	err = row.Scan(
		&movie.Res.Id,
		&movie.Res.MovieName,
		&movie.Res.AgeLimit,
	)

	logs.Info("Successfully retrieved movie")

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
