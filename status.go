package main

import "log"

type status struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}

type StatusMessage struct {
	Status status `json:"status"`
}

func (statusMessage *StatusMessage) Error() string {
	return statusMessage.Status.Message
}

func NewSuccessStatus() StatusMessage {
	return StatusMessage{
		Status: status{
			ResponseCode: 200,
			Message:      "Query was successful!",
		}}
}

func NewNotFoundErrorStatus(message string, err error) *StatusMessage {
	log.Print("ERROR: ", err)
	return &StatusMessage{
		Status: status{
			ResponseCode: 310,
			Message:      "Not Found: " + message,
		}}

}

func NewDbErrorStatus(message string, err error) *StatusMessage {
	log.Print("ERROR: ", err)
	return &StatusMessage{
		Status: status{
			ResponseCode: 300,
			Message:      "DB Error: " + message,
		}}

}

func NewUnmarshallErrorStatus(message string, err error) *StatusMessage {
	log.Print("ERROR: ", err)

	return &StatusMessage{
		Status: status{
			ResponseCode: 100,
			Message:      "Ums Error: " + message,
		}}
}

func NewUnknownErrorStatus() StatusMessage {
	return StatusMessage{
		Status: status{
			ResponseCode: 999,
			Message:      "Unknown error!",
		}}
}
