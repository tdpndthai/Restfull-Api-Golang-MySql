package models

import (
	"database/sql"
	"restfull-api/entities"
)

type ProcModel struct {
	Db *sql.DB
}

func (prodModel ProcModel) FindAll() (proc []entities.Proc, err error) { //method truyền vào procmodel trả về mảng proc
	rows, err := prodModel.Db.Query("select * from proc")
	if err != nil {
		return nil, err
	} else {
		var procs []entities.Proc
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity) //sao chép các cột trong hàng hiện tại vào các giá trị được chỉ định
			if err2 != nil {
				return nil, err2
			} else {
				proc := entities.Proc{
					ID:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				procs = append(procs, proc)
			}
		}
		return procs, nil
	}
}

func (prodModel ProcModel) Search(keyword string) (proc []entities.Proc, err error) { //method truyền vào procmodel trả về mảng proc
	rows, err := prodModel.Db.Query("select * from proc where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var procs []entities.Proc //khởi tạo 1 mảng tên procs chưa struc Proc
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity) //sao chép các cột trong hàng hiện tại vào các giá trị được chỉ định
			if err2 != nil {
				return nil, err2
			} else {
				proc := entities.Proc{
					ID:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				procs = append(procs, proc)
			}
		}
		return procs, nil
	}
}

func (prodModel ProcModel) SearchPrices(min, max float64) (proc []entities.Proc, err error) { //method truyền vào mảng struc Proc trả về mảng proc
	rows, err := prodModel.Db.Query("select * from proc where price >= ? and price <= ?", min, max)
	if err != nil {
		return nil, err
	} else {
		var procs []entities.Proc
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity) //sao chép các cột trong hàng hiện tại vào các giá trị được chỉ định
			if err2 != nil {
				return nil, err2
			} else {
				proc := entities.Proc{
					ID:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				procs = append(procs, proc)
			}
		}
		return procs, nil
	}
}

func (prodModel ProcModel) Create(proc *entities.Proc) (err error) { //method truyền vào struc Proc trả về ProcModel,
	result, err := prodModel.Db.Exec("insert into proc(id,name,price,quantity) value (?,?,?,?)", proc.ID, proc.Name, proc.Price, proc.Quantity)
	if err != nil {
		return err
	} else {
		proc.ID, _ = result.LastInsertId()
		return nil
	}
}

func (prodModel ProcModel) Update(proc *entities.Proc) (int64, error) { //method truyền vào struc Proc trả về ProcModel,
	result, err := prodModel.Db.Exec("update proc set name = ?,price = ? ,quantity = ? where id =? ", proc.Name, proc.Price, proc.Quantity, proc.ID)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}
}

func (prodModel ProcModel) Delete(id int64) (int64, error) { //method truyền vào struc Proc trả về ProcModel,
	result, err := prodModel.Db.Exec("delete from proc where id = ?", id)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}
}
