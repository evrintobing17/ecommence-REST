package models

type APIResponseOptions struct {
	StatusCode     int
	Message        string
	Errors         error
	Data           interface{}
	ErrorInterface interface{}
}
