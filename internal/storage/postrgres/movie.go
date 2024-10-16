package postrgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
	"strconv"
	"strings"
	"time"
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
		&movie.Res.Season,
		&movie.Res.BackgroundImageUrl,
		&movie.Res.MovieUrl,
		&movie.Res.Studio,
		&movie.Res.Bio,
		&movie.Res.Genres,
	)
	if err != nil {
		logs.Error(err.Error())
	}

	logs.Info("Successfully retrieved movie")
	return movie, nil
}

func (s *MovieStorageImpl) UpdateMovie(ctx context.Context, req *pb.UpdateMovieReq) (*pb.UpdateMovieRes, error) {

	query := `UPDATE movies SET `
	args := []interface{}{}
	argCounter := 1

	if req.MovieName != "" {
		query += `movie_name = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.MovieName)
		argCounter++
	}
	if req.AgeLimit != 0 {
		query += `age_limit = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.AgeLimit)
		argCounter++
	}
	if req.BackgroundImageUrl != "" {
		query += `background_image_url = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.BackgroundImageUrl)
		argCounter++
	}
	if req.MovieUrl != "" {
		query += `movie_url = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.MovieUrl)
		argCounter++
	}
	if req.Studio != "" {
		query += `studio = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Studio)
		argCounter++
	}
	if req.Bio != "" {
		query += `bio = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Bio)
		argCounter++
	}
	if len(req.Genres) > 0 {
		query += `genres = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, pq.Array(req.Genres))
		argCounter++
	}
	if req.Language != "" {
		query += `language = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Language)
		argCounter++
	}
	if req.Season != 0 {
		query += `season = $` + strconv.Itoa(argCounter) + `, `
		args = append(args, req.Season)
		argCounter++
	}

	query = query[:len(query)-2] + ` WHERE id = $` + strconv.Itoa(argCounter)
	args = append(args, req.Id)

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update movie: %w", err)
	}

	return &pb.UpdateMovieRes{Success: true}, nil
}

func (s *MovieStorageImpl) DeleteMovie(ctx context.Context, req *pb.DeleteMovieReq) (*pb.DeleteMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `delete from movies where id = $1`
	_, err = s.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info("Successfully deleted movie")
	return &pb.DeleteMovieRes{Success: true}, nil
}

func (s *MovieStorageImpl) RemoveMovie(ctx context.Context, req *pb.RemoveMovieReq) (*pb.RemoveMovieRes, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	query := `update movie set deleted_at = $1 where id = $2`
	t := time.Now().Unix()
	_, err = s.db.ExecContext(ctx, query, t, req.Id)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Info("Successfully removed movie")
	return &pb.RemoveMovieRes{Success: true}, nil
}

func (s *MovieStorageImpl) GetAllMovies(ctx context.Context, req *pb.GetAllMoviesReq) (*pb.GetAllMoviesRes, error) {
	query := `SELECT id, movie_name, age_limit, background_image_url, movie_url, studio, bio, genres, language, season 
	          FROM movies WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM movies WHERE 1=1`

	filters := []string{}
	args := []interface{}{}
	argCounter := 1

	if req.MovieName != "" {
		filters = append(filters, fmt.Sprintf("movie_name ILIKE $%d", argCounter))
		args = append(args, "%"+req.MovieName+"%")
		argCounter++
	}
	if req.AgeLimit != 0 {
		filters = append(filters, fmt.Sprintf("age_limit = $%d", argCounter))
		args = append(args, req.AgeLimit)
		argCounter++
	}
	if req.BackgroundImageUrl != "" {
		filters = append(filters, fmt.Sprintf("background_image_url = $%d", argCounter))
		args = append(args, req.BackgroundImageUrl)
		argCounter++
	}
	if req.MovieUrl != "" {
		filters = append(filters, fmt.Sprintf("movie_url = $%d", argCounter))
		args = append(args, req.MovieUrl)
		argCounter++
	}
	if req.Studio != "" {
		filters = append(filters, fmt.Sprintf("studio ILIKE $%d", argCounter))
		args = append(args, "%"+req.Studio+"%")
		argCounter++
	}
	if req.Bio != "" {
		filters = append(filters, fmt.Sprintf("bio ILIKE $%d", argCounter))
		args = append(args, "%"+req.Bio+"%")
		argCounter++
	}
	if len(req.Genres) > 0 {
		filters = append(filters, fmt.Sprintf("genres @> $%d", argCounter))
		args = append(args, pq.Array(req.Genres))
		argCounter++
	}
	if req.Language != "" {
		filters = append(filters, fmt.Sprintf("language = $%d", argCounter))
		args = append(args, req.Language)
		argCounter++
	}
	if req.Season != 0 {
		filters = append(filters, fmt.Sprintf("season = $%d", argCounter))
		args = append(args, req.Season)
		argCounter++
	}

	if len(filters) > 0 {
		filterQuery := " AND " + strings.Join(filters, " AND ")
		query += filterQuery
		countQuery += filterQuery
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1)
	args = append(args, req.Limit, req.Offset)

	var totalCount int32
	err := s.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count of movies: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get movies: %w", err)
	}
	defer rows.Close()

	var movies []*pb.MovieModel
	for rows.Next() {
		var movie pb.MovieModel
		var genres []*pb.Genres
		err := rows.Scan(&movie.Id, &movie.MovieName, &movie.AgeLimit, &movie.BackgroundImageUrl, &movie.MovieUrl, &movie.Studio, &movie.Bio, pq.Array(&genres), &movie.Language, &movie.Season)
		if err != nil {
			return nil, fmt.Errorf("failed to scan movie: %w", err)
		}
		movie.Genres = genres
		movies = append(movies, &movie)
	}

	return &pb.GetAllMoviesRes{
		Movies:     movies,
		TotalCount: totalCount,
	}, nil
}
