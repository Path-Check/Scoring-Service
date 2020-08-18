package model

import "encoding/json"
import "fmt"
import "testing"

func TestOneExposure(t* testing.T) {
  requestData := []byte(`
    {"new_exposure_summary":
      { "date_received": 1597482000,
        "timezone_offset": 32400,
        "seq_no_in_day": 1,
        "attenuation_durations": {"low": 900, "medium": 0, "high": 0},
        "matched_key_count": 1,
        "days_since_last_exposure": 1,
        "maximum_risk_score": 1,
        "risk_score_sum": 1 }
    }`)

  var parsedRequest ExposureNotificationRequest;
  error := json.Unmarshal(requestData, &parsedRequest);
  if (error != nil) {
    fmt.Println(error)
  }

  responseData, _ := ScoreV1(&parsedRequest)

  response, error := json.Marshal(responseData)
  if (error != nil) {
    fmt.Println(error)
  }
  fmt.Println(string(response))

  expected := `{"notifications":[{"exposure_summaries":[{"date_received":1597482000,"timezone_offset":32400,"seq_no_in_day":1,"attenuation_durations":{"low":900,"medium":0,"high":0},"matched_key_count":1,"days_since_last_exposure":1,"maximum_risk_score":1,"risk_score_sum":1}],"duration_seconds":900 "date_of_exposure":1597395600}]}`

  // TODO: Why are fields that I didn't fill out still included in the response?
  // TODO: Use asserts/expects.
  if (string(response) != expected) {
    t.Errorf("Expected: %s Got: %s", expected, string(response));
  }
}

func TestInsufficientExposure(t* testing.T) {
  requestData := []byte(`
    {"new_exposure_summary":
      { "date_received": 1597482000,
        "timezone_offset": 32400,
        "seq_no_in_day": 1,
        "attenuation_durations": {"low": 90, "medium": 0, "high": 0},
        "matched_key_count": 1,
        "days_since_last_exposure": 1,
        "maximum_risk_score": 1,
        "risk_score_sum": 1 }
    }`)

  var parsedRequest ExposureNotificationRequest;
  error := json.Unmarshal(requestData, &parsedRequest);
  if (error != nil) {
    fmt.Println(error)
  }

  responseData, _ := ScoreV1(&parsedRequest)

  response, error := json.Marshal(responseData)
  if (error != nil) {
    fmt.Println(error)
  }
  fmt.Println(string(response))
}
