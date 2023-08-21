package module

type PhoneList struct {
	Phones []string `json:"phones,omitempty"`
}

type Response struct {
	Response string `json:"response"`
}

type IteungV1Message struct {
	Phone_number string  `json:"phone_number"`
	Group_name   string  `json:"group_name"`
	Alias_name   string  `json:"alias_name"`
	Messages     string  `json:"messages"`
	Is_group     string  `json:"is_group"`
	Filename     string  `json:"filename"`
	Filedata     string  `json:"filedata"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Api_key      string  `json:"api_key"`
}

type IteungMessage struct {
	Phone_number string  `json:"phone_number"`
	Group_name   string  `json:"group_name"`
	Group_id     string  `json:"group_id"`
	Group        string  `json:"group"`
	Alias_name   string  `json:"alias_name"`
	Messages     string  `json:"messages"`
	Is_group     string  `json:"is_group"`
	Filename     string  `json:"filename"`
	Filedata     string  `json:"filedata"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
