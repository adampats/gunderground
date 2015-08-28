package main

import (
  "fmt"
  // "github.com/docopt/docopt-go"
  // "github.com/jmcvetta/napping"
  "net/http"
  "io/ioutil"
)

func main() {
//   usage := `gunderground.
//     weather from wunderground.com
// Usage:
//   gunderground get seattle,wa
//   gunderground -v | --version
//
// Options:
//   -h --help     Show this help screen.
//   -v --version  Show version.`
//
//   arguments, _ := docopt.Parse(usage, nil, true, "gunderground 1.0", false)
//  fmt.Println(arguments)
  fmt.Println("hi")

  getWeather()
}

func validateCity() {

}

func getWeather() {
  city := "http://api.wunderground.com/api/ef11d4f7b98252f9/conditions/q/WA/Seattle.json"
  response, err := http.Get(city)
  if err != nil {
    fmt.Printf("%s", err)
  } else {
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Printf("%s\n", string(contents))
  }
}
