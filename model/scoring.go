package model
// TODO: This should perhaps be in a different package. But hacky-Lina doesn't
// know how to write Go so she's just gonna do this as quickly as she can.


import "fmt"

func WeightedDuration(exposureSummary *ExposureSummary) (int) {
  // TODO: use values from config
  return int(1 * float32(exposureSummary.AttenuationDurations.Low) +
    // TODO: Does this do what I want it to do?
    0.5 * float32(exposureSummary.AttenuationDurations.Medium) +
    0 * float32(exposureSummary.AttenuationDurations.High))
}

// Calculate the day that the last exposure happened.
func GetExposureDay(exposureSummary *ExposureSummary) (int) {
  return exposureSummary.DateReceived - exposureSummary.DaysSinceLastExposure * 24 * 3600
}

func CreateNotification(exposureSummary *ExposureSummary) (*ExposureNotificationResponse) {
  response := &ExposureNotificationResponse{
    Notifications: []Notification{
      {ExposureSummaries: []ExposureSummary{*exposureSummary},
       DurationSeconds: WeightedDuration(exposureSummary)},
    },
  }
  notification := &response.Notifications[0]

  lastExposureDate := GetExposureDay(exposureSummary)

  if (exposureSummary.MatchedKeyCount == 1) {
    notification.DateOfExposure = lastExposureDate
  } else {
    notification.DateMostRecentExposure = lastExposureDate
    notification.MatchedKeyCount = exposureSummary.MatchedKeyCount
  }

  return response
}

func ScoreV1(request *ExposureNotificationRequest) (*ExposureNotificationResponse, error) {
  if (request.NewExposureSummary.MatchedKeyCount == 0) {
    fmt.Println("Matched key count was 0, this shouldn't have been sent.");
    response_data := &ExposureNotificationResponse{
      Notifications: []Notification{},
    }
    return response_data, nil
  }

  weightedDuration := WeightedDuration(&request.NewExposureSummary);

  if (request.NewExposureSummary.MatchedKeyCount == 1) {
    // TODO: Use config here.
    if (weightedDuration >= 15 * 60) {
      // We had a single exposure and it met the threshold, now create a
      // notification for it.
      return CreateNotification(&request.NewExposureSummary), nil
    }
  }

  response_data := &ExposureNotificationResponse{
    Notifications: []Notification{
      {ExposureSummaries: []ExposureSummary{
         {DateReceived: 1597654800,
          SeqNoInDay: 1},
         {DateReceived: 1597654800,
          SeqNoInDay: 2}},
       DurationSeconds: 1800,
       DateOfExposure: 1597482000},
      {ExposureSummaries: []ExposureSummary{
         {DateReceived: 1597568400,
          SeqNoInDay: 1}},
       DurationSeconds: 900,
       DateMostRecentExposure: 1597482000,
       MatchedKeyCount: 3},
    },
  }

  return response_data, nil
}
