package main

import (
	"net/url"
	"os"
	"strconv"
)

const (
	EnvKeySendPinByEmail     = "SEND_PIN_BY_EMAIL"
	EnvKeyProjectPageBaseURL = "PROJECT_PAGE_BASE_URL"
)

func GetEnvBool(key string) bool {
	if val, err := strconv.ParseBool(os.Getenv(key)); err != nil {
		return val
	}

	return false
}

func GetEnvProjectPageBaseURL() url.URL {
	if bu := os.Getenv(EnvKeyProjectPageBaseURL); bu != "" {
		u, err := url.Parse(bu)
		if err == nil {
			return *u
		}
	}

	return url.URL{Scheme: "https", Host: "p.apartomat.ru"}
}
