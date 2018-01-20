package main

import (
  "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
  Use: "postal",
  Short: "postal code",
  Long: `Batch for postal code of JAPAN POST`,
}

var Quiet bool

func init(){
  cobra.OnInitialize()
  RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "quiet mode")
}


func main() {
	RootCmd.Execute()
}

