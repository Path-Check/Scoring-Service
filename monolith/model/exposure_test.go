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

  // SRSLY, is this how testing works in Go? Because that's terrible. There's,
  // like, you know, equals, asserts, shit like that.
  if (parsedRequest.ExposureSummaries[0].DateReceived != 1597482000) {
    t.Errorf("Want: 1597482000 Got: %d",
             parsedRequest.ExposureSummaries[0].DateReceived)
  }
//  fmt.Printf("Request: %+v", parsedRequest);
}
