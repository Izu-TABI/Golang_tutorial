package greetings

import (
  "errors" 
  "fmt"
)


// Hello returns a greeting for named person.
func Hello(name string) (string, error) {
  // If no name given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message, nil
}
