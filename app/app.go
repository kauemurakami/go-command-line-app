package app

import "github.com/urfave/cli"

// return the application command line ready to use
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Commanda Line Applications"    //nome
	app.Usage = "Search IP's and servers name" //utilização
	return app
}
