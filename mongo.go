package module

import (
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Keyword2Module(message string, db *mongo.Database) Module {
	filter := bson.M{"phone_number": message}
	collection := "module"
	atdb.GetOneDoc[Module](db, collection, filter)
	return Module{}

}
