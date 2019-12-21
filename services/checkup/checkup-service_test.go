package checkup

import (
	"testing"

	"go-supervise/entities"
)

func resetCheckupService() {
	checkUpInstance = &checkUpService{}
}

func TestGetCheckupService(t *testing.T) {
	defer resetCheckupService()
	s := GetCheckUpService()
	if s != checkUpInstance {
		t.Fail()
	}

	resetCheckupService()
}

func TestConfigureService(t *testing.T) {
	defer resetCheckupService()
	c := Config{UniqueIDs: true}
	GetCheckUpService().ConfigureService(c)
	if c != checkUpInstance.Config {
		t.Fail()
	}

	resetCheckupService()
}

func TestGetWorkers(t *testing.T) {
	defer resetCheckupService()
	if GetCheckUpService().GetWorkers() != nil {
		t.Fail()
	}

	resetCheckupService()
}

func TestSaveWorker(t *testing.T) {
	defer resetCheckupService()
	testWorker := &entities.Worker{WorkerID: "1"}
	if err := GetCheckUpService().SaveWorker(testWorker); err != nil {
		t.Error(err)
	}

	if checkUpInstance.Workers[0].WorkerID != "1" {
		t.Fail()
	}
}

func TestSaveWorkerWhenNotUnique(t *testing.T) {
	defer resetCheckupService()
	c := Config{
		UniqueIDs: true,
	}
	GetCheckUpService().ConfigureService(c)
	if !checkUpInstance.UniqueIDs {
		t.Fail()
	}

	testWorker := &entities.Worker{WorkerID: "1"}
	if err := GetCheckUpService().SaveWorker(testWorker); err != nil {
		t.Error(err)
	}

	if err := GetCheckUpService().SaveWorker(testWorker); err == nil {
		t.Fail()
	} else {
		if err.Error() != ErrWorkerAlreadyExists.Error() {
			t.Error(err)
		}
	}

}

func TestDeleteWorker(t *testing.T) {
	defer resetCheckupService()

	w := &entities.Worker{WorkerID: "1"}
	if err := GetCheckUpService().SaveWorker(w); err != nil {
		t.Error(err)
	}

	if len(GetCheckUpService().GetWorkers()) == 0 {
		t.Errorf("Worker not saved")
	}

	if err := GetCheckUpService().DeleteWorker(w.WorkerID); err != nil {
		t.Error(err)
	}

	if len(GetCheckUpService().GetWorkers()) != 0 {
		t.Error("Worker not deleted")
	}

	if err := GetCheckUpService().DeleteWorker(w.WorkerID); err == nil {
		t.Error("Worker deleted when not present")
	} else {
		if err.Error() != ErrWorkerNotFound.Error() {
			t.Error("Worker not found error not returned")
		}
	}
}

func TestFindWorkerByID(t *testing.T) {
	defer resetCheckupService()

	w := &entities.Worker{WorkerID: "1"}
	if err := GetCheckUpService().SaveWorker(w); err != nil {
		t.Error(err)
	}

	if len(GetCheckUpService().GetWorkers()) == 0 {
		t.Errorf("Worker not saved")
	}

	if i, wf := GetCheckUpService().FindWorkerByID(w.WorkerID); i < 0 || wf == nil {
		t.Error("Unable to find worker")
	} else {
		if wf != w {
			t.Error("Worker found does not equal worker saved")
		}
	}

	resetCheckupService()
	if i, wf := GetCheckUpService().FindWorkerByID(w.WorkerID); i >= 0 || wf != nil {
		t.Error("worker found when not present")
	}
}

func TestGetWorkerByID(t *testing.T) {
	defer resetCheckupService()

	w := &entities.Worker{WorkerID: "1"}
	if err := GetCheckUpService().SaveWorker(w); err != nil {
		t.Error(err)
	}

	if len(GetCheckUpService().GetWorkers()) == 0 {
		t.Errorf("Worker not saved")
	}

	if wf := GetCheckUpService().GetWorkerByID(w.WorkerID); wf == nil {
		t.Error("Worker was not able to be found")
	} else {
		if wf != w {
			t.Error("Worker found is not correct")
		}
	}
}
