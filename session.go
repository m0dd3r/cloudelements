package h2go

import (
	"fmt"
	"gopkg.in/jmcvetta/napping.v2"
	"net/http"
)

const (
	BASE_URL = "https://api.cloud-elements.com/elements/api-v2"
)

type Session struct {
	user_secret    string
	org_secret     string
	element_secret string
	napping.Session
}

func NewSession(us, os, es string) *Session {
	h := http.Header{}
	h.Add("Authorization", fmt.Sprintf("User %s, Organization %s, Element %s", us, os, es))
	h.Add("Content-type", "application/json")
	h.Add("Accept", "application/json")

	s := &Session{
		user_secret:    us,
		org_secret:     os,
		element_secret: es,
		Session: napping.Session{
			Header: &h,
		},
	}
	return s
}
