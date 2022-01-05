package domain

type Customer struct {
	UUID        string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	// ById(string) (*Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			UUID:        "e13a6c89-70cf-46ab-976d-55e9e20c85c0",
			Name:        "Irfan",
			City:        "Cimahi",
			Zipcode:     "40535",
			DateofBirth: "1988-11-29",
			Status:      true,
		},
		{
			UUID:        "ad17788f-8a44-4b1a-9f2a-36fac753f0c8",
			Name:        "Didi",
			City:        "Cimahi",
			Zipcode:     "40535",
			DateofBirth: "1996-11-29",
			Status:      true,
		},
	}

	return CustomerRepositoryStub{customers: customers}
}
