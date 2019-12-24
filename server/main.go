package main

import (
	"context"
	"github.com/AbdulAbdulkadir/ascii/models"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
)

type server struct{}

func main() {

	models.StartMongoDB()
	models.SeedDB()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	//Create grpc server
	srv := grpc.NewServer()
	//Register server
	proto.RegisterAddServiceServer(srv, &server{})
	//Serialize and Deserialize data
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}



func (s *server) DisplayAscii(_ context.Context, _ *proto.DisplayRequest) (*proto.DisplayResponse, error) {

	log.Printf("Returning random ascii")

	//clearDB(*mongoClient)
	//
	//collection := mongoClient.Database("art").Collection("ascii")
	//
	//artArray := getAsciiArtFromFile()
	//
	//if err := insertAsciiArtDB(artArray, collection); err != nil {
	//	log.Printf("unexpected database error %+v", err)
	//	return nil, status.Error(codes.Internal, "internal server error - please try again later")
	//}

	asciiArray, err := models.GetAsciiArtFromDB()
	if err != nil {
		log.Printf("Could not get ascii from database %+v", err)
	}

	//	return nil, status.Error(codes.AlreadyExists, "there is already an ascii called cat")

	randIndex := int64(rand.Intn(len(asciiArray)))

	result := models.SelectAsciiArt(randIndex, asciiArray)

	return &proto.DisplayResponse{Sp: result}, nil
}


func (s *server) UploadAscii(_ context.Context, request *proto.UploadRequest) (*proto.UploadResponse, error) {

	log.Printf("Uploading ascii to server")

	err := models.UploadAsciiArt(request.Upload)
	if err != nil{
		log.Printf("Could not get upload ascii to server: %+v", err)
	}

	return &proto.UploadResponse{}, nil
}