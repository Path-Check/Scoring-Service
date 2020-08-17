package model

type ExposureNotificationRequest struct {
  ExposureSummaries      []ExposureSummary   `json:"exposure_summaries"`
}

type ExposureSummary struct {
  DateReceived           int     `json:"date_received, omitempty"`
  TimezoneOffset         int     `json:"timezone_offset, omitempty"`
  SeqNoInDay             int     `json:"seq_no_in_day, omitempty"`
  AttenuationDurations   AttenuationDurations     `json:"attenuation_durations, omitempty"`
  MatchedKeyCount        int     `json:"matched_key_count, omitempty"`
  DaysSinceLastExposure  int     `json:"days_since_last_exposure, omitempty"`
  MaximumRiskScore       int     `json:"maximum_risk_score, omitempty"`
  RiskScoreSum           int     `json:"risk_score_sum, omitempty"`
}

type AttenuationDurations struct {
  Low    int `json:"low, omitempty"`
  Medium int `json:"medium, omitempty"`
  High   int `json:"high, omitempty"`
}

type ExposureNotificationResponse struct {
  Notifications          []Notification       `json:"notifications"`
}

type Notification struct {
  ExposureSummaryRefs    []ExposureSummaryRef `json:"exposure_summary_refs"`
  DurationSeconds        int     `json:"duration_seconds"`
  // Note: we must have dateOfExposure OR dateMostRecentExposure + matchedKeyCount, but NOT both.
  DateOfExposure         int     `json:"date_of_exposure, omitempty"`
  DateMostRecentExposure int     `json:"date_most_recent_exposure, omitempty"`
  MatchedKeyCount        int     `json:"matched_key_count, omitempty"`
}

type ExposureSummaryRef struct {
  DateReceived           int     `json:"date_received, omitempty"`
  SeqNoInDay             int     `json:"seq_no_in_day, omitempty"`
}
