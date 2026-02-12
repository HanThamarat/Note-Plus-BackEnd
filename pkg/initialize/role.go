package initial

import (
	"fmt"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"gorm.io/gorm"
)

func RoleInit(db *gorm.DB) {
	var role [3]domain.Role;
	var count int64;

	recheck := db.Model(&role).Count(&count);

	if recheck.Error != nil {
		fmt.Println("Have something wrong in initialize progress : ", recheck.Error);
		return;
	}

	if count != 0 {
		fmt.Println("✅ Member Role initialize success.");
		return;
	}

	role[0].Name 	= "Owner";
	role[0].Status 	= true;
	role[1].Name 	= "Viewer";
	role[1].Status 	= true;
	role[2].Name 	= "Editor";
	role[2].Status 	= true;

	creareNewRoles := db.Create(&role);

	if creareNewRoles.Error != nil {
		fmt.Println("Have something wrong in initialize progress : ", creareNewRoles.Error);
		return;
	}

	fmt.Println("✅ Member Role initialize success.");
}