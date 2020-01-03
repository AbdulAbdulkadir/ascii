package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	StartTestDB()
	ClearDB()
	code := m.Run()
	CloseDB()
	os.Exit(code)
}

func TestInsertAsciiArtDB(t *testing.T) {

	temp1 := AsciiArt{
		Name:  "animal",
		Art: "content",
	}

	temp2 := AsciiArt{
		Name:  "fish",
		Art: "content",
	}

	tempArray := []*AsciiArt{&temp1, &temp2}

	t.Run("should fail when artArray is empty", func(t *testing.T) {
		err := InsertAsciiArtDB([]*AsciiArt{})
		if err == nil {
			t.Fail()
		}
	})
	t.Run("there should not be an error when artArray is not empty", func(t *testing.T) {
		err := InsertAsciiArtDB(tempArray)
		if err != nil {
			t.Fail()
		}

	})
}

func TestUploadAsciiArt(t *testing.T) {

	t.Run("there should not be an error when parameter is not empty", func(t *testing.T) {
		err := UploadAsciiArt("dummyName","something")
		if err != nil {
			t.Fail()
		}
	})
}
