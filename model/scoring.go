package model

import (
	"errors"
)

var (
	// Buckets are capped at 30 minutes of exposure each.
	maxBucketDuration = 30 * 60
	ErrNoConfig       = errors.New("No config.json input to scoring function")
)

func MaxWeightedDuration(attenuationWeights []float32) int {
	return int((attenuationWeights[0] + attenuationWeights[1] + attenuationWeights[2]) * float32(maxBucketDuration))
}

func WeightedDuration(exposureSummary *ExposureSummary, attenuationWeights []float32) int {
	return int(attenuationWeights[0]*float32(exposureSummary.AttenuationDurations.Low) +
		// TODO: Does this do what I want it to do?
		attenuationWeights[1]*float32(exposureSummary.AttenuationDurations.Medium) +
		attenuationWeights[2]*float32(exposureSummary.AttenuationDurations.High))
}

// Calculate the day that the last exposure happened.
func GetExposureDay(exposureSummary *ExposureSummary) int {
	return exposureSummary.DateReceived - exposureSummary.DaysSinceLastExposure*24*3600
}

func CreateNotificationAggregated(newExposureSummary *ExposureSummary, unusedExposures *[]ExposureSummary, weightedDuration int, attenuationWeights []float32) *ExposureNotificationResponse {
	// Start off the response by creating one using the newest ExposureSummary.
	response := CreateNotification(newExposureSummary, attenuationWeights)

	notification := &response.Notifications[0]
	// Append the additional ExposureSummaries.
	// TODO: This is really ugly. I want to append the entire array to the
	// existing array, but don't know how to do that? I did try just
	// append(unusedExposures, newExposureSummary), but that seems a bit ugly -
	// that's ignoring that newExposureSummary is actually already in there and
	// just overwrites it. Also, less important but still bugs me,
	// newExposureSummary ends up being last which is not what I expect.
	for _, exp := range *unusedExposures {
		notification.ExposureSummaries = append(notification.ExposureSummaries, exp)
	}

	// Update the duration with the weighted duration from all the exposures.
	notification.DurationSeconds = weightedDuration

	return response
}

func CreateNotification(exposureSummary *ExposureSummary, attenuationWeights []float32) *ExposureNotificationResponse {
	response := &ExposureNotificationResponse{
		Notifications: []Notification{
			{ExposureSummaries: []ExposureSummary{*exposureSummary},
				DurationSeconds: WeightedDuration(exposureSummary, attenuationWeights)},
		},
	}
	notification := &response.Notifications[0]

	lastExposureDate := GetExposureDay(exposureSummary)

	if exposureSummary.MatchedKeyCount == 1 {
		notification.DateOfExposure = lastExposureDate
	} else {
		notification.DateMostRecentExposure = lastExposureDate
		notification.MatchedKeyCount = exposureSummary.MatchedKeyCount
	}

	return response
}

func FilterExposuresByDate(exposureSummaries *[]ExposureSummary, date int) *[]ExposureSummary {
	filteredExposures := []ExposureSummary{}
	for _, exposureSummary := range *exposureSummaries {
		// Skip any exposures where we're not sure of the day they occurred.
		if exposureSummary.MatchedKeyCount != 1 {
			continue
		}

		// If this exposure happened on a single day and it was the date we're
		// looking for, add it to the set.
		if GetExposureDay(&exposureSummary) == date {
			filteredExposures = append(filteredExposures, exposureSummary)
		}
	}

	return &filteredExposures
}

func ScoreV1(request *ExposureNotificationRequest) (*ExposureNotificationResponse, error) {
	emptyResponse := &ExposureNotificationResponse{}
	attenuationWeights := request.ExposureConfiguration.AttenuationBucketWeights

	if len(request.ExposureConfiguration.AttenuationDurationThresholds) == 0 {
		return emptyResponse, ErrNoConfig
	}

	if request.NewExposureSummary.MatchedKeyCount == 0 {
		return emptyResponse, errors.New("Matched key count was 0.")
	}

	weightedDuration := WeightedDuration(&request.NewExposureSummary, attenuationWeights)

	if request.NewExposureSummary.MatchedKeyCount == 1 {
		// TODO: Use config here.
		if weightedDuration >= request.ExposureConfiguration.TriggerThresholdWeightedDuration*60 {
			// We had a single exposure and it met the threshold, now create a
			// notification for it.
			return CreateNotification(&request.NewExposureSummary, attenuationWeights), nil
		}

		// Check if there were other exposures on the same day, where the
		// matchedKeyCount was 1, and if so aggregate and see if we exceed the
		// threshold in a day.
		unusedExposuresSameDay := FilterExposuresByDate(
			&request.UnusedExposureSummaries, GetExposureDay(&request.NewExposureSummary))
		// Add in any older exposures that occurred on the same day.
		for _, unusedExposure := range *unusedExposuresSameDay {
			weightedDuration += WeightedDuration(&unusedExposure, attenuationWeights)
		}
		// TODO: Again, use config.
		if weightedDuration >= request.ExposureConfiguration.TriggerThresholdWeightedDuration*60 {
			return CreateNotificationAggregated(&request.NewExposureSummary, unusedExposuresSameDay, weightedDuration, attenuationWeights), nil
		}

		return emptyResponse, nil
	} else if request.NewExposureSummary.MatchedKeyCount == 2 ||
		request.NewExposureSummary.MatchedKeyCount == 3 {
		// TODO: use config here.
		// But also note, as written this does tie us into 15 minutes of
		// exposure to some extent, since the choice here has been made based on
		// the cap on buckets.
		if weightedDuration/request.NewExposureSummary.MatchedKeyCount >= 15*60 {
			// The average duration was over the threshold, so we know there
			// was at least one day that was over the threshold.
			return CreateNotification(&request.NewExposureSummary, attenuationWeights), nil
		}
	} else if request.NewExposureSummary.MatchedKeyCount >= 4 {
		if weightedDuration >= MaxWeightedDuration(attenuationWeights) {
			return CreateNotification(&request.NewExposureSummary, attenuationWeights), nil
		}
	}

	// TODO: Add case for when MatchedKeyCount is 4+.

	return emptyResponse, nil
}
