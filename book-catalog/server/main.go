package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"book-catalog/ent"
	"book-catalog/graph"

	"book-catalog/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open("mysql", "root@tcp(localhost:3306)/ent_book_catalog?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	srv := handler.NewDefaultServer(graph.NewSchema(logger, client))
	http.Handle("/",
		playground.Handler("Book Catalog", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
