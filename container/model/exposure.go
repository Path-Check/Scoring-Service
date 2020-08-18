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

// Note: we must have dateOfExposure OR dateMostRecentExposure + matchedKeyCount, but NOT both.
type Notification struct {
	ExposureSummaryRefs    []ExposureSummaryRef `json:"exposure_summary_refs"`
	DurationSeconds        int                  `json:"duration_seconds"`
	DateOfExposure         int                  `json:"date_of_exposure, omitempty"`
	DateMostRecentExposure int                  `json:"date_most_recent_exposure, omitempty"`
	MatchedKeyCount        int                  `json:"matched_key_count, omitempty"`
}

type ExposureSummaryRef struct {
	DateReceived int `json:"date_received, omitempty"`
	SeqNoInDay   int `json:"seq_no_in_day, omitempty"`
}
