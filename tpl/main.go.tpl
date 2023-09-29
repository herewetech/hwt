/*
 * Copyright (C) ###__PROJ_AUTHOR__### - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file main.go
 * @package main
 * @author ###__PROJ_AUTHOR__###
 * @since ###__TODAY__###
 */

package main

import (
	"os"

	"###__PROJ_NAME__###/handler"
	"###__PROJ_NAME__###/runtime"

	"github.com/urfave/cli/v2"
)

func actionServe(c *cli.Context) error {
	handler.InitMisc()

	return runtime.Serve()
}

func actionInitdb(c *cli.Context) error {
	// We do not initialize database here
	return nil
}

// Portal

// @title ###__PROJ_NAME__### API
// @version 0.0.1
// @description ###__PROJ_NAME__### API
// @contact.name ###__PROJ_AUTHOR__###

// @host localhost
// @BasePath /
func main() {
	runtime.LoadConfig()
	runtime.InitLogger()
	runtime.InitServer()
	runtime.InitNats()
	runtime.InitDB()

	app := &cli.App{
		Name: runtime.AppName,
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Run service",
				Action: actionServe,
			},
			{
				Name:   "initdb",
				Usage:  "Initialize database tables",
				Action: actionInitdb,
			},
		},
		DefaultCommand: "serve",
	}

	// Startup app
	if err := app.Run(os.Args); err != nil {
		runtime.Logger.Fatal(err)
	}
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
