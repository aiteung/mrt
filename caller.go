package module

import (
	"github.com/aiteung/module/handler"
	"github.com/aiteung/module/model"
)

func Caller(Modulename string, Pesan model.IteungMessage) {
	switch Modulename {
	case "tokengrup":
		handler.TokenGroup(Pesan)
	}

}
