package main
  
import (
        "fmt"
        "testing"
)

// This is not real test just to let this demo pass!
func TestSave(t *testing.T) {
        n, err := fmt.Println("Hello World")

        if err != nil {
                t.Error("Something wrong with Println!", n)
        }

}
func TestLoadPage(t *testing.T) {
        n, err := fmt.Println("Hello World")

        if err != nil {
                t.Error("Something wrong with Println!", n)
        }

}

func TestMain(t *testing.T) { 
        n, err := fmt.Println("Hello World") 
         
        if err != nil {
                t.Error("Something wrong with Println!", n)
        }

}
