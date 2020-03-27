package main

import (
  "fmt"
  "log"
  "os"
  "strings"
  "github.com/urfave/cli"
)

var app = cli.NewApp()
var pizza = []string{"enjoy your pizza with..."}

func info() {
  app.Name = "meow"
  app.Usage = "A calendar preview CLI written in Golang"
  app.Author = "epylinkn"
  app.Version = "0.1.0"
}

func commands() {
  app.Commands = []cli.Command{
    {
      Name: "peppers",
      Aliases: []string{"p"},
      Usage: "Add peppers to your pizza",
      Action: func(c *cli.Context) {
        pe:= "peppers"
        peppers := append(pizza, pe)
        m := strings.Join(peppers, " ")
        fmt.Println(m)
      },
    },
  }
}

func main() {
  info()
  commands()

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
