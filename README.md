[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-command-line-app/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-command-line-app/blob/main/README.md)

## Application
The application has two actions, one with the objective of receiving a web address, such as google.com for example, and returning the public IP of the address, the second with the objective of returning the name of the server where the address is hosted. Let's use an external package for this.  

### Get Start
Create the folder ```go-command-line-app```, in the terminal, using the folder created as root, create a module with ```go mod init command-line-app```, and a module ` ``go.mod``` will be created in your directory, when opening it you should see something like:  
```go
module command-line-app // nome do module

go 1.22.1 // sua versão do go
```
Now we will add the external package that will help us with this application:  
```shell
$ go get github.com/urfave/cli
```
With this, it will be installed and imported into your ```go.mod``` module  
```go
...
require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli v1.22.15 // indirect
)
```
In addition to creating a ```go.sum``` file, don't worry about it for now, but it is a type of .lock, it stores information about the history of installations, deductions, etc., it is not good to touch it or the module, it is recommended to let ```go``` take care of it.<br/><br/>

Now create the ```main.go``` file with the initial configuration:  
```go
package main

func main() {

}
```
Now let's create another directory within our root directory called ```/app```, inside it we will create the file ```app.go```, which is where we will have our command line application, the command line will be done here, let's create the initial structure and analyze it later:  
```go
package app

import "github.com/urfave/cli"

func Gerar() *cli.App {
  app := cli.newApp() // tipo do app é um *cli.App como precisamos
	return app
}
```
Here we are importing the ```cli``` package, remembering that we reference the package using the word after the last slash /.  
In our function we are referencing the package ```cli``` with our suffix (*), retrieving a type ```App``` with ```*cli.App```, which is a type contained internally within from the package.<br/>
With this we need to return a type equal to ```*cli.App```, for this we create the variable ```app``` and assign a function ```cli.newApp()``` that returns this type. <br/><br/>
But before returning ```app```, we have to configure a few things:   
```go
func Gerar() *cli.App {
	app := cli.NewApp()                        //generate instance of *cli.App
	app.Name = "Commanda Line Applications"    //nome
	app.Usage = "Search IP's and servers name" //utilização
	return app
}
```
Note that our function still doesn't do anything, we'll just import it into our ```main.go```, follow the example below:  
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
Here we receive the result of the generate function contained in ```app/app.go```, the function ```Run(os.Args)``` receives the default argument ```os.Args```, serves so that our operating system commands are recognized by the command line.<br/><br/>

*Handling possible errors*  
Note that our function ```.Run(os.Args)``` can return an error, so we must handle it.  
```go
...
func main() {
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
    log.Fatal(erro)
  }
}
```
Here we receive the error if error is different from null or ```nil```, if this occurs we will do a ```log.Fatal(error)```, different from ```print```, it displays more information about the error, and for our application in this case.  
If you have any doubts about the assignment, if the function works correctly it does not have a return, so the error would be ```nil``` in case of success, only in case of error would it have a value.
Another thing is that:  
```go
if erro := application.Run(os.Args); erro != nil {
    log.Fatal(erro)
  }
```
IT IS THE SAME AS  
```go
erro := application.Run(os.Args)
if erro != nil {
  log.Fatar(erro)
}
```
But we can do this directly in the assignment.<br/><br/>


*Adding IP search by host/domain in the Generate function*  
In your ```app/app.go``` file, make the changes and we'll talk later:  
```go
...
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Commanda Line Applications"    // name
	app.Usage = "Search IP's and servers name" // use
	// Add commands use this property, it is slice of the commands
	app.Commands = []cli.Command{
		{
			Name:  "ip", // name of the command
      //description of the use
			Usage: "Search IP's of the address of the internet www.google.com",
      // flags --<flag> <value> we using
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host", // name of the flag
					Value: "google.com.br", // default value
				},
			},
			Action: searchIps, // action/function when command is executed
		},
	}
	return app
}
```
We can see some configurations and functionalities added to our ```app```, we then created the ```ip``` functionality, with a parameter via ```flag``` called ```host```, with a default value, if no argument is passed, in our case google.  
Now our ```searchIps``` function, let's create it outside the scope of the ```Generate()``` function:  
```go
.......
		Action: searchIps, // action/function when command is executed
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host") // we receive the value passed with the flag --host <host-value>
  
  // we use pacakge net to search for the ip
	ips, erro := net.LookupIP(host) 
  // If everything goes well, we will receive a slice of []net.IP in the ips variable
  // in case of error we receive it and give another log.Fatal(error)
	if erro != nil {
		log.Fatal(erro)
	}
  
  // let's go through this slice of ips to retrieve each ip contained in ips
  // and we will show this on the command line with a common print
	for _,ip := range ips {
		fmt.Println(ip)
	}
}
```
I commented line by line for better understanding, now let's run our created command.  
```shell
$ go run main.go ip --host amazon.com.br
```
If you passed the host correctly you should see one or more IP addresses as output, using ```amazon.com.br``` I received the following results.  
```json
54.239.26.87
52.94.225.243
72.21.203.171
54.239.26.87
52.94.225.243
72.21.203.171
```  
If you just run ```$go run main.go ip``` you will receive the goole IP that we defined as default.<br/><br/>

*Search by server*  
Initially, let's start by changing the ```Generate()``` function, assigning the Flags in a variable ```flags```
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
		//search by ip
    {
			Name:   "ip", // command name
      // description of the usage
			Usage:  "Search IP's of the address of the internet www.google.com",
			Flags:  flags, // used flags
			Action: searchIps, // function of the action
		},
    //search server
    {
			Name:   "servers", //command name
			Usage:  "Search host server name", // description of the usage
			Flags:  flags, // used flags
			Action: searchHostServers, // function of the action
		},
  }
  ......
```
Note that we do the same steps, the only difference is that we use a different function from our ```net``` package  

To run now:  
```shell
go run main.go servers --host amazon.com.br
```
### Bônus
Function to create a ```directory``` and ```file.go```, first we will add a ```cli.Command```, called ```create```, then we will add a new function in ```app.go```, outside the scope of the ```Generate()``` function, as well as the other methods:  
```go
...
app.Commands = []cli.Command{
		{
			Name:   "ip", // command name
			Usage:  "Search IP's of the address of the internet www.google.com",
			Flags:  flags,
			Action: searchIps,
		},
		...
add>{
			Name:   "create",
			Usage:  "Create a new directory with a .go file inside",
			Flags:  flags,
			Action: createDirectory,
		}, << 
...
// bonus function 
func createDirectory(c *cli.Context) {
	dirName := c.String("dirname") // Gets the directory name from the --dirname flag

	if dirName == "" {
		fmt.Println("Please provide a directory name using --dirname flag.")
		return
	}

	err := os.Mkdir(dirName, 0755) // Create Folder/Directory
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(dirName, dirName+".go") // Creates the file name with .go extension
	file, err := os.Create(fileName)                  // Create file
	if err != nil {
		log.Fatal(err)
	}

	// Add content to file .go
	content := fmt.Sprintf(`package %s

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}`, dirName)

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Directory '%s' and file '%s' created successfully.\n", dirName, fileName)
}
```
Run:  
```shell
$  go run main.go create --dirname <name>
```
The ```.go``` file will have the same name as the ```directory``` set after ```--dirname``` and will come with the basic contents of a ```.go``` file , with the package also following the ```dirname``` nomenclature.