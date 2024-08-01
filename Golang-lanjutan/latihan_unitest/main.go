package main

import (
  "fmt"
  "regexp"
  "testing"
)
func IsValidEmail(email string) bool {
  regexPattern := `^[^@]+@[^@]+$`
  match, _ := regexp.MatchString(regexPattern, email)
  return match
}

func TestValid(t *testing.T) {
  tests := []struct {
    email    string
    expected bool
  }{
    {"afrd@gmail.com", true},
    {"afrd@gmail", false},
    {"afrdgmail.com", false},
    {"afrd@.com", false},
    {"afrd@@gmail.com", false}, 
  }

  for _, test := range tests {
    result := IsValidEmail(test.email)
    if result != test.expected {
      t.Errorf("Test failed for email '%s'. Expected: %v, Got: %v", test.email, test.expected, result)
    }
  }
}

func main() {
  // Example usage of IsValidEmail function
  email := "user@example.com"
  fmt.Printf("Is '%s' a valid email address? %t\n", email, IsValidEmail(email))
}
