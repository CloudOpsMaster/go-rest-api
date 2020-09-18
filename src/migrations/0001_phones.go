package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upPhones, downPhones)
}

func upPhones(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS phones (
  				"id" SERIAL PRIMARY KEY, 
                "user_name" VARCHAR(50),
                "phone" VARCHAR(15));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downPhones(tx *sql.Tx) error {
	query := `DROP TABLE phones;`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
