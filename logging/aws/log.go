package main

type LogRequest struct {
}

type LogResponse struct {
}

func Logger(req *LogRequest) (LogResponse, error) {
	// Log Something
	res := LogResponse{}
	return res, nil
}
