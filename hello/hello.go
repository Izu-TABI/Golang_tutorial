package main

import(
  "fmt"
  "log"

  "example.com/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)


  // get greeting
  message, err := greetings.Hello("Gladys")

  if err != nil {
    log.Fatal(err)
  }



  fmt.Println(message)
}
