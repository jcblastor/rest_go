package main

import (
	"github.com/jcblastor/rest_go/pkg/storage"
)

func main() {
	storage.NewPostgresDB()
}
