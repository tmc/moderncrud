package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tmc/moderncrud/ent"
)

func main() {
	client := newConnection()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, world")
		widgets, err := client.Widget.Query().All(r.Context())
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		for _, widget := range widgets {
			fmt.Fprintln(w, widget)
		}
	})
	mux.HandleFunc("/widgets", func(w http.ResponseWriter, r *http.Request) {
		widget := ent.Widget{}
		if r.Method == http.MethodPost {
			json.NewDecoder(r.Body).Decode(&widget)
			defer r.Body.Close()
			err := client.Widget.
				Create().
				SetNote(widget.Note).
				Exec(r.Context())
			if err != nil {
				fmt.Fprintln(w, err)
			} else {
				fmt.Fprintln(w, "ok")
			}
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(0)
	}()
	log.Printf("Handling HTTP requests on %s.", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalln(err)
	}
}

// Open new connection to the database.
func newConnection() *ent.Client {
	// var (
	// 	dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
	// 	dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
	// 	instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	// 	dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
	// )

	// socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	// if !isSet {
	// 	socketDir = "/cloudsql"
	// }

	// dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	db, err := sql.Open("sqlite3", "file:db.sqlite?_fk=1&busy_timeout=1000")
	if err != nil {
		log.Fatal(err)
	}

	// db.Exec("pragma busy_timeout = 5000")
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.SQLite, db)
	c := ent.NewClient(ent.Driver(drv))
	return c
}

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
