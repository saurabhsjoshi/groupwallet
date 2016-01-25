package main

type Status struct {
	ResponseCode int
	Message string
}

func NewSuccessStatus() Status{
	return Status{
		ResponseCode: 200,
		Message: "Query was successful!",
	}
}

func NewUnknownErrorStatus()  Status{
	return Status{
		ResponseCode: 999,
		Message: "Unknown error!",
	}
}
