package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	var url string
	var upload string

	flag.StringVar(&url, "url", "localhost:4040", "a string var")
	flag.StringVar(&upload, "upload", "null", "a string var")
	flag.Parse()





	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := proto.NewAddServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//upload file
	if upload != "null" {

		_, err = client.UploadAscii(ctx,&proto.UploadRequest{Upload:upload})
		if err != nil {
			log.Fatalf("Could not upload: %v", err)
		}

	}

	//display ascii
	var r *proto.DisplayResponse

	r, err = client.DisplayAscii(ctx, &proto.DisplayRequest{})
	if err != nil {
		log.Fatalf("Could not display: %v", err)
	}

	fmt.Printf("\n" + r.GetSp())

}
