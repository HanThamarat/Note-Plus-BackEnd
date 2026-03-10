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

func (u *ProjectUsecase) ProjectCreateStartWorker() {
	msgs, err := u.Listener.Consume("project.create");
	if err != nil {
		fmt.Println("Error starting consumer project.create :", err);
        return;
	}

	fmt.Println(" [*] Waiting for project.create messages. To exit press CTRL+C");

	for d := range msgs {
		result, err := u.prjectCreateProcess(d.Body);
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

func (u *ProjectUsecase) FindAllProjectByOrg() {
	msgs, err := u.Listener.Consume("project.find_all");
	if err != nil {
		fmt.Println("Error starting consumer project.find_all : ", err);
        return;
	}

	fmt.Println(" [*] Waiting for project.find_all messages. To exit press CTRL+C");

	for d := range msgs {
		result, err := u.findAllProjectByOrgProcess(d.Body);
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

func (u *ProjectUsecase) prjectCreateProcess(data []byte) (*domain.Project, error) {
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

func (u *ProjectUsecase) findAllProjectByOrgProcess(data []byte) (*[]domain.Project, error) {
	var orgId int;
	if err := json.Unmarshal(data, &orgId); err != nil {
		return nil, err;
	}

	var result []domain.Project;

	if err := u.db.Model(&result).Where("org_id = ?", orgId).Scan(&result).Error; err != nil {
		return nil, err;
	}

	return &result, nil;
}