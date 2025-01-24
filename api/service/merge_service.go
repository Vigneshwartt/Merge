package service

import (
	"allcaps/api/repository"
	"allcaps/pkg/models"
)

type IServiceMerge interface {
	// GetClientData(url string) (map[string]interface{}, error)
	GetClientData(url string) (*models.ApiResponse, error)
	PostClientData(url string, data []map[string]interface{}) (*models.ApiResponse, error)
}

type Serv struct {
	Repo repository.IRepoInterface
}

func InitServiceClient(service repository.IRepoInterface) IServiceMerge {
	return &Serv{Repo: service}
}

func (service *Serv) GetClientData(url string) (*models.ApiResponse, error) {
	return service.Repo.GetClientData(url)
}

// func (service *Serv) GetClientData(url string) (map[string]interface{}, error) {
// 	return service.Repo.GetClientData(url)
// }

func (service *Serv) PostClientData(url string, data []map[string]interface{}) (*models.ApiResponse, error) {
	return service.Repo.PostClientData(url, data)
}
