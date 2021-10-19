package di

import (
	"database/sql"
	"github.com/marcos-dev88/go-store-back/pkg/client"
	"github.com/marcos-dev88/go-store-back/pkg/database"
	"github.com/marcos-dev88/go-store-back/pkg/product"
	"github.com/marcos-dev88/go-store-back/pkg/store"
	"log"
)

type Container interface {
	GetDb() (*sql.DB, error)
	GetClientRepository() client.Repository
	GetClientService() client.Service
	GetProductRepository() product.Repository
	GetProductService() product.Service
	GetStoreRepository() store.Repository
	GetStoreService() store.Service
}

type container struct {
	Db          *sql.DB
	storeRepo   store.Repository
	storeService store.Service
	clientRepo  client.Repository
	clientService client.Service
	productRepo product.Repository
	productService product.Service
}

func NewContainer() *container {
	return &container{}
}

func (co container) GetDb() (*sql.DB, error) {
	if co.Db == nil {
		var err error
		co.Db, err = database.GetConn()
		if err != nil {
			return nil, err
		}
	}
	return co.Db, nil
}

func (co container) GetClientRepository() client.Repository {
	if co.clientRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.clientRepo = client.NewRepository(db)
	}
	return co.clientRepo
}

func (co container) GetClientService() client.Service {
	if co.clientService == nil {
		co.clientService = client.NewService(co.GetClientRepository())
	}
	return co.clientService
}

func (co *container) GetProductRepository() product.Repository {
	if co.productRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.productRepo = product.NewRepository(db)
	}
	return co.productRepo
}

func (co container) GetProductService() product.Service {
	if co.productService == nil {
		co.productService = product.NewService(co.GetProductRepository())
	}
	return co.productService
}

func (co container) GetStoreRepository() store.Repository {
	if co.storeRepo == nil {
		db, err := co.GetDb()
		if err != nil {
			log.Fatal("Error on Db DI-> ", err)
			return nil
		}
		co.storeRepo = store.NewRepository(db, co.GetClientRepository(), co.GetProductRepository())
	}
	return co.storeRepo
}

func (co container) GetStoreService() store.Service {
	if co.storeService == nil {
		co.storeService = store.NewService(co.GetStoreRepository())
	}
	return co.storeService
}
