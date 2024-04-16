package main

import (
	"context"
	"log"
	"net"

	"github.com/AbdulAbdulkadir/ascii/proto"
	"github.com/AbdulAbdulkadir/ascii/server/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/service.proto
type server struct {
	proto.UnimplementedAsciiServiceServer
}

func main() {

	err := models.StartMongoDB()
	if err != nil {
		log.Println("could not start database")
		return
	}

	isEmpty, err := models.IsDatabaseEmpty()
	if err != nil {
		log.Printf("Error checking if databse is empty")
		return
	}

	if isEmpty {
		log.Printf("Database empty, will seed...")
		err := models.SeedDB()
		if err != nil {
			log.Println("could not seed database", err)
			return
		}
		log.Printf("Seeding complete!")
	}

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// Create grpc server
	srv := grpc.NewServer()
	// Register server
	proto.RegisterAsciiServiceServer(srv, &server{})
	// Serialize and Deserialize data
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) DisplayAscii(_ context.Context, _ *proto.DisplayRequest) (*proto.DisplayResponse, error) {

	log.Printf("Returning random ascii")

	isEmpty, err := models.IsDatabaseEmpty()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	if isEmpty {
		return nil, status.Error(codes.FailedPrecondition, "database empty")
	}

	result, err := models.GetRandomArt()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &proto.DisplayResponse{DisplayAscii: result.Art}, nil
}

func (s *server) UploadAscii(_ context.Context, request *proto.UploadRequest) (*proto.UploadResponse, error) {

	if request.Content == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty String")
	}

	err := models.UploadAsciiArt(request.Filename, request.Content)
	if err != nil {
		log.Printf("Could not upload ascii to server: %+v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	log.Printf("Successfully uploaded file")
	return &proto.UploadResponse{}, nil
}
