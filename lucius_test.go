package lucius

import (
  "testing"
	"labix.org/v2/mgo/bson"
)

func setDb() {
  SetDB("127.0.0.1", "crudtestdb")
}

var collectionName string = "testcollection"

func TestCreate(t *testing.T) {
  setDb()
  
  document := bson.M{
    "username": "joe",
    "password": "pass",
    "isactive": true,
    "visits": 523,
  }

  err := Create(&document, collectionName)

  if err != nil {
    t.Error("Expected error to be nil")
  }

}

func TestFindByAndFindById(t *testing.T) {
  setDb()

  condition := bson.M{"username": "joe"}
  user, err := FindBy(condition, collectionName)

  if err != nil {
    t.Error("Expected error for FindBy() to be nil")
  }

  userObj := *user
  ID := userObj["_id"].(bson.ObjectId)
  _, err = FindByID(ID, collectionName)

  if err != nil {
    t.Error("Expected error for FindByID() to be nil")
  }

}

func TestFindAll(t *testing.T) {
  setDb()
  
  _, err := FindAll(collectionName)

  if err != nil {
    t.Error("Expected error for FindAll() to be nil")
  }
}

func TestUpdate(t *testing.T) {
  setDb()
 
  condition := bson.M{"username": "joe"}
  change := bson.M{"username": "john", "visits": 524}
  err := Update(condition, change, collectionName)

  if err != nil {
    t.Error("Expected error for Update() to be nil")
  }
}

func TestPatch(t *testing.T) {
  setDb()

  document := bson.M{
    "username": "jenny",
    "password": "pass",
    "isactive": false,
    "visits": 523,
  }
  
  err := Create(&document, collectionName)

  if err != nil {
    t.Error("Expected error for Create() to be nil")
  }

  condition := bson.M{"username": "jenny"}
  change := bson.M{"username": "jane", "isactive": true}
  err = Patch(condition, change, collectionName)

  if err != nil {
    t.Error("Expected error for Patch() to be nil")
  }
}

func TestDelete(t *testing.T) {
  setDb()

  document := bson.M{
    "username": "deleteUser",
    "password": "pass",
    "isactive": true,
    "visits": 523,
  }
  
  err := Create(&document, collectionName)

  if err != nil {
    t.Error("Expected error for Create() to be nil")
  }

  condition := bson.M{"username": "deleteUser"}
  user, err := FindBy(condition, collectionName)

  if err != nil {
    t.Error("Expected error for FindBy() to be nil")
  }

  userObj := *user
  ID := userObj["_id"].(bson.ObjectId)
  
  err = Delete(ID, collectionName)

  if err != nil {
    t.Error("Expected error for Delete() to be nil")
  }

}
