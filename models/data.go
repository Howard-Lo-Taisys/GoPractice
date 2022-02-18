package models

import (
 "log"
 db "crawler2/database"
)

type Attributes struct {
	Id			int		`json:"id"`
	Rank		int     `json:"rank"`
	Address		string  `json:"holder"`
	Amount		string  `json:"amount"`
	Percentage	float64 `json:"percent"`
}

func (d *Attributes) GetData() (datas []Attributes, err error) {
	datas = make([]Attributes, 0)
	rows, err := db.SqlDB.Query("SELECT id, numb, address, amount, percentage FROM users")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var data Attributes
		rows.Scan(&data.Id, &data.Rank, &data.Address, &data.Amount, &data.Percentage)
		datas = append(datas, data)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}


func (d *Attributes) AddData() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO users(numb, address, amount, percentage) VALUES(?, ?, ?, ?)", d.Rank, d.Address, d.Amount, d.Percentage)
	if err != nil {
	 return
	}
	id, err = rs.LastInsertId()
	return
}

func (d *Attributes) DeleteData() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM users WHERE id=?", d.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}