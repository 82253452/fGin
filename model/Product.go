package model

import (
	"github.com/jinzhu/gorm"
	. "fGin/config"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) Save() {
	Db.Save(p)
}
func (p *Product) delete() {
	Db.Delete(p)
}
func (p *Product) update() {
	Db.Update(p)
}
func (p *Product) find() {
	Db.Select(p)
}
func init() {
	Db.AutoMigrate(&Product{})
	//ids := []string{"1", "2"}
	//Db.Unscoped().Delete(Product{}, "code in (?)", ids) //in 匹配
	//p := Product{}
	//p.Code = "123"
	//Db.Save(&p)
}
