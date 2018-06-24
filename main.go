package main

import (
    "../ddos/ddos";
    "time";
    "fmt"
)
  func main() {
  	workers := 100000
  	d, err := ddos.New("http://127.0.0.1:3000", workers)
  	if err != nil {
  		panic(err)
  	}
  	d.Run()
    time.Sleep(time.Second)
    d.Stop()
    fmt.Println("DDoS attack server: http://0.0.0.0:3000")
  	// Output: DDoS attack server: http://127.0.0.1:80
  }
