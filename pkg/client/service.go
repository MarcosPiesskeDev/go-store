package client

import (
	"errors"
	"strconv"
	"time"
)

type Service interface {
	getClient(clientId string) (Client, error)
	getClientList() ([]Client, error)
	createClient(client Client) (Client, error)
	updateClient(clientId string, client Client) (Client, error)
	deleteClient(clientId string) bool
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

//Get Client
func (s *service) getClient(clientId string) (Client, error) {
	idconv, _ := strconv.Atoi(clientId)

	if len(clientId) == 0 {
		return *NewClient(0, 0, "", "", "", "", "", 0, time.Time{}), errors.New("id is empty")
	}

	client, er := s.repository.getClientById(idconv)

	if er != nil {
		return client, er
	}

	return client, nil
}

//Get Client List
func (s *service) getClientList() ([]Client, error) {
	clientList, err := s.repository.getAllClient()

	if err != nil {
		return []Client{}, err
	}

	return clientList, nil
}

//Create Client
func (s *service) createClient(client Client) (Client, error) {

	createdClient, er := s.repository.createClient(client)

	if er != nil {
		return *NewClient(0, 0, "", "", "", "", "", 0, time.Time{}), er
	}

	return createdClient, nil
}

//Update Client
func (s *service) updateClient(clientId string, client Client) (Client, error) {
	idconv, _ := strconv.Atoi(clientId)

	if len(clientId) == 0 {
		return *NewClient(0, 0, "", "", "", "", "", 0, time.Time{}), errors.New("id is empty")
	}

	updatedClient, er := s.repository.updateClient(idconv, client)

	if er != nil {
		return *NewClient(0, 0, "", "", "", "", "", 0, time.Time{}), er
	}

	return updatedClient, nil
}

//Delete a client
func (s *service) deleteClient(clientId string) bool {
	idconv, _ := strconv.Atoi(clientId)

	if len(clientId) == 0 {
		return false
	}
	isDeletedUser, er := s.repository.deleteClientById(idconv)

	if er != nil {
		return false
	}

	if !isDeletedUser {
		return false
	}
	return true
}
