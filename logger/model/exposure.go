package model

type ExposureNotificationRequest struct {
	NewExposureSummary     ExposureSummary   `json:"newExposureSummary"`
	CurrentExposureSummary []ExposureSummary `json:"currentExposureSummary, omitempty"`
}

type ExposureSummary struct {
	AttenuationDurations  AttenuationDurations `json:"attenuationDurations, omitempty"`
	MatchedKeyCount       int                  `json:"matchedKeyCount, omitempty"`
	DaysSinceLastExposure int                  `json:"daysSinceLastExposure, omitempty"`
	TimeReceived          int                  `json:"timeReceived"`
}

type AttenuationDurations struct {
	LowBucketDurationSeconds    int `json:"lowBucketDurationSeconds, omitempty"`
	MediumBucketDurationSeconds int `json:"mediumDurationSeconds, omitempty"`
	HighBucketDurationSeconds   int `json:"highBucketDurationSeconds, omitempty"`
}

type ExposureNotificationResponse struct {
	ExposureSummaries []ExposureSummary `json:"exposureSummaries"`
}
