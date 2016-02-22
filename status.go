package main

type status struct {
	ResponseCode int    `json:"response_code"`
	Message      string `json:"message"`
}

type StatusMessage struct {
	Status status `json:"status"`
}

func NewSuccessStatus() StatusMessage {
	return StatusMessage{
		Status: status{
			ResponseCode: 200,
			Message:      "Query was successful!",
		}}
}

func NewDbErrorStatus() StatusMessage {
	return StatusMessage{
		Status: status{
			ResponseCode: 300,
			Message:      "Problem connecting with database!",
		}}
}

func NewUnknownErrorStatus() StatusMessage {
	return StatusMessage{
		Status: status{
			ResponseCode: 999,
			Message:      "Unknown error!",
		}}
}
