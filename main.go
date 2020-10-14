package main

import (
  "fmt"
  "plugin"
)

func main() {
  p, err := plugin.Open("plugin/plugin.so")
  if err != nil {
    fmt.Println(err)
  }

  f, _ := p.Lookup("Hello")
  j := f.(func() map[string]interface{})()
  fmt.Println(j)
}
