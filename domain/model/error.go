package model

import "fmt"

type UpperLimitErr struct {
	Request    uint
	Balance    uint
	UpperLimit uint
}

func NewUpperLimitErr(request uint, balance uint, upperLimit uint) *UpperLimitErr {
	return &UpperLimitErr{request, balance, upperLimit}
}

func (self *UpperLimitErr) Error() string {
	return fmt.Sprintf("%d requested, upper limit is %d")
}

type InsufficientFundsErr struct {
	Request uint
	Balance uint
}

func NewInsufficientFundsErr(request uint, balance uint) *InsufficientFundsErr {
	return &InsufficientFundsErr{request, balance}
}

func (self *InsufficientFundsErr) Error() string {
	return fmt.Sprintf("%d requested, balance is only %d")
}
