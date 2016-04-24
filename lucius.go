package lucius

import (
	"errors"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var dbhost, dbname string

func SetDBHost(host string) {
	dbhost = host
}

func SetDBName(name string) {
	dbname = name
}

func SetDB(host string, name string) {
	dbhost = host
	dbname = name
}

func Create(document *bson.M, collectionName string) error {
	
	if collectionName == "" {
		return errors.New("Variable collectionName is empty")
	}

  session, err := mgo.Dial(dbhost)
	if err != nil {
		return err
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)
	err = coll.Insert(&document)
	return err
}

func FindByID(ID bson.ObjectId, collectionName string) (*bson.M, error) {
	
	result := bson.M{}
	
	if collectionName == "" {
		return &result, errors.New("Variable collectionName is empty")
	}

  session, err := mgo.Dial(dbhost)
	if err != nil {
		return &result, err
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	condition := bson.M{"_id": ID}
	err = coll.Find(condition).One(&result)
	return &result, err
}

func FindBy(condition interface{}, collectionName string) (*bson.M, error) {
	
	result := bson.M{}
	
	if collectionName == "" {
		return &result, errors.New("Variable collectionName is empty")
	}

  session, err := mgo.Dial(dbhost)
	if err != nil {
		return &result, err
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	err = coll.Find(condition).One(&result)
	return &result, err
}

func FindAll(collectionName string) (*[]bson.M, error) {

	results := []bson.M{}

	if collectionName == "" {
		return &results, errors.New("Variable collectionName is empty")
	}

	session, err := mgo.Dial(dbhost)
	if err != nil {
		return &results, err
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	condition := bson.M{}
	err = coll.Find(condition).Sort("-timestamp").All(&results)
	return &results, err
}

func Update(condition bson.M, change bson.M, collectionName string) error {

	if collectionName == "" {
		return errors.New("Variable collectionName is empty")
	}

	session, err := mgo.Dial(dbhost)
	if err != nil {
		return err
	}


	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	err = coll.Update(condition, change)
	return err
}

func Patch(condition bson.M, change bson.M, collectionName string) error {

	if collectionName == "" {
		return errors.New("Variable collectionName is empty")
	}

	session, err := mgo.Dial(dbhost)
	if err != nil {
		return err
	}


	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	update := bson.M{}
	err = coll.Find(condition).One(&update)
	if err != nil {
		return err
	}


	for k, v := range change {
		update[k] = v
	}

	err = coll.Update(condition, update)
	return err
}

func Delete(ID bson.ObjectId, collectionName string) error {

	if collectionName == "" {
		return errors.New("Variable collectionName is empty")
	}

	session, err := mgo.Dial(dbhost)
	if err != nil {
		return err
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

  // collection
	coll := session.DB(dbname).C(collectionName)

	err = coll.Remove(bson.M{"_id": ID})
	return err
}