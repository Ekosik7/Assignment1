package Bank

type BankAccount struct {
ID      string
Owner   string
Balance float64
}

func NewBankAccount(id, owner string) *BankAccount {
return &BankAccount{
ID:    id,
Owner: owner,
}
}

func (b *BankAccount) Deposit(amount float64) {
b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) bool {
if amount > b.Balance {
return false
}
b.Balance -= amount
return true
}

func (b *BankAccount) GetBalance() float64 {
return b.Balance
}
