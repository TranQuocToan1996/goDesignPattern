package creationalpatterns

// Product
type BankAccount struct {
	balance              int64
	interestRate         float64
	name                 string
	identificationNumber string
	branch               string
}

// Interface Builder
type Builder interface {
	setBalance(balance int64)
	setInterestRate()
	setName(name string)
	setIdentificationNumber(identificationNumber string)
	setBranch(branch string)
	getBankAccount() BankAccount
}

// Concrete builder
type BuilderNormalAccount struct {
	product BankAccount
}

func NewBuilderBankAccount() *BuilderNormalAccount {
	return &BuilderNormalAccount{}
}

func (b *BuilderNormalAccount) setBalance(balance int64) {
	b.product.balance = balance
}

func (b *BuilderNormalAccount) setInterestRate() {
	b.product.interestRate = getInterestRate()
}

func (b *BuilderNormalAccount) setName(name string) {
	b.product.name = name
}

func (b *BuilderNormalAccount) setIdentificationNumber(identificationNumber string) {
	b.product.identificationNumber = identificationNumber
}

func (b *BuilderNormalAccount) setBranch(branch string) {
	b.product.branch = branch
}

func (b *BuilderNormalAccount) getBankAccount() BankAccount {
	return BankAccount{
		balance:              b.product.balance,
		interestRate:         b.product.interestRate,
		name:                 b.product.name,
		identificationNumber: b.product.identificationNumber,
		branch:               b.product.branch,
	}
}

// Director

type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) build(amount int64, branch, id, name string) BankAccount {
	d.builder.setBalance(amount)
	d.builder.setBranch(branch)
	d.builder.setIdentificationNumber(id)
	d.builder.setInterestRate()
	d.builder.setName(name)
	return d.builder.getBankAccount()
}

func getInterestRate() float64 {
	return 0.0641
}
