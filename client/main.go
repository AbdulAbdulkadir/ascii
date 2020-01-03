package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	var url string
	var fileName string

	flag.StringVar(&url, "url", "localhost:4040", "a string var")
	flag.StringVar(&fileName, "upload", "null", "a string var")
	flag.Parse()

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := proto.NewAsciiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//upload file
	if fileName != "null" {

		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Printf("could not read file")
		}

		_, err = client.UploadAscii(ctx, &proto.UploadRequest{Filename: fileName, Content: string(content)})
		if err != nil {
			log.Fatalf("Could not upload: %v", err)
		}
		log.Printf("Successfully uploaded " + fileName)
	}

	//display ascii
	var r *proto.DisplayResponse

	r, err = client.DisplayAscii(ctx, &proto.DisplayRequest{})
	if err != nil {
		log.Fatalf("Could not display: %v", err)
	}
	if r.GetDisplayAscii() == "empty" {
		fmt.Printf("Database is empty\n")
	} else {
		fmt.Printf("\n" + r.GetDisplayAscii())
	}

}
