package main

import (
	"github.com/gofiber/fiber/v2"
	bolt "go.etcd.io/bbolt"
	"log"
)

func main() {
	app := fiber.New()

	apiV1 := app.Group("/api/v1")
	registerApiV1Routes(apiV1)

	log.Fatal(app.Listen(":3000"))
}

func createKVStore() (*bolt.DB, error) {
	dbm, err := bolt.Open(o.path, o.mode, nil)
	if err != nil {
		panic(err)
	}
	err = dbm.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(o.bucket))
		if err != nil {
			return err
		}
		return nil
	})
}
