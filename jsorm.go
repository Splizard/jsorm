/*
	A simple library for sending json status messages.
*/
package jsorm

import (
	"net/http"
	"fmt"
	"encoding/json"
)

const (
	Success = iota
	Failure
	Unauthorised
	Warning
)

type ServerMessage struct {
	Message string
	Code int
}

func marshal(object interface{}) string {
	data, err := json.Marshal(object)
	if err == nil {
		return string(data)
	}
	return ""
}

func Error(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(ServerMessage{Code:Failure, Message:message}))
}

func Warn(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(ServerMessage{Code:Warning, Message:message}))
}

func Code(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(ServerMessage{Code:code, Message:message}))
}

func AuthenticationError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(ServerMessage{Code:Unauthorised, Message:message}))
}

func AuthError(w http.ResponseWriter, message string) {
	AuthenticationError(w, message)
}

func Message(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(ServerMessage{Code:Success, Message:message}))
}

func Send(w http.ResponseWriter, object {}interface) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, marshal(object}))
}
