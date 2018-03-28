package main
  
import (
        "fmt"
        "testing"
)

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