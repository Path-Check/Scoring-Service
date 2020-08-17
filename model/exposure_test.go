package model

import "encoding/json"
import "fmt"
import "testing"

func TestParseRequest(t* testing.T) {
  requestData := []byte(`
    {"exposure_summaries": [
      { "date_received": 1597482000,
        "timezone_offset": 32400,
        "seq_no_in_day": 1,
        "attenuation_durations": {"low": 180, "medium": 60, "high": 500},
        "matched_key_count": 1,
        "days_since_last_exposure": 1,
        "maximum_risk_score": 1,
        "risk_score_sum": 1 },
      { "date_received": 1597482000,
        "timezone_offset": 32400,
        "seq_no_in_day": 2,
        "attenuation_durations": {"low": 0, "medium": 300, "high": 300},
        "matched_key_count": 1,
        "days_since_last_exposure": 3,
        "maximum_risk_score": 1,
        "risk_score_sum": 1 }
    ]}`)

  var parsedRequest ExposureNotificationRequest;
  error := json.Unmarshal(requestData, &parsedRequest);
  if (error != nil) {
    fmt.Println(error)
  }

  if (parsedRequest.ExposureSummaries[0].DateReceived != 1597482000) {
    t.Errorf("Want: 1597482000 Got: %d",
             parsedRequest.ExposureSummaries[0].DateReceived)
  }
}

func TestWriteResponse(t* testing.T) {
  response_data := &ExposureNotificationResponse{
    Notifications: []Notification{
      {ExposureSummaryRefs: []ExposureSummaryRef{
         {DateReceived: 1597654800,
          SeqNoInDay: 1},
         {DateReceived: 1597654800,
          SeqNoInDay: 2}},
       DurationSeconds: 1800,
       DateOfExposure: 1597482000},
      {ExposureSummaryRefs: []ExposureSummaryRef{
         {DateReceived: 1597568400,
          SeqNoInDay: 1}},
       DurationSeconds: 900,
       DateMostRecentExposure: 1597482000,
       MatchedKeyCount: 3},
    },
  }
  response, error := json.Marshal(response_data)
  if (error != nil) {
    fmt.Println(error)
  }

  // TODO: Make less ugly test.
  expected := `{"notifications":[{"exposure_summary_refs":[{"date_received":1597654800,"seq_no_in_day":1},{"date_received":1597654800,"seq_no_in_day":2}],"duration_seconds":1800,"date_of_exposure":1597482000,"date_most_recent_exposure":0,"matched_key_count":0},{"exposure_summary_refs":[{"date_received":1597568400,"seq_no_in_day":1}],"duration_seconds":900,"date_of_exposure":0,"date_most_recent_exposure":1597482000,"matched_key_count":3}]}`

  // TODO: Use asserts/expects.
  if (string(response) != expected) {
    t.Errorf("Expected: %s Got: %s", expected, string(response));
  }
}
