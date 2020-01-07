package models

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
)

var (
	Collection = "art"
	DB         = "ascii"
)

type AsciiArt struct {
	Name string
	Art  string
}

var mongoClient *mongo.Client

// Starts up test database
func StartTestDB() {
	DB = "test"
	Collection = "test"
	err := StartMongoDB()
	if err != nil {
		log.Println("could not start database")
	}
}

// Clears database
func ClearDB() {
	mongoClient.Database(DB).Collection(Collection).Drop(context.TODO())
}

// Starts up the database
func StartMongoDB() error {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("could not ping client: %v", err)
	}

	log.Println("Connected to MongoDB!")

	mongoClient = client

	return nil
}

// Seeds the database with ascii art
func SeedDB() error {

	ClearDB()

	artArray, err := GetAsciiArtFromFile()
	if err != nil {
		return fmt.Errorf("could not get ascii from file: %v", err)
	}

	err = InsertAsciiArtDB(artArray)
	if err != nil {
		return fmt.Errorf("could not insert art to DB: %v", err)
	}

	return nil
}

// Takes in an array of ascii art as a parameter and
// inserts them into the database
func InsertAsciiArtDB(artArray []*AsciiArt) error {

	if len(artArray) == 0 {
		return errors.New("array is empty")
	}

	for _, v := range artArray {

		_, err := mongoClient.Database(DB).Collection(Collection).InsertOne(context.TODO(), v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Selects random ascii art and returns it
func GetRandomArt() (*AsciiArt, error) {
	var temp *AsciiArt
	// Selects random ascii art from database
	cur, err := mongoClient.Database(DB).Collection(Collection).Aggregate(context.TODO(), []bson.M{{"$sample": bson.M{"size": 1}}})
	if err != nil {
		return nil, fmt.Errorf("could not make cursor for DB: %v", err)
	}

	for cur.Next(context.TODO()) {
		var elem AsciiArt
		err := cur.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("could not decode cursor for DB: %v", err)
		}

		temp = &elem
	}

	return temp, nil
}

// Retrieves ascii art from file, stores them in a string array and returns it
func GetAsciiArtFromFile() ([]*AsciiArt, error) {
	// Enters the AsciiArt directory
	files, err := ioutil.ReadDir("AsciiArt")
	if err != nil {
		return nil, fmt.Errorf("could not enter directory: %v", err)
	}

	var artArray []*AsciiArt

	// Iterates through the array of text files and retrieves the data while
	// storing them into a new array of strings
	for _, file := range files {
		content, err := ioutil.ReadFile("AsciiArt/" + file.Name())
		if err != nil {
			return nil, fmt.Errorf("could not take assciiArt from text file: %v", err)
		}

		temp := AsciiArt{
			Name: file.Name(),
			Art:  string(content),
		}

		artArray = append(artArray, &temp)
	}

	return artArray, nil
}

// Uploads ascii art to the database
func UploadAsciiArt(fileName string, content string) error {

	temp := AsciiArt{
		Name: fileName,
		Art:  content,
	}

	_, err := mongoClient.Database(DB).Collection(Collection).InsertOne(context.TODO(), temp)
	if err != nil {
		return fmt.Errorf("could not insert item to database: %v", err)
	}

	return nil
}

// Checks to see if the Database is empty
func IsDatabaseEmpty() (bool, error) {

	count, err := mongoClient.Database(DB).Collection(Collection).CountDocuments(context.TODO(), bson.M{}, nil)
	if err != nil {
		return false, fmt.Errorf("could not count database: %v", err)
	}

	if count == 0 {
		return true, nil
	}

	return false, nil
}

// Closes the connection to the database
func CloseDB() error {

	err := mongoClient.Disconnect(context.TODO())
	if err != nil {
		return fmt.Errorf("could not disconnect database: %v", err)
	}

	log.Println("Connection to MongoDB closed.")
	return nil
}
