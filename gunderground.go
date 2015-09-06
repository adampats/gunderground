/*
* gunderground
*
* Usage:
*  wunderground dallas,tx
*
* Your personal wunderground.com api key must be stored locally in
* the following file:
*  wunderground.key
*
*/

package main

import (
  "fmt"
  "flag"
  "os"
  "strings"
  "net/http"
  "io/ioutil"
)

var (
  apiKey string
)

func main() {
  if len(os.Args) > 1 {
    location := parseLocation(os.Args[1])
    apiKey := readConfig()
    getWeather(apiKey,location)
  } else {
    usage()
  }
}

func usage() {
  fmt.Fprintf(os.Stderr, "Usage: %s dallas,tx\n", os.Args[0])
  flag.PrintDefaults()
  os.Exit(2)
}

func parseLocation(location string) []string {
  l := strings.Split(location, ",")
  return l
}

func readConfig() string {
  var s string
  if conf, err := ioutil.ReadFile("wunderground.key"); err == nil {
    s := string(conf)
    return strings.TrimSpace(s)
  } else {
    fmt.Println("wunderground.key not found!")
    os.Exit(0)
  }
  return s
}

func getWeather(key string, location []string) {
  urlpath := "http://api.wunderground.com/api/" + key + "/conditions/q/" +
  location[1] + "/" + location[0] + ".json"
  response, err := http.Get(urlpath)
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
