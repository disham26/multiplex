package model

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	SimDBInit()

}

func TestGetSingleObjectForShow1(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("1")
	log.Println(object.ShowID)

	if object.ShowID != "1" {
		t.Error("Got some other show object ")
	}
}
func TestGetSingleObjectForShow2(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("2")
	log.Println(object.ShowID)

	if object.ShowID != "1" {
		t.Error("Got some other show object ")
	}
}
func TestGetSingleObjectForShow3(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("3")
	log.Println(object.ShowID)

	if object.ShowID != "1" {
		t.Error("Got some other show object ")
	}
}

func TestGetSingleObjectForShowOther(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("5")

	if object.ShowID != "" {
		t.Error("Got some other show object ")
	}
}

func TestUpdateObject1(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("1")
	newObject := UpdateObject(object, 0, 0)
	if !newObject.Rows[0].Seats[0].Taken {
		t.Error("Had to be true")
	}
}

func TestUpdateObject2(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("2")
	newObject := UpdateObject(object, 0, 0)
	if !newObject.Rows[0].Seats[0].Taken {
		t.Error("Had to be true")
	}
}

func TestUpdateObject3(t *testing.T) {
	SimDBInit()
	object := GetSingleObject("3")
	newObject := UpdateObject(object, 0, 0)
	if !newObject.Rows[0].Seats[0].Taken {
		t.Error("Had to be true")
	}
}

func TestDeleteObject(t *testing.T) {
	SimDBInit()
	if !DeleteObject() {
		t.Error("Should have been deleted")
	}
}
