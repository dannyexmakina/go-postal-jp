package main

import (
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert",
	Long:  `convert`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Convert()
	},
}

func init() {
	RootCmd.AddCommand(convertCmd)

}
