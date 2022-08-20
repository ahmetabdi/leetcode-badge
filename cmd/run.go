package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/haozibi/leetcode-badge/app"
)

// NewRunCommand Run Web Command
func NewRunCommand() *cobra.Command {

	var opt app.Config

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run Web",
		Run: func(cmd *cobra.Command, args []string) {

			a := app.New(opt)
			err := a.Run()
			if err != nil {
				fmt.Printf("%+v\n", err)
				os.Exit(1)
			}
			fmt.Println("See you")
		},
	}

	flag := cmd.Flags()

	flag.StringVarP(&opt.Address, "address", "", ":8080", "http listen address")
	flag.IntVarP(&opt.DebugAddress, "debug-address", "", 18080, "http debug address, only localhost")
	flag.StringVarP(&opt.SqlitePath, "sqlite-path", "", "./lc.db", "sqlite3 file path")
	flag.BoolVarP(&opt.EnableCron, "enable-cron", "", false, "if enable cron")

	return cmd
}
