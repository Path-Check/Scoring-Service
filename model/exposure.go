package model

type ExposureNotificationRequest struct {
	NewExposureSummary      ExposureSummary   `json:"newExposureSummary"`
	UnusedExposureSummaries []ExposureSummary `json:"unusedExposureSummaries,omitempty"`
}

type ExposureSummary struct {
	DateReceived          int                  `json:"dateReceived"`
	TimezoneOffset        int                  `json:"timeZoneOffset"`
	SeqNoInDay            int                  `json:"seqNoInDay"`
	AttenuationDurations  AttenuationDurations `json:"attenuationDurations"`
	MatchedKeyCount       int                  `json:"matchedKeyCount"`
	DaysSinceLastExposure int                  `json:"daysSinceLastExposure"`
	MaximumRiskScore      int                  `json:"maximumRiskScore"`
	RiskScoreSum          int                  `json:"riskScoreSum"`
}

// TODO: This whole omitemtpy business is very confusing. There doesn't seem to be a difference between the default value explicitly being filled out, and it missing. Those are not the same. Can we deal with this better?
type AttenuationDurations struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

type ExposureNotificationResponse struct {
	Notifications []Notification `json:"notifications,omitempty"`
}

type Notification struct {
	ExposureSummaries []ExposureSummary `json:"exposureSummaries,omitempty"`
	DurationSeconds   int               `json:"durationSeconds"`
	// Note: we must have dateOfExposure OR dateMostRecentExposure + matchedKeyCount, but NOT both.
	DateOfExposure         int `json:"dateOfExposure,omitempty"`
	DateMostRecentExposure int `json:"dateMostRecentExposure,omitempty"`
	MatchedKeyCount        int `json:"matchedKeyCount,omitempty"`
}
