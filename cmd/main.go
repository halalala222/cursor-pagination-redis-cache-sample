package main

import (
	"github.com/halalala222/cursor-pagination-redis-cache-sample/internal/api/sample"
	"log"
)

func main() {
	engine := sample.Run()
	if err := engine.Run(":8088"); err != nil {
		log.Fatal(err)
	}
}
