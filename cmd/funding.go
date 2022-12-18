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

// fundingCmd represents the funding command
var fundingCmd = &cobra.Command{
	Use:   "funding",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("funding called")
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		key := os.Getenv("TREMENDOUS_API_KEY")

		base := "https://testflight.tremendous.com"
		client := tremendous.NewClient(http.DefaultClient, base, key)

		fundingSources, err := client.Funding.List()

		if err != nil {
			return err
		}

		fmt.Println(fundingSources)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fundingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fundingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fundingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
