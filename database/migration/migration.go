package migration

import (
	"fmt"

	"alfredtomo.dev/todos-app/database"
	"alfredtomo.dev/todos-app/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Todo{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Migration Completed")
}