/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vlab-research/go-tremendous/tremendous"
	"net/http"
	"os"
)

// productsCmd represents the products command
var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		key := os.Getenv("TREMENDOUS_API_KEY")

		base := "https://testflight.tremendous.com"
		client := tremendous.NewClient(http.DefaultClient, base, key)

		country, err := cmd.Flags().GetString("country")
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString("currency")
		if err != nil {
			return err
		}

		products, err := client.Products.List(country, currency)

		if err != nil {
			return err
		}

		fmt.Println(products)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(productsCmd)

	productsCmd.Flags().String("country", "", "country")
	productsCmd.Flags().String("currency", "", "currency")
}
