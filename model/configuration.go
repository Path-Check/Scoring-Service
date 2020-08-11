package model

type Configuration struct {
	MinimumRiskScore                 int   `json:"minimumRiskScore"`
	AttenuationDurationThresholds    []int `json:"attenuationDurationThresholds"`
	AttenuationLevelValues           []int `json:"attenuationLevelValues"`
	DaysSinceLastExposure            []int `json:"daysSinceLastExposure"`
	DurationLevelValues              []int `json:"durationLevelValues"`
	TransmissionRiskLevelValues      []int `json:"transmissionRiskLevelValues"`
	AttenuationBucketWeights         []int `json:"attenuationBucketWeights"`
	TriggerThresholdWeightedDuration int   `json:"triggerThresholdweightedDuration"`
}
