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

func NewDbErrorStatus()  Status{
	return Status{
		ResponseCode: 300,
		Message: "Problem connecting with database!",
	}
}

func NewUnknownErrorStatus()  Status{
	return Status{
		ResponseCode: 999,
		Message: "Unknown error!",
	}
}



