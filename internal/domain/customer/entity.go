package customer

type Customer struct {
	Id int32
}

func NewCustomer(id int32) Customer {
	return Customer{Id: id}
}
