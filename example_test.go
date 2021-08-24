package moderncrud

import (
	"context"
	"fmt"
	"log"
	"os"

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
	wt, err := client.WidgetType.Create().
		SetName("shiny").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a widgettype: %v", err)
	}
	widget1, err := client.Widget.Create().
		SetType(wt).
		SetNote("this is an example widget").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a widget: %v", err)
	}
	fmt.Println(widget1.ID)
	// Output:
	// 1
}

func Example_DumpSchema() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed writing schema: %v", err)
	}
	// Output:
	// BEGIN;
	// CREATE TABLE `widgets`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `note` varchar(255) NOT NULL, `created_at` datetime NOT NULL, `status` varchar(255) NOT NULL DEFAULT 'draft', `priority` integer NOT NULL DEFAULT 0, `widget_type` integer NULL, FOREIGN KEY(`widget_type`) REFERENCES `widget_types`(`id`) ON DELETE SET NULL);
	// CREATE TABLE `widget_types`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NOT NULL);
	// COMMIT;
}
