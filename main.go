package main

import (
    "../ddos/ddos";
    "time";
    "fmt";
)
  func main() {
  	workers := 100
  	d, err := ddos.New("http://0.0.0.0:3000", workers)
  	if err != nil {
  		panic(err)
  	}
  	d.Run()
    time.Sleep(time.Second)
    d.Stop()
    fmt.Println("DDoS attack server: http://0.0.0.0:3000")
  }
