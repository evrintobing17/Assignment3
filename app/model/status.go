package model

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Data Status `json:"status"`
}

type StatusBahaya struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusAir   string `json:"status_air"`
	StatusAngin string `json:"status_angin"`
}
