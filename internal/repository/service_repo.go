package repository

import (
	"database/sql"
	"status-dashboard/internal/models"
)

// ServiceRepo handles CRUD for services.
type ServiceRepo struct {
	DB *sql.DB
}

func NewServiceRepo(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{DB: db}
}

func (r *ServiceRepo) Create(s *models.Service) error {
	// TODO: implement insert
	return nil
}

func (r *ServiceRepo) List() ([]models.Service, error) {
	// TODO: implement query
	return nil, nil
}

func (r *ServiceRepo) GetByID(id int) (*models.Service, error) {
	// TODO: implement
	return nil, nil
}

func (r *ServiceRepo) Delete(id int) error {
	// TODO: implement
	return nil
}
