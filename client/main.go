package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AbdulAbdulkadir/ascii/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	var url string
	var fileName string

	flag.StringVar(&url, "url", "localhost:4040", "a string var")
	flag.StringVar(&fileName, "upload", "", "a string var")
	flag.Parse()

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	client := proto.NewAsciiServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Upload file
	if fileName != "" {
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Printf("could not read file")
			return
		}

		_, err = client.UploadAscii(ctx, &proto.UploadRequest{Filename: fileName, Content: string(content)})
		if err != nil {
			if status.Convert(err).Code() == codes.InvalidArgument {
				log.Printf("Empty string provided")
			} else {
				log.Printf("Server error")
			}
			return
		}
		log.Printf("Successfully uploaded " + fileName)
	}

	// Display ascii
	var r *proto.DisplayResponse

	r, err = client.DisplayAscii(ctx, &proto.DisplayRequest{})
	if err != nil {
		if status.Convert(err).Code() == codes.FailedPrecondition {
			log.Printf("Database is empty")
		} else {
			log.Printf("Server error")
		}
		return
	}

	fmt.Printf("\n" + r.GetDisplayAscii())

}
