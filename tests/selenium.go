package main

import (
  "log"
  "os"
  "bitbucket.org/tebeka/selenium"
)

func main() {
  caps := selenium.Capabilities{"browserName": "firefox"}
  wd, err := selenium.NewRemote(caps, "")

  if err != nil {
    log.Println("Failed to start selenium:", err.Error())
    os.Exit(1)
  }

  defer wd.Quit()

  wd.Get("http://play.golang.org/?simple=1")
}
