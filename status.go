package main

type Status struct {
	ResponseCode int
	Message string
}

const (
	SuccessStatus = Status{
		ResponseCode: 200,
		Message: "Query was successful!",
	}

	UnknownErrorStatus = Status{
		ResponseCode: 999,
		Message: "Unknown error!",
	}
)

