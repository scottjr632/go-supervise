package db

import (
	"log"
	"testing"
)

func TestConnectToDB(t *testing.T) {
	c := Config{
		ConnectionString: "mongodb://localhost:27017",
	}
	_, err := NewDB(c)
	if err != nil {
		t.Errorf("Abs(-1) = %d; want 1", err)
	}
}

// func TestInsertDB(t *testing.T) {
// 	c := Config{
// 		ConnectionString: "mongodb://localhost:27017",
// 		DBName:           "test-config",
// 	}
// 	d, err := NewDB(c)
// 	if err != nil {
// 		t.Errorf("Abs(-1) = %d; want 1", err)
// 	}

// 	checkup := &entities.CheckUp{
// 		Worker: entities.Worker{
// 			WorkerID: "test2",
// 		},
// 		ActualResult: "test",
// 		ResponseCode: "204",
// 	}
// 	if err := d.SaveCheckup(checkup); err != nil {
// 		t.Error(err)
// 	}
// 	checkupd, _ := d.GetCheckUpsByWorkerID("test2")
// 	log.Println(checkupd)
// }

// func TestGetByResponseCode(t *testing.T) {
// c := Config{
// 	ConnectionString: "mongodb://localhost:27017",
// 	DBName:           "test-config",
// }
// d, err := NewDB(c)
// if err != nil {
// 	t.Errorf("Abs(-1) = %d; want 1", err)
// }
// 	checkups, err := d.GetCheckUpBySResponseCode("200")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	for _, c := range checkups {
// 		log.Println(*c)
// 	}
// }

func TestGetAllCheckups(t *testing.T) {
	c := Config{
		ConnectionString: "mongodb://localhost:27017",
		DBName:           "test-config",
	}
	d, err := NewDB(c)
	if err != nil {
		t.Errorf("Abs(-1) = %d; want 1", err)
	}
	cs, err := d.GetAllCheckups()
	if err != nil {
		t.Error(err)
	}

	for _, c := range cs {
		log.Println(*c)
	}
}
