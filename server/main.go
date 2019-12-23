package main

import (
	"context"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"github.com/AbdulAbdulkadir/ascii/server/models"
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



func (s *server) DisplayAscii(_ context.Context, _ *proto.AsciiRequest) (*proto.AsciiResponse, error) {

	log.Printf("returning random ascii")

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

	asciiArray := models.RetrieveAsciiArtArray()

	//	return nil, status.Error(codes.AlreadyExists, "there is already an ascii called cat")

	randIndex := int64(rand.Intn(len(asciiArray)))

	selectIndex := models.SelectAsciiArt(randIndex, asciiArray)

	return &proto.AsciiResponse{Sp: selectIndex}, nil
}


