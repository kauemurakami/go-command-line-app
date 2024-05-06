[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-command-line-app/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-command-line-app/blob/main/README.md)

## Applicação
Aplicação tem duas ações, uma com objetivo de receber um endereço web, como google.com por exemplo, e retornar o IP público do endereço, a segunda tem como objetivo retornar o nome do servidor onde o endereço está hospedado. Vamos usar um pacote externo para isso.  

### Iniciando
Crie a pasta ```go-command-line-app```, no terminal, tendo a pasta criada como raiz, crie um módulo com ```go mod init command-line-app```, e um módulo ```go.mod``` será criado no seu diretório, ao abrir ele você deve ver algo como:  
```go
module command-line-app // nome do module

go 1.22.1 // sua versão do go
```
Agora iremos adicionar o pacote externo que irá nos ajudar nessa aplicação:  
```shell
$ go get github.com/urfave/cli
```
Com isso, ele será instalado e importado no seu módulo ```go.mod```  
```go
...
require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli v1.22.15 // indirect
)
```
Além de criar um arquivo ```go.sum```, não se preocupe com ele por enquanto, mas é um tipo de .lock, ele guarda as informções de histórico das instalações, deduções etc, não é bom mexer nele nem no módulo, recomendavel deixar o ```go``` cuidar disso.<br/><br/>

Agora crie o arquivo ```main.go``` com a configuração inicial:  
```go
package main

func main() {

}
```
Agora vamos criar outro diretório dentro do nosso diretório raiz chamado ```/app```, dentro dele vamos criar o arquivo ```app.go```, que é onde teremos nossa aplicação de linha de comando, a aplicação de linha de comando será feita aqui, vamos criar a estrutura inicial e analisar depois:  
```go
package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
  app := cli.newApp() // tipo do app é um *cli.App como precisamos
	return app
}
```
Aqui estamos importando o package ```cli```, lembrando que referênciamos o package usando a palavra que está após a última barra /.  
Na nossa função estamos referênciando o package ```cli``` com nosso sufixo (*), recuperando um tipo ```App``` com ```*cli.App```, que é um tipo contido internamente dentro do package.<br/>
Com isso precisamos retornar um tipo igual ```*cli.App```, pra isso criamos a variável ```app``` e atribuímos uma função ```cli.newApp()``` que retorna esse tipo.<br/><br/>
Mas antes de retornar o ```app```, temos que configurar algumas coisas:  
```go
func Gerar() *cli.App {
	app := cli.NewApp()                        //generate instance of *cli.App
	app.Name = "Commanda Line Applications"    //nome
	app.Usage = "Search IP's and servers name" //utilização
	return app
}
```
Reparem que nossa função ainda não faz nada, vamos apenas importa-lá em nossa ```main.go```, segue o exemplo abaixo:  
```go
...
  import (
	"command-line-app/app"
	"fmt"
)

func main() {
	application := app.Gerar()
	application.Run(os.Args)
}
```
Aqui recebemos o resultado da função gerar contida em ```app/app.go```, a função ```Run(os.Args)``` recebe o argumento padrão ```os.Args```, serve para que nossos comandos do sistema operacional sejam reconhecidos pela linha de comando.<br/><br/>

*Tratando possíveis erros*  
Repare que nossa função ```.Run(os.Args)``` pode retornar um erro, portanto devemos trata-lo.  
```go
...
func main() {
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
    log.Fatal(erro)
  }
}
```
Aqui recebemos o erro caso erro seja diferente de nulo ou ```nil```, caso isso ocorra faremos um ```log.Fatal(erro)```, diferente do ```print```, ele exibe mais informações sobre o erro, e para nossa aplicação neste caso.  
Caso tenha ficado alguma dúvida sobre a atribuição, se afunção funcionar corretamente ela não possui retorno por isso o erro seria ```nil``` em caso de sucesso, apenas em caso de erro ele teria um valor.  
Outra coisa é que:  
```go
if erro := application.Run(os.Args); erro != nil {
    log.Fatal(erro)
  }
```
É O MESMO QUE
```go
erro := application.Run(os.Args)
if erro != nil {
  log.Fatar(erro)
}
```
Mas podemos fazer isso direto na atribuição.<br/><br/>


*Adicionando busca de IP por host/domínio na função Gerar*  
Em seu arquivo ```app/app.go``` faça as alterações e falaremos após:  
```go
...
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Commanda Line Applications"    // nome
	app.Usage = "Search IP's and servers name" // utilização
	// Add commands use this property, it is slice of the commands
	app.Commands = []cli.Command{
		{
			Name:  "ip", // nome do comando
      //descrição do uso
			Usage: "Search IP's of the address of the internet www.google.com",
      // flags --<flag> <value> que vamos criar/usar
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host", // nome da flag
					Value: "google.com.br", // valor padrão
				},
			},
			Action: searchIps, // ação/função quando comando for executado
		},
	}
	return app
}
```
Podemos ver algumas configurações e funcionalidades adicionadas em nosso ```app```, criamos então a funcionalidade ```ip```, com um parâmetro via ```flag``` chamado ```host```, com um valor padrão, caso nenhum argumento seja passado, no nosso caso o google.  
Agora nossa função ```searchIps```, vamos criar ela fora do escopo da função ```Gerar()```:  
```go
.......
		Action: searchIps, // ação/função quando comando for executado
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host") // recebemos o valor passado com a flag --host <host-value>
  
  // usamos o pacakge net para fazer a busca do ip
	ips, erro := net.LookupIP(host) 
  // caso tudo corra bem receberemos na variáel ips um slice de []net.IP
  // em caso de erro recebemos ele e damos outro log.Fatal(erro)
	if erro != nil {
		log.Fatal(erro)
	}
  
  // vamos percorrer esse slice de ips para recuperarmos cada ip contido em ips
  // e vamos mostrar isso na linha comando com um print comum
	for _,ip := range ips {
		fmt.Println(ip)
	}
}
```
Comentei linha por linha para melhor entendimento, agora vamos rodar esse nosso comando criado.  
```shell
$ go run main.go ip --host amazon.com.br
```
Caso tenha passado o host corretamente você deve ver um ou mais endereços de ip como output, usando ```amazon.com.br``` recebi os seguintes resultados.  
```json
54.239.26.87
52.94.225.243
72.21.203.171
54.239.26.87
52.94.225.243
72.21.203.171
```  
Caso rode apenas ```$go run main.go ip``` receberá o IP do goole que definimos como padrão.<br/><br/>

*Busca por servidor*  
Inicialmente vamos começar alterando a função ```Gerar()```, atribuindo as Flags em uma variável ```flags```  
```go
......
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}
	// Add commands use this property, it is slice of the commands
	app.Commands = []cli.Command{
		//busca por ip
    {
			Name:   "ip", // command name
			Usage:  "Search IP's of the address of the internet www.google.com",
			Flags:  flags,
			Action: searchIps,
		},
    //busca por servidor
    {
			Name:   "servers",
			Usage:  "Search host server name",
			Flags:  flags,
			Action: searchHostServers,
		},
  }
  ......
```
Repare que fazemos os mesmos passos, a única diferença é que usamos uma função diferente do nosso package ```net```  

Para rodar agora:  
```shell
go run main.go servers --host amazon.com.br
```
