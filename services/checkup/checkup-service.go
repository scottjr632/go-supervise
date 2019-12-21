package checkup

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"go-supervise/entities"

	"golang.org/x/sync/errgroup"
)

var checkUpOnce sync.Once
var checkUpInstance *checkUpService

// checkup service errors
var (
	ErrWorkerNotFound = errors.New("Unable to find worker by ID")
)

type checkUpService struct {
	sync.Mutex

	Workers []*entities.Worker
	Config
}

// Client ...
type Client interface {
	Get(url string) (resp *http.Response, err error)
}

// CheckUpRepo ...
type CheckUpRepo interface {
	SaveCheckup(*entities.CheckUp) error
}

// CheckUpService ...
type CheckUpService interface {
	GetWorkers() []*entities.Worker
	DoCheckUps(Client, CheckUpRepo) error
	SaveWorker(*entities.Worker) error
	DeleteWorker(string) error
	FindWorkerByID(string) (int, *entities.Worker)
	ConfigureService(config Config)
	GetWorkerByID(workerID string) *entities.Worker
	RunWithInterval(interval time.Duration, client Client, repo CheckUpRepo) error
}

func readResponse(r *http.Response) (string, error) {
	bod, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	return string(bod), err
}

// GetCheckUpService ...
func GetCheckUpService() CheckUpService {
	checkUpOnce.Do(func() { checkUpInstance = &checkUpService{} })
	return checkUpInstance
}

func (cs *checkUpService) ConfigureService(config Config) {
	cs.Config = config
}

func doCheckUp(worker *entities.Worker, client Client, repo CheckUpRepo) error {
	checkUp := &entities.CheckUp{Worker: worker}
	resp, err := client.Get(worker.CheckUpURI)
	if err != nil {
		checkUp.ResponseCode = "500"
		checkUp.ActualResult = err.Error()
	} else {
		checkUp.ResponseCode = resp.Status
		if worker.ExpectedResponse != "" {
			checkUp.ActualResult, err = readResponse(resp)
		}
	}

	if err := repo.SaveCheckup(checkUp); err != nil {
		return err
	}
	return nil
}

func (cs *checkUpService) GetWorkers() []*entities.Worker {
	return cs.Workers
}

func (cs *checkUpService) DoCheckUps(client Client, repo CheckUpRepo) error {
	log.Println("!!! RUNNINGS CHECKUP SERVICE !!!")
	g := errgroup.Group{}
	for _, worker := range cs.Workers {
		g.Go(func() error {
			return doCheckUp(worker, client, repo)
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (cs *checkUpService) SaveWorker(worker *entities.Worker) error {
	cs.Lock()
	defer cs.Unlock()

	if cs.UniqueIDs {
		if idx, _ := cs.FindWorkerByID(worker.WorkerID); idx > -1 {
			return errors.New("Worker already exists")
		}
	}

	cs.Workers = append(cs.Workers, worker)
	return nil
}

func (cs *checkUpService) DeleteWorker(workerID string) error {
	cs.Lock()
	defer cs.Unlock()
	if i, worker := cs.FindWorkerByID(workerID); worker != nil {
		cs.Workers[i] = cs.Workers[len(cs.Workers)-1]
		cs.Workers = cs.Workers[:len(cs.Workers)-1]
		return nil
	}
	return ErrWorkerNotFound
}

func (cs *checkUpService) FindWorkerByID(workerID string) (i int, worker *entities.Worker) {
	for i, worker := range cs.Workers {
		if worker.WorkerID == workerID {
			return i, worker
		}
	}
	return -1, nil
}

func (cs *checkUpService) GetWorkerByID(workerID string) *entities.Worker {
	_, worker := cs.FindWorkerByID(workerID)
	return worker
}

func (cs *checkUpService) RunWithInterval(interval time.Duration, client Client, repo CheckUpRepo) error {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			if err := cs.DoCheckUps(client, repo); err != nil {
				return err
			}
		}
	}
}
