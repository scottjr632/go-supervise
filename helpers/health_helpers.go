package helpers

import (
	"go-supervise/entities"

	"errors"

	"log"
)

type HealthWrapper struct {
	Status   string              `json:"status"`
	CheckUps []*entities.CheckUp `json:"checkUps,omitempty"`
	*entities.Worker
}

type CheckUpRepo interface {
	GetCheckUpsByWorkerID(workerID string) ([]*entities.CheckUp, error)
	GetAllCheckups() ([]*entities.CheckUp, error)
}

func GetHealthByWorkerID(workerID string, repo CheckUpRepo) (*entities.Health, error) {
	checkups, err := repo.GetCheckUpsByWorkerID(workerID)
	if err != nil || checkups == nil {
		return nil, err
	}

	health := &entities.Health{Worker: checkups[0].Worker}
	health.SetHealthStatus(checkups)
	return health, nil
}

func GetHealthForAllWorkers(workerRepo WorkerRepo, checkUpRepo CheckUpRepo) ([]*entities.Health, error) {
	workers := workerRepo.GetWorkers()
	if workers == nil {
		return nil, errors.New("No workers present")
	}

	var healthStatus []*entities.Health
	for _, worker := range workers {
		log.Print(checkUpRepo)
		health := &entities.Health{}
		health.Worker = workerRepo.GetWorkerByID(worker.WorkerID)

		checkups, err := checkUpRepo.GetCheckUpsByWorkerID(worker.WorkerID)
		switch true {
		case err != nil:
			return nil, err
		case checkups != nil:
			health.SetHealthStatus(checkups)
		}

		healthStatus = append(healthStatus, health)
	}
	return healthStatus, nil
}
