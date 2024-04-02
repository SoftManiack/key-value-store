package main

import "testing"

func TestStore(t *testing.T) {
	err := Put("index", "1")

	if err != nil {
		t.Errorf("put error")
	}

	indexValue, _ := Get("index")

	if indexValue != "1" {
		t.Errorf("put error")
	}

	if err = Delete("index"); err != nil {
		t.Errorf("delete error")
	}

}

//go test ./... -cover
