package store

import (
	"github.com/marcos-dev88/go-store-back/pkg/database"
	"strconv"
)

type Service interface {
	getStoreList() ([]Store, error)
	getStoreById(storeId string) (Store, error)
	createStore(store Store) (Store, error)
	updateStore(storeId string, store Store) (Store, error)
	deleteStore(storeId string) (bool, error)
}

type service struct{
	repository Repository
}

func (s service) getStoreList() ([]Store, error){
	storeList, err := s.repository.getStoreList()

	if err != nil {
		return nil, err
	}

	return storeList, nil
}

func (s service) getStoreById(storeId string) (Store, error) {
	idConverted, err := strconv.Atoi(storeId)
	newStore := *NewStore(0, "", "", "", "", "", nil, nil)

	if err != nil {
		return newStore, err
	}

	store, err := s.repository.getStoreById(idConverted)

	if err != nil {
		return newStore, err
	}

	return store, nil
}

func (s service) createStore(store Store) (Store, error){
	newStore := *NewStore(0, "", "", "", "", "", nil, nil)

	createdStore, err := s.repository.createStore(store)

	if err != nil {
		return newStore, err
	}

	return createdStore, nil
}

func (s service) updateStore(storeId string, store Store) (Store, error){
	newStore := *NewStore(0, "", "", "", "", "", nil, nil)

	idConverted, err := strconv.Atoi(storeId)

	if err != nil {
		return newStore, err
	}

	updatedStore, err := s.repository.updateStore(idConverted, store)

	if err != nil {
		return newStore, err
	}

	return updatedStore, nil
}

func (s service) deleteStore(storeId string) (bool, error){
	idConverted, err := strconv.Atoi(storeId)

	if err != nil {
		return false, err
	}

	idExists := database.VerifySExists(idConverted, "store")

	if !idExists {
		return false, nil
	}

	isDeletedStore, err := s.repository.deleteStore(idConverted)

	if err != nil {
		return false, err
	}

	return isDeletedStore, nil
}
