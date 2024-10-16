package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/grpc/reflection"
	"movie-service/genproto/movieservice"
	pb "movie-service/genproto/movieservice"
	"movie-service/internal/logger"
	"movie-service/internal/service/comment"
	"movie-service/internal/service/movie"
	"movie-service/internal/storage/postrgres"
	"net"
)

func main() {

	logs, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	db, err := postrgres.Connect()
	if err != nil {
		logs.Error("Error connecting to database")
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":5052")
	if err != nil {
		logs.Error("Error starting listener")
	}
	defer listener.Close()

	logs.Info("Movie service running on port: 5052")

	s := grpc.NewServer()

	Conn, err := grpc.NewClient(":5052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Error("Error starting gRPC client")
	}
	defer Conn.Close()

	movieClient := movieservice.NewMovieServiceClient(Conn)
	commentClient := movieservice.NewCommentsServiceClient(Conn)

	moviedb := postrgres.NewMovieStorage(db)
	commentdb := postrgres.NewCommentsStorage(db)

	moviesv := movie.NewMovieService(moviedb, movieClient)
	commentsv := comment.NewCommentsService(commentdb, commentClient)

	pb.RegisterMovieServiceServer(s, moviesv)
	pb.RegisterCommentsServiceServer(s, commentsv)

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		logs.Error("Error starting server")
		panic(err)
	}

}
