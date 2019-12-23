package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
)

const (
	Collection = "art"
	DB = "ascii"
)

type AsciiArt struct {
	Id  int
	Art string
}

var mongoClient *mongo.Client

func StartMongoDB() {
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

	mongoClient = client
}

func SeedDB() error {

	mongoClient.Database(DB).Collection(Collection).Drop(context.TODO())

	artArray := GetAsciiArtFromFile()

	if err := InsertAsciiArtDB(artArray); err != nil {
		return err
	}

	return nil
}

func InsertAsciiArtDB(artArray []string) error {

	//iterates through the artArray and inserts into the database
	for i, v := range artArray {

		temp := AsciiArt{
			Id:  i,
			Art: v,
		}

		_, err := mongoClient.Database(DB).Collection(Collection).InsertOne(context.TODO(), temp)
		if err != nil {
			return err
		}
	}
	return nil
}

func RetrieveAsciiArtArray() []*AsciiArt {
	// Pass options to the Find method
	findOptions := options.Find()

	// an array which stores the decoded documents
	var results []*AsciiArt

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := mongoClient.Database(DB).Collection(Collection).Find(context.TODO(), bson.D{{}}, findOptions)
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

func CloseDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GetAsciiArtFromFile() []string {
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

func SelectAsciiArt(sq int64, results []*AsciiArt) string {
	var sp string
	for i, v := range results {

		if int64(i) == (sq) {
			sp = v.Art
		}
	}
	return sp
}