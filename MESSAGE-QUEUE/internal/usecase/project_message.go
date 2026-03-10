package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/HanThamarat/NOTE-PLUS-MESSAGE-QUEUE/internal/domain"
	"gorm.io/gorm"
)

type ProjectUsecase struct {
	Listener 	domain.QueueListener
	db			*gorm.DB
}

func NewProjectUsecase(l domain.QueueListener, db *gorm.DB) *ProjectUsecase {
	return &ProjectUsecase{
		Listener: l,
		db: db,
	}
}

func (u *ProjectUsecase) Start() {
	msgs, err := u.Listener.Consume("project");
	if err != nil {
		fmt.Println("Error starting consumer:", err);
        return;
	}

	fmt.Println(" [*] Waiting for project messages. To exit press CTRL+C");

	for d := range msgs {
		result, err := u.process(d.Body);
		var messageBodyDTO domain.MessageBodyDTO;

		if d.ReplyTo != "" {
			messageBodyDTO.Status = true;
			resultByte, _ := json.Marshal(result);
			messageBodyDTO.Body = resultByte;
			response, _ := json.Marshal(messageBodyDTO);

			if err != nil {
				messageBodyDTO.Status = false;
				messageBodyDTO.Body = []byte(err.Error());
				response, _ = json.Marshal(messageBodyDTO);
			}

			u.Listener.Publish(d.ReplyTo, d.CorrelationId, response);
		}
	}
}

func (u *ProjectUsecase) process(data []byte) (*domain.Project, error) {
    var dto domain.ProjectDTO;
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, err;
	}

	var project domain.Project;
	
	project.Name 		= dto.Name;
	project.Description = dto.Description;
	project.Status 		= dto.Status;
	project.CreatedBy	= *dto.UserId;
	project.OrgId		= dto.OrgId;
    
	if err := u.db.Create(&project).Error; err != nil {
		return  nil, err;
	}

	return &project, nil;
}