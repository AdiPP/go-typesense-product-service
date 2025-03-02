/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package syncprododucts

import (
	"fmt"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/typesense"
	"log"
	"strconv"

	"github.com/AdiPP/go-typesense-product-service/internal/app/config"
	"github.com/AdiPP/go-typesense-product-service/internal/app/database/pgsql"
	"github.com/AdiPP/go-typesense-product-service/internal/app/service"
	"github.com/spf13/cobra"
)

// SyncProductsCmd represents the syncProducts command
var SyncProductsCmd = &cobra.Command{
	Use:   "sync-products [product_ids]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you must provide product ids as an argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("syncProducts called")

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %s", err)
		}

		pgsqlRepo, err := pgsql.NewRepository(cfg)
		if err != nil {
			log.Fatalf("Failed to load pgsql repo: %s", err)
		}

		typesenseRepo, err := typesense.NewRepository(cfg)
		if err != nil {
			log.Fatalf("Failed to load typesense repo: %s", err)
		}

		productIds, err := convertArgsToInt64(args)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		err = service.NewProductSynchronizerService(pgsqlRepo, typesenseRepo).
			SyncBatch(&service.SyncBatchProductsParam{
				ProductIDs: productIds,
			})
		if err != nil {
			log.Printf("failed to sync batch: %s", err)
			return
		}
	},
}

func init() {

}

func convertArgsToInt64(args []string) ([]int64, error) {
	var intArgs []int64
	for _, arg := range args {
		num, err := strconv.ParseInt(arg, 10, 64) // Convert string to int64
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s': %v", arg, err)
		}
		intArgs = append(intArgs, num)
	}
	return intArgs, nil
}
