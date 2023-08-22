package module

import (
	"strings"

	"github.com/aiteung/atdb"
	"github.com/aiteung/musik"
	"go.mongodb.org/mongo-driver/mongo"
)

func NormalizeAndTypoCorrection(message *string, MongoConn *mongo.Database, TypoCollection string) {
	typos := atdb.GetAllDoc[[]Typo](MongoConn, TypoCollection)
	*message = musik.NormalizeString(*message)
	for _, typo := range typos {
		*message = strings.ReplaceAll(*message, typo.From, typo.To)
	}

}
