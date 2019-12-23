package main

import (
	"context"
	"fmt"
	"github.com/AbdulAbdulkadir/ascii/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
)

// You will be using this Trainer type later in the program
type AsciiArt struct {
	Id  int
	Art string
}

type server struct{}

var mongoClient *mongo.Client

func main() {

	mongoClient = startMongoDB()

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

	log.Printf("")

	clearDB(*mongoClient)

	collection := mongoClient.Database("art").Collection("ascii")

	artArray := getAsciiArtFromFile()

	insertAsciiArtDB(artArray, collection)

	results := retrieveAsciiArtArray(collection)

	//	return nil, status.Error(codes.AlreadyExists, "there is already an ascii called cat")

	randIndex := int64(rand.Intn(len(results)))

	selectIndex := selectAsciiArt(randIndex, results)

	return &proto.AsciiResponse{Sp: selectIndex}, nil
}

func startMongoDB() *mongo.Client {
	// Set client options

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func getAsciiArtFromFile() []string {
	//enters the AsciiArt directory
	files, err := ioutil.ReadDir("AsciiArt")
	if err != nil {
		log.Fatal(err)
	}

	var artArray []string
	//iterates through the array of text files and retrieves the data while
	//storing them into a new array of strings
	for _, file := range files {
		//Take asciiArt from text file
		content, err := ioutil.ReadFile("AsciiArt/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		artArray = append(artArray, string(content))
	}

	return artArray
}

func insertAsciiArtDB(artArray []string, collection *mongo.Collection) {

	//iterates through the artArray and inserts into the database
	for i, v := range artArray {

		temp := AsciiArt{
			Id:  i,
			Art: v,
		}

		_, err := collection.InsertOne(context.TODO(), temp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func clearDB(c mongo.Client) {
	//Clears out the database
	collection1 := c.Database("art").Collection("ascii")

	collection1.Drop(context.TODO())
}

func retrieveAsciiArtArray(collection *mongo.Collection) []*AsciiArt {
	// Pass options to the Find method
	findOptions := options.Find()

	// an array which stores the decoded documents
	var results []*AsciiArt

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value which a single document can be decoded
		var elem AsciiArt
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}

func closeDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func selectAsciiArt(sq int64, results []*AsciiArt) string {
	var sp string
	for i, v := range results {

		if int64(i) == (sq) {
			sp = v.Art
		}
	}
	return sp
}
