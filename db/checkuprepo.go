package db

import (
	"context"
	"go-supervise/entities"
	"go-supervise/helpers"
	"go-supervise/services/checkup"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const checkupCollectionName = "checkups"

type CheckUpRepo interface {
	checkup.CheckUpRepo
	helpers.CheckUpRepo
	DeleteByWorkerID(workerID string) error
	GetCheckUpBySResponseCode(responseCode string) ([]*entities.CheckUp, error)
}

func (d *db) getCollection() *mongo.Collection {
	return d.Database(d.DBName).Collection("checkups")
}

func getCheckups(d *db, filter bson.D) ([]*entities.CheckUp, error) {
	collection := d.getCollection()
	var checkups []*entities.CheckUp
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		checkup := &entities.CheckUp{}
		if err := cur.Decode(checkup); err != nil {
			return nil, err
		}
		checkups = append(checkups, checkup)
	}
	return checkups, nil
}

func (d *db) SaveCheckup(checkup *entities.CheckUp) error {
	collection := d.Database(d.DBName).Collection("checkups")
	_, err := collection.InsertOne(context.TODO(), checkup)
	return err
}

func (d *db) GetCheckUpsByWorkerID(workerID string) ([]*entities.CheckUp, error) {
	filter := bson.D{{"worker.workerid", workerID}}
	return getCheckups(d, filter)
}

func (d *db) DeleteByWorkerID(workerID string) error {
	collection := d.getCollection()
	filter := bson.D{{"worker.workerid", workerID}}
	_, err := collection.DeleteMany(context.TODO(), filter)
	return err
}

func (d *db) GetCheckUpBySResponseCode(responseCode string) ([]*entities.CheckUp, error) {
	filter := bson.D{{"responsecode", responseCode}}
	return getCheckups(d, filter)
}

func (d *db) GetAllCheckups() ([]*entities.CheckUp, error) {
	filter := bson.D{{}}
	return getCheckups(d, filter)
}
