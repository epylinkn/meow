package main

import (
  "fmt"
  "io/ioutil"
  _ "log"
  "net/http"
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

    fmt.Printf("%s\n", lines[i])
  }
}

func print_ical_buddy() {
  out, err := exec.Command("icalBuddy", "eventsToday").Output()
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Printf("%s\n", out)
}

func print_wttr() {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://wttr.in?1nqF", nil)
  if err != nil {
    fmt.Printf("%s", err)
    return
  }

  req.Header.Set("User-Agent", "curl/7.60.0")
  res, err := client.Do(req)
  if err != nil {
    fmt.Printf("%s", err)
    return
  }

  defer res.Body.Close()
  contents, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("%s", err)
    return
  }

  fmt.Printf("%s\n", string(contents))
}

func main() {
  // info()
  // commands()

  _, _, day := time.Now().Date()
  print_wttr()
  print_cal(day)
  print_ical_buddy()

  // error := app.Run(os.Args)
  // if error != nil {
  //   log.Fatal(error)
  // }
}
