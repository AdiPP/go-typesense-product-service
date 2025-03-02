/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package migratetypesenseschemas

import (
	"fmt"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"github.com/spf13/cobra"
)

// MigrateTypesenseSchemasCmd represents the migrateTypesenseSchemas command
var MigrateTypesenseSchemasCmd = &cobra.Command{
	Use:   "migrate-typesense-schemas",
	Short: "Migrates and initializes Typesense collection schemas",
	Long: `The migrateTypesenseSchemas command is responsible for creating or updating
Typesense collections based on predefined schemas. 

This command is useful for ensuring that the database schema in Typesense
is aligned with the application's requirements. It can be run during
deployment or whenever schema updates are needed.

This command interacts with the Typesense API and requires the Typesense
server to be running. Ensure that the API key and server URL are correctly
configured before executing this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrateTypesenseSchemas called")

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		migrator, err := typesense.NewMigrator(cfg)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = migrator.Migrate()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {

}
