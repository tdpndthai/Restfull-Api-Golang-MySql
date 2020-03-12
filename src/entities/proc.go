package entities

import "fmt"

type Proc struct {
	ID       int64   `json :"id"`
	Name     string  `json :"name"`
	Price    float64 `json :"price"`
	Quantity int64   `json :"quantity"`
}

func (proc Proc) ToString() string { //method truyền vào Proc,tên Tostring trả về kiểu string
	return fmt.Sprint("id : %d\nname: %s\nprice: %0.1f\nquantity: %d", proc.ID, proc.Name, proc.Price, proc.Quantity)
}
