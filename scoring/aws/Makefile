# Makefiles, putting CI/CD businesses out of business since 1976 lmao
# .PHONY calls the test and clean functions
.PHONY: test clean

# test calls deploy.done then issues a curl request to the deployed lambda function and API Gateway
test: deploy.done
	curl -fsSL -d '{"new_exposure_summary":{"date_received": 1597482000,"timezone_offset": 32400,"seq_no_in_day": 1,"attenuation_durations": {"low": 900, "medium": 0, "high": 0},"matched_key_count": 1,"days_since_last_exposure": 1,"maximum_risk_score": 1,"risk_score_sum": 1}}' -H "Content-Type: application/json" -X POST "$$(terraform output url)"

# clean destroys the infrastructure and removes the builds
clean:
	terraform destroy
	rm -f scoring.zip scoring

# init.done initializes terraform 
init.done:
	terraform init
	echo terraform initialized

# deploy.done calls init.done and passes in main.tf builds scoring.zip
deploy.done: init.done main.tf scoring.zip
	terraform apply
	echo terraform applied and application deployed

# scoring.zip zips the compiled build into a zip file
scoring.zip: scoring
	zip $@ $<

# scoring compiles and builds a binary according to linux amd64
scoring: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@
