package greetings

import (
  "errors" 
  "fmt"
  "math/rand"
  "time"
)


// Hello returns a greeting for named person.
func Hello(name string) (string, error) {
  // If no name given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func init() {
  rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
  formats := []string {
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  return formats[rand.Intn(len(formats))]
}
