package main

import (
  "github.com/spf13/cobra"

)

var downloadCmd = &cobra.Command{
  Use: "download",
  Short: "download",
  Long: `download zip from JAPAN POST`,
  RunE: func(cmd *cobra.Command, args []string ) error {
    name, e := Download()
    if e != nil {
      return e
    }
    _,ze := Unzip(name)
    if ze != nil {
      return ze
    }
    return nil
  },
}

var convertCmd = &cobra.Command{
  Use: "convert",
  Short: "convert",
  Long: `convert`,
  Run: func(cmd *cobra.Command, args []string ){
  },
}

var convertSqlCmd = &cobra.Command{
  Use: "sql",
  Short: "convert csv to sql",
  Long: `convert csv to sql`,
  RunE: func(cmd *cobra.Command, args []string ) error {
   return Convert(args)
  },
}

var updateCmd = &cobra.Command{
  Use: "update",
  Short: "update",
  Long: `parse, update and/or insert`,
  RunE: func(cmd *cobra.Command, args []string ) error {
    return Update()
  },
}

var updateAutoCmd = &cobra.Command{
  Use: "auto",
  Short: "auto",
  Long: `dl, parse, update and/or insert`,
  RunE: func(cmd *cobra.Command, args []string ) error {
    return UpdateAuto()
  },
}



func init(){
  RootCmd.AddCommand(downloadCmd)

  RootCmd.AddCommand(convertCmd)
  convertCmd.AddCommand(convertSqlCmd)
  convertSqlCmd.PersistentFlags().StringVarP(&config.Input, "input", "i", "", "input file")
  convertSqlCmd.PersistentFlags().StringVarP(&config.Output, "output", "o", "", "output file or directory ex. /var/www/batch/file")

  RootCmd.AddCommand(updateCmd)
  updateCmd.PersistentFlags().StringVarP(&config.Input, "input", "i", "", "input file")

  updateCmd.AddCommand(updateAutoCmd)
}
