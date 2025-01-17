package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "postal",
	Short: "postal code",
	Long:  `Batch for postal code of JAPAN POST`,
}

var Quiet bool

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "quiet mode")
}

func main() {
	RootCmd.Execute()
}

func Run() (e error) {
	csvpath, err := Download()
	if err != nil {
		return err
	}

	if len(csvpath) < 1 {
		e = errors.New("Download fail")
		return
	}

	ps, err := ParseFile(csvpath)
	if err != nil {
		e = err
		return
	}

	if err := ps.Insert(); err != nil {
		e = err
		return
	}
	temp := filepath.Dir(csvpath)
	os.RemoveAll(temp)

	return
}
