package main

import (
  "fmt"
)

// Consider an application involving different types of payment methods, such as credit card payment, debit card payment, and PayPal payment. Each payment type has unique properties and methods.
type PaymentFactory struct{}
func (p *PaymentFactory) CreatePayment(paymentType string, props map[string]string) (Payment, error) {
  switch paymentType {
  case "creditcard":
    cardNumber, ok := props["CardPayment"]
    if !ok {
      return nil, fmt.Errorf("Credit card di perlukan untuk pemabayaran")
    }
    return &CreditCardPayment{cardNumber: cardNumber}, nil
  case "debitcard":
    cardNumber, ok := props["CardPayment"]
    if !ok {
      return nil, fmt.Errorf("Debit card diperlukan untuk pembayaran")
    }
    return &DebitCardPayment{cardNumber: cardNumber}, nil
  case "paypal":
    email, ok := props["email"]
    if !ok {
      return nil, fmt.Errorf("PayPal payment tidak ada email")
    }
    return &PaypalPayment{email: email}, nil
  default:
    return nil, fmt.Errorf("tidak ada pembayaran: %s", paymentType)
  }
}
type Payment interface {
  processPayment(amount float32) string
}

type CreditCardPayment struct {
  cardNumber string
}

func (c CreditCardPayment) processPayment(amount float32) string {
  return fmt.Sprintf("Processed payment of %f through credit card %s", amount, c.cardNumber)
}

type DebitCardPayment struct {
  cardNumber string
}

func (d DebitCardPayment) processPayment(amount float32) string {
  return fmt.Sprintf("Processed payment of %f through debit card %s", amount, d.cardNumber)
}

type PaypalPayment struct {
  email string
}

func (p PaypalPayment) processPayment(amount float32) string {
  return fmt.Sprintf("Processed payment of %f through PayPal account %s", amount, p.email)
}

func main() {
  creditCardPayment := CreditCardPayment{cardNumber: "1234-5678-9012-3456"}
  fmt.Println(creditCardPayment.processPayment(100.00))

  debitCardPayment := DebitCardPayment{cardNumber: "2345-6789-0123-4567"}
  fmt.Println(debitCardPayment.processPayment(200.00))

  paypalPayment := PaypalPayment{email: "user@example.com"}
  fmt.Println(paypalPayment.processPayment(300.00))
}
