# Scoring Service

Tells mobile app whether or not to notify everyone, whom an infected user has met the last 14 days, if user is infected. It takes a Summary of Exposures for a given `exposureconfiguration.json` (which changes based on health authority) and returns array of Notifications if certain scoring criteria are met.

# General Flow

1. App gets EN Exposure Configuration from SERVER (`/v1/configuration`)
2. App passes configuration to GAEN API
3. GAEN runs exposure check and returns ExposureSummary
4. (If matched_key_count: 0, discard, done)

5. App constructs modified ExposureSummary object using structure from /v1/score input:
    1. Adds date_received (read comments in definition below)
    2. Timezone_offset
    3. Seq_no_in_day: this is saying it’s the n:th ExposureSummary we received today.

6. App sends new ExposureSummary, stores UnusedExposureSummaries to server
7. Server returns Notification array (might be empty) with any new notifications, contains ExposureSummaries on which these notifications were based

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
  - [ ] Add scoring API design for ExposureWindows (Apple version 2) to doc (Lina)
- [x] Basic Go Implementation (Raymond, David)
- [ ] Scoring Go Implementation (Lina)
  - [ ] Finish v1 scoring (Lina)
  - [ ] Make it use config instead of current hardcoded values (Lina? Raymond?
Dave?)
- [X] Deployment on AWS (Raymond, David)
- [X] cURL calls for the Mobile team and this document down below (Raymond)
- [ ] Update Mobile app to use the basic scoring function (Matt?)
- [X] UnitTests with testing data captured by the mobile team (Raymond)

# How to Install in Production

(Raymond, David, please update)


### Installation and Setup

    - Install Golang: https://golang.org/
        - This is necessary to compile and test the go application
    - Install AWS CLI:https://aws.amazon.com/cli/
        - This is necessary to deploy resources to AWS
    - Log Into AWS CLI and configure credentials: https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started
        - run 'aws configure' and input your account credentials
        - Configure your AWS access key and secret key with the aws configure command, or just create a file ~/.aws/credentials containing the keys:
            [default]
            aws_access_key_id = KEY
            aws_secret_access_key = KEY
    - Install Terraform: https://learn.hashicorp.com/tutorials/terraform/install-cli?in=terraform/aws-get-started
        - This is necessary to provision the necessary infrastructure required to deploy the lambda function

### Run Makefile (with AWS CLI and Terraform installed)

    - cd into the target directory (ex: cd scoring/aws)
    - run 'make'


## How to run in Development

1. Fork the repo to your GitHub user.

2. Clone to your computer.

```
git clone https://github.com/Path-Check/Scoring-Service.git
```

3. Run (with AWS CLI and Terraform installed)

```
make
```

## Running the Tests

```
go test -v ./... (all tests, period)
go test -v ./model (all tests in /model)
go test -v ./scoring/aws (all tests in /scoring/aws)
```

With AWS CLI and Terraform installed:
```
make test
```
## [AWS Lambda](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html) Function
- [See AWS lambda function dashboard](https://console.aws.amazon.com/lambda/home?region=us-east-1#/functions/scoring?tab=monitoring)
- To manually test lambda function with an input, click **Test** in upper right corner. You may use "ScoreTest" or configure your own test event.

## [AWS CloudWatch logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html)
- [See CloudWatch logs](https://us-east-1.console.aws.amazon.com/cloudwatch/home?region=us-east-1#logsV2:log-groups/log-group/$252Faws$252Flambda$252Fscoring) for our scoring lambda function. Click on or search "log streams."
- [Run queries on logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_RunSampleQuery.html).

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

## Design Documents

- [Data Format For Data Scoring and Metrics](https://docs.google.com/document/d/18UM5T_8PSZ4mJaRpz0H3UDnwyyta2o_GmxtyI396xYs/edit#heading=h.88dqztbzgbbp)
- [Study Design and Methods](https://docs.google.com/document/d/1FT4J29c2_k5gBdCf04BN7X9HbLCrN1eNmOu0ehgHZjY/edit)
- [GAEN Scoring Design](https://docs.google.com/document/d/12vU48fwOcGvIYLR7Y0jnSK_ZIeGvwvvszs6EjG_HHNE/edit#heading=h.bg7iuv59zi1d)
