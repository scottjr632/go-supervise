package helpers

import "go-supervise/entities"

import "errors"

type CheckUpRepo interface {
	GetCheckUpsByWorkerID(workerID string) ([]*entities.CheckUp, error)
	GetAllCheckups() ([]*entities.CheckUp, error)
}

func GetHealthByWorkerID(workerID string, repo CheckUpRepo) (*entities.Health, error) {
	checkups, err := repo.GetCheckUpsByWorkerID(workerID)
	if err != nil {
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
		checkups, err := checkUpRepo.GetCheckUpsByWorkerID(worker.WorkerID)
		if err != nil {
			return nil, err
		}
		health := &entities.Health{Worker: checkups[0].Worker}
		health.SetHealthStatus(checkups)
		healthStatus = append(healthStatus, health)
	}
	return healthStatus, nil
}
