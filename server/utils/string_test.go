package utils

import (
	"log"
	"testing"
)

func TestGetIdAndName(t *testing.T) {
	str := "READY:XYZ:JOE"
	id, name := GetIdAndName(str)
	if id != "XYZ" || name != "JOE" {
		log.Println("id =", id)
		log.Println("name =", name)
		t.Error("Incorrect extracted values")
	}
}

func TestGetId(t *testing.T) {
	str := "READY:XYZ:JOE"
	str2 := "GETPERSON:XYZ"

	id := GetId(str)
	if id != "XYZ" {
		log.Println("id =", id)
		t.Error("Incorrect extracted value")
	}

	id = GetId(str2)
	if id != "XYZ" {
		log.Println("id =", id)
		t.Error("Incorrect extracted value")
	}

}
