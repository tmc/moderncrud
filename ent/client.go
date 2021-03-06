// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/moderncrud/ent/migrate"

	"github.com/tmc/moderncrud/ent/widget"
	"github.com/tmc/moderncrud/ent/widgettype"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Widget is the client for interacting with the Widget builders.
	Widget *WidgetClient
	// WidgetType is the client for interacting with the WidgetType builders.
	WidgetType *WidgetTypeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Widget = NewWidgetClient(c.config)
	c.WidgetType = NewWidgetTypeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Widget:     NewWidgetClient(cfg),
		WidgetType: NewWidgetTypeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:     cfg,
		Widget:     NewWidgetClient(cfg),
		WidgetType: NewWidgetTypeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Widget.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Widget.Use(hooks...)
	c.WidgetType.Use(hooks...)
}

// WidgetClient is a client for the Widget schema.
type WidgetClient struct {
	config
}

// NewWidgetClient returns a client for the Widget from the given config.
func NewWidgetClient(c config) *WidgetClient {
	return &WidgetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `widget.Hooks(f(g(h())))`.
func (c *WidgetClient) Use(hooks ...Hook) {
	c.hooks.Widget = append(c.hooks.Widget, hooks...)
}

// Create returns a create builder for Widget.
func (c *WidgetClient) Create() *WidgetCreate {
	mutation := newWidgetMutation(c.config, OpCreate)
	return &WidgetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Widget entities.
func (c *WidgetClient) CreateBulk(builders ...*WidgetCreate) *WidgetCreateBulk {
	return &WidgetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Widget.
func (c *WidgetClient) Update() *WidgetUpdate {
	mutation := newWidgetMutation(c.config, OpUpdate)
	return &WidgetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WidgetClient) UpdateOne(w *Widget) *WidgetUpdateOne {
	mutation := newWidgetMutation(c.config, OpUpdateOne, withWidget(w))
	return &WidgetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WidgetClient) UpdateOneID(id int) *WidgetUpdateOne {
	mutation := newWidgetMutation(c.config, OpUpdateOne, withWidgetID(id))
	return &WidgetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Widget.
func (c *WidgetClient) Delete() *WidgetDelete {
	mutation := newWidgetMutation(c.config, OpDelete)
	return &WidgetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WidgetClient) DeleteOne(w *Widget) *WidgetDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WidgetClient) DeleteOneID(id int) *WidgetDeleteOne {
	builder := c.Delete().Where(widget.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WidgetDeleteOne{builder}
}

// Query returns a query builder for Widget.
func (c *WidgetClient) Query() *WidgetQuery {
	return &WidgetQuery{
		config: c.config,
	}
}

// Get returns a Widget entity by its id.
func (c *WidgetClient) Get(ctx context.Context, id int) (*Widget, error) {
	return c.Query().Where(widget.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WidgetClient) GetX(ctx context.Context, id int) *Widget {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryType queries the type edge of a Widget.
func (c *WidgetClient) QueryType(w *Widget) *WidgetTypeQuery {
	query := &WidgetTypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(widget.Table, widget.FieldID, id),
			sqlgraph.To(widgettype.Table, widgettype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, widget.TypeTable, widget.TypeColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WidgetClient) Hooks() []Hook {
	return c.hooks.Widget
}

// WidgetTypeClient is a client for the WidgetType schema.
type WidgetTypeClient struct {
	config
}

// NewWidgetTypeClient returns a client for the WidgetType from the given config.
func NewWidgetTypeClient(c config) *WidgetTypeClient {
	return &WidgetTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `widgettype.Hooks(f(g(h())))`.
func (c *WidgetTypeClient) Use(hooks ...Hook) {
	c.hooks.WidgetType = append(c.hooks.WidgetType, hooks...)
}

// Create returns a create builder for WidgetType.
func (c *WidgetTypeClient) Create() *WidgetTypeCreate {
	mutation := newWidgetTypeMutation(c.config, OpCreate)
	return &WidgetTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WidgetType entities.
func (c *WidgetTypeClient) CreateBulk(builders ...*WidgetTypeCreate) *WidgetTypeCreateBulk {
	return &WidgetTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WidgetType.
func (c *WidgetTypeClient) Update() *WidgetTypeUpdate {
	mutation := newWidgetTypeMutation(c.config, OpUpdate)
	return &WidgetTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WidgetTypeClient) UpdateOne(wt *WidgetType) *WidgetTypeUpdateOne {
	mutation := newWidgetTypeMutation(c.config, OpUpdateOne, withWidgetType(wt))
	return &WidgetTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WidgetTypeClient) UpdateOneID(id int) *WidgetTypeUpdateOne {
	mutation := newWidgetTypeMutation(c.config, OpUpdateOne, withWidgetTypeID(id))
	return &WidgetTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WidgetType.
func (c *WidgetTypeClient) Delete() *WidgetTypeDelete {
	mutation := newWidgetTypeMutation(c.config, OpDelete)
	return &WidgetTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WidgetTypeClient) DeleteOne(wt *WidgetType) *WidgetTypeDeleteOne {
	return c.DeleteOneID(wt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WidgetTypeClient) DeleteOneID(id int) *WidgetTypeDeleteOne {
	builder := c.Delete().Where(widgettype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WidgetTypeDeleteOne{builder}
}

// Query returns a query builder for WidgetType.
func (c *WidgetTypeClient) Query() *WidgetTypeQuery {
	return &WidgetTypeQuery{
		config: c.config,
	}
}

// Get returns a WidgetType entity by its id.
func (c *WidgetTypeClient) Get(ctx context.Context, id int) (*WidgetType, error) {
	return c.Query().Where(widgettype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WidgetTypeClient) GetX(ctx context.Context, id int) *WidgetType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WidgetTypeClient) Hooks() []Hook {
	return c.hooks.WidgetType
}
