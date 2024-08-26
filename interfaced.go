package main

import "fmt"

type Account struct {
	User        string
	accountType string
	AccountNotifier
}

type AccountNotifier interface {
	Notify() string
}

func (acc *Account) createAccount(user, accType string) bool {
	acc.User = user
	acc.accountType = accType
	return true
}

func (acc Account) Notify() string {
	return fmt.Sprintf("Acc %v is created", acc.User)
}

func Email(acc Account) string {
	return acc.AccountNotifier.Notify()
}


/*

slighted modification 
it kinder worked ?!
type Account struct {
	User        string
	accountType string
	AN          AccountNotifier
}

type AccountNotifier interface {
	Notify() string
}

func (acc *Account) createAccount(user, accType string) bool {
	acc.User = user
	acc.accountType = accType
	return true
}

func (acc *Account) Notify() string {
	return fmt.Sprintf("Acc %v is created", acc.User)
}

var a Account
a.createAccount("Mitchie", "Insta")
a.Notify()
*/

//Geometric interface

type Shape interface {
	Area () float64
	Perimeter () int
}

func PrintShapeOut(s Shape) {
	fmt.Printf("Area is %v", s.Area())
	fmt.Printf("Perimeter is %v", s.Perimeter())
}

type Triangle struct {
	length int
}

func (t Triangle) Area () float64 {
	return 0.5 * float64(t.length) * float64(t.length)
}

func (t Triangle) Perimeter () int {
	return t.length * 3
}

type Circle struct {
	diameter int
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.diameter) * 0.5 * float64(c.diameter) * 0.5
}

func (c Circle) Perimeter () int {
	return 3 * c.diameter
}
