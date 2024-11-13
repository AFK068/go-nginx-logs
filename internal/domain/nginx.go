package domain

import (
	"time"
)

type NGINXFields string

const (
	RemoteAddr    NGINXFields = "remote_addr"
	RemoteUser    NGINXFields = "remote_user"
	TimeLocal     NGINXFields = "time_local"
	Request       NGINXFields = "request"
	Status        NGINXFields = "status"
	BodyBytesSent NGINXFields = "body_bytes_sent"
	HTTPReferer   NGINXFields = "http_referer"
	HTTPUserAgent NGINXFields = "http_user_agent"
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

func (n *NGINX) GetFieldValue(field NGINXFields) any {
	fieldMap := map[NGINXFields]any{
		RemoteAddr:    n.RemoteAddr,
		RemoteUser:    n.RemoteUser,
		TimeLocal:     n.TimeLocal,
		Request:       n.Request,
		Status:        n.Status,
		BodyBytesSent: n.BodyBytesSent,
		HTTPReferer:   n.HTTPReferer,
		HTTPUserAgent: n.HTTPUserAgent,
	}

	return fieldMap[field]
}
