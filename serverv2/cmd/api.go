package cmd

import (
	"context"
	"github.com/gx/gx-ytdl/serverv2/config"
	"github.com/gx/gx-ytdl/serverv2/pkg/_const"
	"github.com/gx/gx-ytdl/serverv2/pkg/database"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Run() {
	ctx := context.Background()

	logger.Init(ctx)

	dbCfg, err := config.GetDatabase(_const.DatabaseFileName)
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(dbCfg)
	if err != nil {
		panic(err)
	}

	apiCfg, err := config.GetAPI(_const.APIFileName)
	if err != nil {
		panic(err)
	}

	select {}
}
