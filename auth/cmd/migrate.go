package cmd

import (
	"github.com/open-cloud-initiative/glue/auth/internal/adapters/db"

	"github.com/katallaxie/pkg/dbx"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	RunE: func(cmd *cobra.Command, _ []string) error {
		conn, err := gorm.Open(postgres.Open(cfg.Flags.DatabaseURI), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{},
		})
		if err != nil {
			return err
		}

		store, err := dbx.NewDatabase(conn, db.NewReadTx(), db.NewWriteTx())
		if err != nil {
			return err
		}

		return store.Migrate(
			cmd.Context(),
		)
	},
}
