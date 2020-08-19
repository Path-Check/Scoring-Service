package model

type ExposureNotificationRequest struct {
  NewExposureSummary     ExposureSummary     `json:"new_exposure_summary"`
  UnusedExposureSummaries      []ExposureSummary   `json:"unused_exposure_summaries,omitempty"`
}

type ExposureSummary struct {
  DateReceived           int     `json:"date_received"`
  TimezoneOffset         int     `json:"timezone_offset"`
  SeqNoInDay             int     `json:"seq_no_in_day"`
  AttenuationDurations   AttenuationDurations     `json:"attenuation_durations"`
  MatchedKeyCount        int     `json:"matched_key_count"`
  DaysSinceLastExposure  int     `json:"days_since_last_exposure"`
  MaximumRiskScore       int     `json:"maximum_risk_score"`
  RiskScoreSum           int     `json:"risk_score_sum"`
}

// TODO: This whole omitemtpy business is very confusing. There doesn't seem to be a difference between the default value explicitly being filled out, and it missing. Those are not the same. Can we deal with this better?
type AttenuationDurations struct {
  Low    int `json:"low"`
  Medium int `json:"medium"`
  High   int `json:"high"`
}

type ExposureNotificationResponse struct {
  Notifications          []Notification       `json:"notifications,omitempty"`
}

type Notification struct {
  ExposureSummaries      []ExposureSummary `json:"exposure_summaries,omitempty"`
  DurationSeconds        int     `json:"duration_seconds"`
  // Note: we must have dateOfExposure OR dateMostRecentExposure + matchedKeyCount, but NOT both.
  DateOfExposure         int     `json:"date_of_exposure,omitempty"`
  DateMostRecentExposure int     `json:"date_most_recent_exposure,omitempty"`
  MatchedKeyCount        int     `json:"matched_key_count,omitempty"`
}
