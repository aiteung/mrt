package generator

import "github.com/aiteung/module/model"

func Pattern2Message(Pesan *model.IteungMessage) {
	if (Pesan.Latitude != 0) && (Pesan.Longitude != 0) && (Pesan.Is_group) {
		Pesan.Message = MulaiKelas(Pesan.Group_name)
	}
}
