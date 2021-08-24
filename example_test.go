package moderncrud

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tmc/moderncrud/ent"
)

func Example_Widget() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	widget1, err := client.Widget.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a widget: %v", err)
	}
	fmt.Println(widget1)
	// Output:
	// Widget(id=1)
}
