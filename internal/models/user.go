package models

type User struct {
	Id   uint32 `db:"id"`
	Name string `db:"name"`
	Age  uint8  `db:"age"`
}
