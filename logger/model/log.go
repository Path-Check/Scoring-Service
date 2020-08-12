package model

type Event struct {
	Time string
	Pid  int
	Host string
}

type LogRequest struct {
	ExposureNotifications   []ExposureNotificationResponse `json:"exposureNotifications"`
	UnusedExposureSummaries []ExposureSummary              `json:"unusedExposureSummaries"`
	UserStatus              []TestData                     `json:"userStatus"`
}

type TestData struct {
	TestResult  bool `json:"testResult, omitempty"`
	DateEntered int  `json:"dateEntered, omitempty"`
	DateOfTest  int  `json:"dateOfTest, omitempty"`
}

type LogResponse struct {
	Status string `json:"status, omitempty"`
}
