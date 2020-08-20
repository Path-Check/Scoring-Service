# Scoring Service

Scoring is a simple stateless server that takes a Summary of Exposures for a given ENExposureConfiguration and returns an array of Notifications.

# General Flow

1. App gets EN Exposure Configuration from SERVER (`/v1/configuration`)
2. App passes configuration to GAEN API
3. GAEN runs exposure check and returns ExposureSummary
4. (If matched_key_count: 0, discard, done)

5. App constructs modified ExposureSummary object using structure from /v1/score input:
    1. Adds date_received (read comments in definition below)
    2. Timezone_offset
    3. Seq_no_in_day: this is saying it’s the n:th ExposureSummary we received today.

6. App sends new ExposureSummary, stored UnusedExposureSummaries to server
7. Server returns Notification array (might be empty) with any new notifications, contains ExposureSummaries that these notifications were based on

8. App does:
    1. removes any ExposureSummaries from UnusedExposureSummaries that are present in new notifications
    2. If no new notification: stores the new ExposureSummary in UnusedExposureSummaries
    3. Stores new Notifications if exist
    4. Displays notifications to users based on new Notification object

### State saved in app:

- Map/dictionary of unused exposure summaries (14 days), keyed by “date_received:seq_no_in_day”
- Notification array (14 days)
- Sequence # for ExposureSummary in day

## Environment

- GAEN v1.0 for now
- Working with ExposureSummary (Not ExposureInfo to avoid direct OS Notifications to users)
- Must be deployed in Google/AWS/Azure clouds.

## TODO (Week of Aug 17)

- [x] Scoring API Design (Lina)
---  [ ] Add scoring API design for ExposureWindows to doc (Lina)
- [x] Basic Go Implementation (Ray, Dave)
- [ ] Scoring Go Implementation (Lina)
--- [ ] Finish v1 scoring (Lina)
--- [ ] Make it use config instead of current hardcoded values (Lina? Ray?
Dave?)
- [ ] Deployment on AWS (Ray, David)
- [ ] cURL calls for the Mobile team and this document down below (Ray)
- [ ] Update Mobile app to use the basic scoring function (Matt?) 
- [ ] UnitTests with testing data captured by the mobile team (Ray)

## How to Install in Production

(Ray, David, please update)


### Installation and Setup

    - Install Golang: https://golang.org/
    - Install AWS CLI:https://aws.amazon.com/cli/
    - Log Into AWS CLI and configure credentials: https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started
        - run aws configure and input your account credentials
    - Install Terraform: https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/aws-get-started

### Run Makefile

    - cd into the target directory (ex: cd scoring/aws)
    - run make


## How to run in Development

1. Fork the repo to your GitHub user.

2. Clone to your computer.

```bash
git clone https://github.com/Path-Check/Scoring-Service.git
```

3. Run

```bash
make
```

## Running the Tests

(Ray, Dave, please update)

```bash
make test
```

## Important Links and References

- <https://www.google.com/covid19/exposurenotifications/>
- <https://developers.google.com/android/exposure-notifications/exposure-notifications-api>
- <https://developer.apple.com/documentation/exposurenotification/enexposureconfiguration>
- <https://github.com/google/exposure-notifications-server>
- <https://github.com/google/exposure-notifications-verification-server>
- <https://blog.google/inside-google/company-announcements/update-exposure-notifications>
- <https://blog.google/documents/73/Exposure_Notification_-_FAQ_v1.1.pdf>
- <https://www.aphlblog.org/bringing-covid-19-exposure-notification-to-the-public-health-community>
- <https://www.reuters.com/article/us-health-coronavirus-britain-tracing-ap/britains-covid-19-app-the-game-changer-that-wasnt-idUSKBN2400YM>
- <https://github.com/google/exposure-notifications-server/tree/main/examples/export>
- <https://github.com/google/exposure-notifications-server/blob/main/examples/export/testExport-2-records-1-of-1.zip>
