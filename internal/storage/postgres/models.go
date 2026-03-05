package postgres

import "time"

type Client struct {
	Id       string
	Name     string
	Tls_pem  string
	Valid_to time.Time
}

type Server struct {
	Id      string
	Name    string
	Address string
}

type Permission struct {
	Client_id string
	Server_id string
}
