package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func init() {
	if os.Getenv("SLEEP") == "true" {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	fmt.Println("Some really critical log you need to see, but won't")
	os.Exit(1)
}
