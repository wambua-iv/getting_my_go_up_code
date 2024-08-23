package main

import "fmt"

type Account struct {
	User string
	accountType string
}

type AccountNotifier interface {
	Notify() string
}

func(acc *Account) createAccount (user, accType string) bool {
	acc.User = user
	acc.accountType = accType
	return true
}

func (acc Account)Notify () string {
	return fmt.Sprintf("Acc %v is created", acc.User)
}

func Email(acc Account) string {
	return acc.Notify()
}