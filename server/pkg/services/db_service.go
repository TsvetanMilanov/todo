package services

import (
	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DbService database related methods.
type DbService struct {
	Helpers  types.IHelpers `inject:"helpers"`
	database *mgo.Database
	session  *mgo.Session
}

// InitializeDatabase executes the required steps for the loading of the DB.
func (db *DbService) InitializeDatabase() error {
	session, err := mgo.Dial(db.getConnectionString())

	if err != nil {
		return err
	}

	db.session = session
	db.database = db.session.DB(constants.DbName)

	err = db.seedInitialData()

	return err
}

// GetCollection returns a collection from the database.
func (db *DbService) GetCollection(collection string) *mgo.Collection {
	return db.database.C(collection)
}

// Dispose closes/releases all the db resources.
func (db *DbService) Dispose() error {
	db.database.Logout()
	db.session.Close()

	return nil
}

func (db *DbService) seedInitialData() error {
	users := db.database.C(constants.UsersCollectionName)

	admin := models.User{}
	count, err := users.Find(bson.M{"username": constants.AdminUsername}).Count()

	if err != nil {
		return err
	}

	if count == 0 {
		admin.Username = constants.AdminUsername
		admin.Roles = []string{constants.AdminRole, constants.ModeratorRole, constants.UserRole}
		admin.Todos = []string{}
		pass, err := db.Helpers.EncryptString("adminpass")

		if err != nil {
			return err
		}

		admin.Password = pass

		err = users.Insert(admin)

		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DbService) getConnectionString() string {
	defaultConnectionString := "127.0.0.1"

	return db.Helpers.GetEnvVariableOrDefault(constants.DbServerEnvVar, defaultConnectionString)
}
