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

	flag.StringVar(&url, "url", "localhost:4040", "a string var")
	flag.Parse()

	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := proto.NewAddServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r *proto.AsciiResponse

	r, err = client.DisplayAscii(ctx, &proto.AsciiRequest{})
	if err != nil {
		log.Fatalf("Could not display: %v", err)
	}

	fmt.Printf("\n" + r.GetSp())

}
