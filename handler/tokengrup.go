package handler

import (
	"fmt"
	"os"

	"github.com/aiteung/module/model"
	"github.com/whatsauth/watoken"
)

func TokenGroup(Pesan model.IteungMessage) string {
	group_id := Pesan.Group_id
	token, _ := watoken.EncodeforHours(group_id, os.Getenv("PRIVATEKEY"), 256)
	fmt.Println(token)
	return token
}
