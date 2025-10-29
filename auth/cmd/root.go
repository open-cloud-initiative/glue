package cmd

import (
	"context"
	"fmt"

	"github.com/open-cloud-initiative/glue/auth/internal/adapters/db"
	"github.com/open-cloud-initiative/glue/auth/internal/config"
	"github.com/open-cloud-initiative/glue/auth/internal/controllers/saml"

	"github.com/gofiber/fiber/v3"
	logger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/katallaxie/pkg/dbx"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var cfg = config.New()

const versionFmt = "%s (%s %s)"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func Init() error {
	ctx := context.Background()

	err := cfg.InitDefaultConfig()
	if err != nil {
		return err
	}

	RootCmd.AddCommand(Migrate)

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true

	err = RootCmd.ExecuteContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

var RootCmd = &cobra.Command{
	Use:   "tags",
	Short: "tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), args...)
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(ctx context.Context, _ ...string) error {
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

	mc := saml.NewMetadataController(store)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())

	err = app.Listen(cfg.Flags.Addr)
	if err != nil {
		return err
	}

	return nil
}
