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

		base := "https://testflight.tremendous.com"
		client := tremendous.NewClient(http.DefaultClient, base, key)

		orders, err := client.Orders.Create(
			"AYXFGRP96E6G",
			[]string{"TBAJH7YLFVS5"},
			10.0,
			"USD",
			"LINK",
			tremendous.Recipient{Name: "Study Participant"},
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

	rootCmd.AddCommand(ordersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ordersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ordersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
