package storage

import (
	"log"

	"github.com/RenzoMedina/api_go_echo_demo/pkg"
)

func StartedMigrate() {
	NewConnection()
	mydog := NewMySQLDog(Pool())
	myDogService := pkg.NewServices(mydog)

	if err := myDogService.Migrate(); err != nil {
		log.Fatalf("pkg.Migrate() %v", err)
	}
}
