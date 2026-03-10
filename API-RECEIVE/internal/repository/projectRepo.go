package repository

import (
	"encoding/json"
	"fmt"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"gorm.io/gorm"
)

type projectRepository struct {
	db		*gorm.DB
	rabmq 	domain.QueueProvider
}

func NewProjectRepository(db *gorm.DB, rabmq domain.QueueProvider) domain.ProjectRepository {
	return &projectRepository{
		db: db,
		rabmq: rabmq,
	}
}

func (r *projectRepository) CreateProject(dto domain.ProjectDTO) (*domain.Project, error) {
	body, err := json.Marshal(dto);
    if err != nil {
        return nil, err;
    }

	fmt.Printf("Usecase: Sending order %s to queue\n", dto.Name);
	r.rabmq.Publish("project", body);
    return nil, nil;
}