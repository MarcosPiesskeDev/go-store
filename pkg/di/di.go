package di

import (
	"database/sql"
	"log"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/repository"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/database"
)

type Container struct {
	Db          *sql.DB
	storeRepo   *repository.StoreModel
	clientRepo  *repository.ClientModel
	productRepo *repository.ProductModel
}

func (co *Container) GetDb() (*sql.DB, error) {
	if co.Db == nil {
		var err error
		co.Db, err = database.GetConn()
		if err != nil {
			return nil, err
		}
	}
	return co.Db, nil
}

func (co *Container) GetStoreModel() *repository.StoreModel {
	if co.storeRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.storeRepo = repository.NewStoreModel(db)
	}
	return co.storeRepo
}

func (co *Container) GetClientModel() *repository.ClientModel {
	if co.clientRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.clientRepo = repository.NewClientModel(db)
	}
	return co.clientRepo
}

func (co *Container) GetProductModel() *repository.ProductModel {
	if co.productRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.productRepo = repository.NewProductModel(db)
	}
	return co.productRepo
}
