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

var ordersCmd = &cobra.Command{
	Use:   "orders",
	Short: "A brief description of your command",
	Long:  "blahlbah",
}

var listCmd = &cobra.Command{
	Use:   "list",
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

		orders, err := client.Orders.List()

		if err != nil {
			return err
		}

		fmt.Println(orders)

		return nil
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new order",
	RunE: func(cmd *cobra.Command, args []string) error {

		key := os.Getenv("TREMENDOUS_API_KEY")
		base := os.Getenv("TREMENDOUS_BASE_URL")

		client := tremendous.NewClient(http.DefaultClient, base, key)

		funding, err := cmd.Flags().GetString("funding")
		if err != nil {
			return err
		}

		products, err := cmd.Flags().GetStringSlice("products")
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString("currency")
		if err != nil {
			return err
		}

		delivery, err := cmd.Flags().GetString("delivery")
		if err != nil {
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		orders, err := client.Orders.Create(
			funding,
			products,
			amount,
			currency,
			delivery,

			// TODO: add optional email/phone
			tremendous.Recipient{Name: name},
		)

		if err != nil {
			return err
		}

		fmt.Println(orders)

		return nil
	},
}

func init() {
	ordersCmd.AddCommand(listCmd)
	ordersCmd.AddCommand(createCmd)

	createCmd.Flags().String("funding", "", "funding")
	createCmd.Flags().StringSlice("products", []string{}, "products")
	createCmd.Flags().Float64("amount", 0.0, "amount")
	createCmd.Flags().String("currency", "", "currency")
	createCmd.Flags().String("delivery", "", "delivery")
	createCmd.Flags().String("name", "", "name")

	rootCmd.AddCommand(ordersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ordersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ordersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
