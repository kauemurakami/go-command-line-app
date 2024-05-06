package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// return the application command line ready to use
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Commanda Line Applications"    //nome
	app.Usage = "Search IP's and servers name" //utilização
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}
	// Add commands use this property, it is slice of the commands
	app.Commands = []cli.Command{
		{
			Name:   "ip", // command name
			Usage:  "Search IP's of the address of the internet www.google.com",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search host server name",
			Flags:  flags,
			Action: searchHostServers,
		},
	}
	return app
}

func searchHostServers(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host) //NS = Name Server

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
