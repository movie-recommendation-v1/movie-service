package movie

import (
	"context"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
	"movie-service/internal/storage/postrgres"
)

type MovService interface {
	AddMovie(ctx context.Context, req *pb.AddMovieReq) (*pb.AddMovieRes, error)
	GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq) (*pb.GetMovieByIdRes, error)
	UpdateMovie(ctx context.Context, req *pb.UpdateMovieReq) (*pb.UpdateMovieRes, error)
	DeleteMovie(ctx context.Context, req *pb.DeleteMovieReq) (*pb.DeleteMovieRes, error)
	RemoveMovie(ctx context.Context, req *pb.RemoveMovieReq) (*pb.RemoveMovieRes, error)
	GetAllMovies(ctx context.Context, req *pb.GetAllMoviesReq) (*pb.GetAllMoviesRes, error)
}

type MovServiceImpl struct {
	movie  postrgres.MovieStorage
	client pb.MovieServiceClient
	pb.UnimplementedMovieServiceServer
}

func NewMovieService(movie postrgres.MovieStorage, client pb.MovieServiceClient) MovService {
	return &MovServiceImpl{movie: movie, client: client}
}

func (s *MovServiceImpl) AddMovie(ctx context.Context, req *pb.AddMovieReq) (*pb.AddMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.AddMovie(ctx, req)
	if err != nil {
		logs.Error("Error while calling Add movie")
	}
	logs.Info("Successfully add movie")
	return resp, nil
}

func (s *MovServiceImpl) GetMovieById(ctx context.Context, req *pb.GetMovieByIdReq) (*pb.GetMovieByIdRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.GetMovieById(ctx, req)
	if err != nil {
		logs.Error("Error while calling Get movie by id")
		return nil, err
	}

	logs.Info("Successfully retrieved movie")
	return resp, nil
}

func (s *MovServiceImpl) UpdateMovie(ctx context.Context, req *pb.UpdateMovieReq) (*pb.UpdateMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.UpdateMovie(ctx, req)
	if err != nil {
		logs.Error("Error while calling Update movie")
		return nil, err
	}
	logs.Info("Successfully updated movie")
	return resp, nil
}

func (s *MovServiceImpl) DeleteMovie(ctx context.Context, req *pb.DeleteMovieReq) (*pb.DeleteMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.DeleteMovie(ctx, req)
	if err != nil {
		logs.Error("Error while calling Delete movie")
		return nil, err
	}
	logs.Info("Successfully deleted movie")
	return resp, nil
}

func (s *MovServiceImpl) RemoveMovie(ctx context.Context, req *pb.RemoveMovieReq) (*pb.RemoveMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.RemoveMovie(ctx, req)
	if err != nil {
		logs.Error("Error while calling Remove movie")
		return nil, err
	}
	logs.Info("Successfully removed movie")
	return resp, nil
}

func (s *MovServiceImpl) GetAllMovies(ctx context.Context, req *pb.GetAllMoviesReq) (*pb.GetAllMoviesRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	resp, err := s.movie.GetAllMovies(ctx, req)
	if err != nil {
		logs.Error("Error while calling Get all movies")
		return nil, err
	}
	logs.Info("Successfully retrieved all movies")
	return resp, nil
}
