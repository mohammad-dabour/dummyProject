package main

import (
	"fmt"
	"testing"
)

func TestViewHandler(t *testing.T) {
	n, err := fmt.Println("Hello World")

	if err != nil {
		t.Error("Something wrong with Println!", n)
	}

}


func TestRunServer(t *testing.T) {
	n, err := fmt.Println("Hello World")

	if err != nil {
		t.Error("Something wrong with Println!", n)
	}

}
