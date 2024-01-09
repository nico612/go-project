package data

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nico612/go-project/example/ent/internal/data/ent"
	"github.com/nico612/go-project/example/ent/internal/data/ent/migrate"

	"log"
)

type Data struct {
	db *ent.Client
}

func NewData(db *ent.Client) *Data {
	return &Data{db: db}
}

func (d *Data) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}
	return nil
}

func NewEntClient(driverName, datasourceName string) *ent.Client {

	client, err := ent.Open(
		driverName,
		datasourceName,
		ent.Debug(),
	)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	// Run the auto migration tool.
	if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
