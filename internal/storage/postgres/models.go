package postgres

import "time"

type Client struct {
	Id      string
	Name    string
	TlsPem  string
	ValidTo time.Time
}

type Server struct {
	Id      string
	Name    string
	Address string
}

type Permission struct {
	ClientId string
	ServerId string
}
