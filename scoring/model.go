package main

type ExposureNotificationRequest struct {
	NewExposureSummary    ExposureSummary   `json:"newExposureSummary"`
	UnusedExposureSummary []ExposureSummary `json:"unusedExposureSummary"`
}

type ExposureSummary struct {
	DateReceived          int                  `json:"dateReceived,omitempty"`
	TimezoneOffset        int                  `json:"timezoneOffset,omitempty"`
	SeqNoInDay            int                  `json:"seqNoInDay,omitempty"`
	AttenuationDurations  AttenuationDurations `json:"attenuationDurations,omitempty"`
	MatchedKeyCount       int                  `json:"matchedKeyCount,omitempty"`
	DaysSinceLastExposure int                  `json:"daysSinceLastExposure,omitempty"`
	MaximumRiskScore      int                  `json:"maximumRiskScore,omitempty"`
	RiskScoreSum          int                  `json:"riskScoreSum,omitempty"`
}

type AttenuationDurations struct {
	Low    int `json:"low,omitempty"`
	Medium int `json:"medium,omitempty"`
	High   int `json:"high,omitempty"`
}

type ExposureNotificationResponse struct {
	Notifications []Notification `json:"notifications"`
}

// Notification Note: we must have dateOfExposure OR dateMostRecentExposure + matchedKeyCount, but NOT both.
type Notification struct {
	ExposureSummaries      []ExposureSummary `json:"exposureSummaries"`
	DurationSeconds        int               `json:"durationSeconds"`
	DateOfExposure         int               `json:"dateOfExposure,omitempty"`
	DateMostRecentExposure int               `json:"dateMostRecentExposure,omitempty"`
	MatchedKeyCount        int               `json:"matchedKeyCount,omitempty"`
}
