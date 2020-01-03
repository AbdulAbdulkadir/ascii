package main

import (
	"context"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"github.com/AbdulAbdulkadir/ascii/server/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"testing"
)

var client proto.AsciiServiceClient

func startServer() {
	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}

	//Create grpc server
	srv := grpc.NewServer()
	//Register server
	proto.RegisterAsciiServiceServer(srv, &server{})
	//Serialize and Deserialize data
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func TestMain(m *testing.M) {
	models.StartTestDB()
	go startServer()
	conn, err := grpc.Dial("localhost:4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = proto.NewAsciiServiceClient(conn)

	code := m.Run()
	models.CloseDB()
	os.Exit(code)
}

func TestServer_DisplayAscii(t *testing.T) {
	_, err := client.DisplayAscii(context.TODO(), &proto.DisplayRequest{})
	if err != nil {
		t.Fail()
	}
}

func TestServer_UploadAscii(t *testing.T) {

	t.Run("should run with string", func(t *testing.T) {
		name := "dummyName"
		str := "something"
		_, err := client.UploadAscii(context.TODO(), &proto.UploadRequest{Filename: name, Content: str})
		if err != nil {
			t.Fail()
		}
	})

	t.Run("should fail with empty string", func(t *testing.T) {
		name := "dummyName"
		str := ""
		_, err := client.UploadAscii(context.TODO(), &proto.UploadRequest{Filename: name, Content: str})
		if err == nil {
			t.Fail()
		}
	})

}
