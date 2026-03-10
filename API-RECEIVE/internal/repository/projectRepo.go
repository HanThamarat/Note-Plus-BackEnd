package repository

import (
	"encoding/json"
	"errors"
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
	response, err := r.rabmq.Call("project", body);
	if err != nil {
        return nil, err;
    }

	var result domain.Project;
	var message domain.MessageBodyDTO;
	json.Unmarshal(response, &message);

	if message.Status == false {
		return nil, errors.New("Insert data to database failed.");
	}

	json.Unmarshal(message.Body, &result);
	
    return &result, nil;
}