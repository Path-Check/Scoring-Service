package model

import "errors"

func WeightedDuration(exposureSummary *ExposureSummary) int {
	// TODO: use values from config
	return int(1*float32(exposureSummary.AttenuationDurations.Low) +
		// TODO: Does this do what I want it to do?
		0.5*float32(exposureSummary.AttenuationDurations.Medium) +
		0*float32(exposureSummary.AttenuationDurations.High))
}

// Calculate the day that the last exposure happened.
func GetExposureDay(exposureSummary *ExposureSummary) int {
	return exposureSummary.DateReceived - exposureSummary.DaysSinceLastExposure*24*3600
}

func CreateNotificationAggregated(newExposureSummary *ExposureSummary, unusedExposures *[]ExposureSummary, weightedDuration int) *ExposureNotificationResponse {
	// Start off the response by creating one using the newest ExposureSummary.
	response := CreateNotification(newExposureSummary)

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

func CreateNotification(exposureSummary *ExposureSummary) *ExposureNotificationResponse {
	response := &ExposureNotificationResponse{
		Notifications: []Notification{
			{ExposureSummaries: []ExposureSummary{*exposureSummary},
				DurationSeconds: WeightedDuration(exposureSummary)},
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
	empty_response := &ExposureNotificationResponse{}

	if request.NewExposureSummary.MatchedKeyCount == 0 {
		return empty_response, errors.New(
			"Matched key count was 0, this shouldn't have been sent.")
	}

	weightedDuration := WeightedDuration(&request.NewExposureSummary)

	if request.NewExposureSummary.MatchedKeyCount == 1 {
		// TODO: Use config here.
		if weightedDuration >= 15*60 {
			// We had a single exposure and it met the threshold, now create a
			// notification for it.
			return CreateNotification(&request.NewExposureSummary), nil
		}

		// Check if there were other exposures on the same day, where the
		// matched_key_count was 1, and if so aggregate and see if we exceed the
		// threshold in a day.
		unusedExposuresSameDay := FilterExposuresByDate(
			&request.UnusedExposureSummaries, GetExposureDay(&request.NewExposureSummary))
		// Add in any older exposures that occurred on the same day.
		for _, unusedExposure := range *unusedExposuresSameDay {
			weightedDuration += WeightedDuration(&unusedExposure)
		}
		// TODO: Again, use config.
		if weightedDuration >= 15*60 {
			return CreateNotificationAggregated(&request.NewExposureSummary, unusedExposuresSameDay, weightedDuration), nil
		}

		return empty_response, nil
	} else if request.NewExposureSummary.MatchedKeyCount == 2 ||
		request.NewExposureSummary.MatchedKeyCount == 3 {
		// TODO: use config here.
		// But also note, as written this does tie us into 15 minutes of
		// exposure to some extent, since the choice here has been made based on
		// the cap on buckets.
		if weightedDuration/request.NewExposureSummary.MatchedKeyCount >= 15*60 {
			// The average duration was over the threshold, so we know there
			// was at least one day that was over the threshold.
			return CreateNotification(&request.NewExposureSummary), nil
		}
	}

	// TODO: Add case for when MatchedKeyCount is 4+.

	return empty_response, nil
}
