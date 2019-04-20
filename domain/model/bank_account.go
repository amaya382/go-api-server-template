package model

type BankAccount struct {
	AccountNumber string `gorm:"primary_key"`
	AccountHolder string `gorm:"not null"`
	Balance       uint   `gorm:"not null"`
	UpperLimit    uint   `gorm:"not null"`
}

func (self *BankAccount) Deposit(amount uint) error {
	if self.Balance+amount > self.UpperLimit {
		return NewUpperLimitErr(amount, self.Balance, self.UpperLimit)
	}

	self.Balance += amount
	return nil
}

func (self *BankAccount) Withdraw(amount uint) error {
	if self.Balance-amount < 0 {
		return NewInsufficientFundsErr(amount, self.Balance)
	}

	self.Balance -= amount
	return nil
}
