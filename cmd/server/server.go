package main

import (
	"github.com/gofiber/fiber/v2"
	bolt "go.etcd.io/bbolt"
	"log"
)

type Server struct {
	db *bolt.DB
	r  *fiber.App
}

func main() {
	s := &Server{
		db: createKVStore(),
		r:  fiber.New(),
	}

	apiV1 := s.r.Group("/api/v1")
	registerApiV1Routes(apiV1)

	log.Fatal(s.r.Listen(":3000"))
}
