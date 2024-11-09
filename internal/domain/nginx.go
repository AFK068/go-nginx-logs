package domain

import (
	"time"
)

type NGINX struct {
	RemoteAddr    string
	RemoteUser    string
	TimeLocal     time.Time
	Request       string
	Status        int
	BodyBytesSent int
	HTTPReferer   string
	HTTPUserAgent string
}

func NewNGINX(remoteAddr, remoteUser string, timeLocal time.Time,
	request string, status, bodyBytesSent int, htttpReferer, httpUserAgent string) *NGINX {
	return &NGINX{remoteAddr, remoteUser, timeLocal, request, status, bodyBytesSent, htttpReferer, httpUserAgent}
}
