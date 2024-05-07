package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

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
		cli.StringFlag{
			Name:  "dirname",
			Usage: "Name of the directory to create",
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
		{
			Name:   "create",
			Usage:  "Create a new directory with a .go file inside",
			Flags:  flags,
			Action: createDirectory,
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

// bonus function
func createDirectory(c *cli.Context) {
	dirName := c.String("dirname") // Obtém o nome do diretório da flag --dirname

	if dirName == "" {
		fmt.Println("Please provide a directory name using --dirname flag.")
		return
	}

	err := os.Mkdir(dirName, 0755) // Cria o diretório
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(dirName, dirName+".go") // Cria o nome do arquivo com extensão .go
	file, err := os.Create(fileName)                  // Cria o arquivo
	if err != nil {
		log.Fatal(err)
	}

	// Conteúdo do arquivo .go
	content := fmt.Sprintf(`package %s

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}`, dirName)

	// Escreve o conteúdo no arquivo
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Directory '%s' and file '%s' created successfully.\n", dirName, fileName)
}
