package main

import(
  "fmt"

  "example.com/greetings"
)

func main() {
  // get greeting
  message := greetings.Hello("Gladys")
  fmt.Println(message)
}
