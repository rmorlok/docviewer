package models


import (
	"golangRestfulAPISample/app/models/orm"
	"time"
)

type (
	Document struct {
		BaseModel
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)

func (m *Document) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

func (m *Document) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

func Create(m *Document) (*Document, error) {
	var err error
	err = orm.Create(&m)
	return m, err
}

func (m *Document) Update() error {
	var err error
	err = orm.Save(&m)
	return err
}

func (m *Document) Delete() error {
	var err error
	err = orm.Delete(&m)
	return err
}

func FindByID(id uint64) (Document, error) {
	var (
		document Document
		err  error
	)
	err = orm.FindOneByID(&document, id)
	return document, err
}

func All() ([]Document, error) {
	var (
		documents []Document
		err   error
	)
	err = orm.FindAll(documents)
	return documents, err
}