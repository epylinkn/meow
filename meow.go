package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "strings"
  "regexp"
  "time"
  "github.com/urfave/cli"
  "github.com/gookit/color"
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
        pe := "peppers"
        peppers := append(pizza, pe)
        m := strings.Join(peppers, " ")
        fmt.Println(m)
      },
    },
  }
}

func print_cal(day int) {
  cmd := exec.Command("cal", "-3")
  stdout, err := cmd.Output()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  lines := strings.Split(string(stdout), "\n")
  re_backspace := regexp.MustCompile(`(_\x08\d)+`)

  for i := range lines {
    // NB. This hack removes the calendar backspace hack,
    //     and colors the date instead.
    if re_backspace.MatchString(lines[i]) {
      split := re_backspace.Split(lines[i], -1)
      fmt.Print(split[0])
      color.Error.Print(day)
      fmt.Println(split[1])
      continue
    }

    fmt.Println(lines[i])
  }
}

func main() {
  info()
  commands()

  _, _, day := time.Now().Date()
  print_cal(day)

  error := app.Run(os.Args)
  if error != nil {
    log.Fatal(error)
  }
}
