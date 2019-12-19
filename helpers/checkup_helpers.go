package helpers

import "go-supervise/entities"

import "go-supervise/common"

import "errors"

var (
	ErrInvalidHTTPURI = errors.New("CheckUpURI must be valid http or https URI")
)

type WorkerRepo interface {
	SaveWorker(worker *entities.Worker) error
	GetWorkers() []*entities.Worker
	GetWorkerByID(workerID string) *entities.Worker
}

func validateNewWorker(worker *entities.Worker) error {
	if ok := common.ValidateHTTP(worker.CheckUpURI); !ok {
		return ErrInvalidHTTPURI
	}
	return nil
}

func AddWorker(worker *entities.Worker, repo WorkerRepo) error {
	if err := validateNewWorker(worker); err != nil {
		return err
	}
	return repo.SaveWorker(worker)
}
