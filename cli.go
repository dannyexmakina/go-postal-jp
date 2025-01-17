package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run",
	Long:  `run`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Run()
	},
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download",
	Long:  `download`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Cli_Download()
	},
}
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "parse",
	Long:  `parse`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Cli_Parse(args)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  `get`,
}

var getSqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql",
	Long:  `sql`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Cli_Get_SQL()
	},
}

var getJsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json",
	Long:  `json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Cli_Get_JSON()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(downloadCmd)
	//
	RootCmd.AddCommand(parseCmd)
	//
	RootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getSqlCmd)
	getCmd.AddCommand(getJsonCmd)

}

func Cli_Download() (e error) {
	cspath, err := Download()
	if err != nil {
		return err
	}
	fmt.Println(cspath)
	return
}

func Cli_Parse(args []string) (e error) {
	ps, err := ParseFile(args[0])
	if err != nil {
		e = err
		return
	}

	if err := ps.Insert(); err != nil {
		e = err
		return
	}

	return
}

func Cli_Get_SQL() (e error) {

	b, err := GetSQL()
	if err != nil {
		e = err
		return
	}

	fmt.Println(string(b))

	return
}

func Cli_Get_JSON() (e error) {

	b, err := GetJSON()
	if err != nil {
		e = err
		return
	}

	fmt.Println(string(b))

	return
}
