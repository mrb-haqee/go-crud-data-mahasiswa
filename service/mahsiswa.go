package service

import (
	"github.com/mrb-haqee/go-crud-data-mahasiswa/config"
	"github.com/mrb-haqee/go-crud-data-mahasiswa/model"
	"gorm.io/gorm"
)

type crudapp struct {
	db *gorm.DB
}

func NewDB() *crudapp {
	conn, err := config.Connect()
	if err != nil {
		panic(err)
	}
	conn.AutoMigrate(model.DataMahasiswa{})
	return &crudapp{conn}
}

func (c *crudapp) FindAll() ([]model.DataMahasiswa, error) {
	var GetData []model.DataMahasiswa
	err := c.db.Find(&GetData).Error
	if err != nil {
		return nil, err
	}
	return GetData, nil
}

func (c *crudapp) FindId(id int, mahasiswa *model.DataMahasiswa) error {
	return c.db.Model(&model.DataMahasiswa{}).Where("id = ?", id).Scan(mahasiswa).Error
}

func (c *crudapp) Add(mahasiswa model.Send) error {
	var get model.Mahasiswa
	get = model.Mahasiswa(mahasiswa)
	return c.db.Create(&model.DataMahasiswa{Mahasiswa: get}).Error
}

func (c *crudapp) Update(id int, mahasiswa model.Mahasiswa) error {
	return c.db.Model(&model.DataMahasiswa{}).Where("id = ?", id).Updates(model.DataMahasiswa{Model: gorm.Model{}, Mahasiswa: mahasiswa}).Error
}

func (c *crudapp) Delete(id int) error {
	return c.db.Where("id = ?", id).Delete(&model.DataMahasiswa{}).Error
}
