# .PHONY calls the test and clean functions
.PHONY: test clean

# test calls deploy.done then issues a curl request to the deployed lambda function and API Gateway
test: deploy.done
	curl -fsSL -d '{"newExposureSummary":{"dateReceived": 1597482000,"timeZoneOffset": 32400,"seqNoInDay": 1,"attenuationDurations": {"low": 900, "medium": 0, "high": 0},"matchedKeyCount": 1,"daysSinceLastExposure": 1,"maximumRiskScore": 1,"riskScoreSum": 1}}' -H "Content-Type: application/json" -X POST "$$(terraform output url)"

# clean destroys the infrastructure and removes the builds
clean:
	terraform destroy
	rm -f scoring.zip scoring
	@echo "\nScoring infrastructure destroyed, zip file removed, and scoring binary removed!\n"

# init.done initializes terraform
init.done:
	terraform init
	@echo "\nTerraform initialized!\n"

# deploy.done calls init.done and passes in main.tf builds scoring.zip
deploy.done: init.done main.tf scoring.zip
	terraform apply
	@echo "\nTerraform applied and application deployed!\n"

# scoring.zip zips the compiled build into a zip file
scoring.zip: scoring
	zip $@ $<
	@echo "\nScoring binary zipped!\n"

# scoring compiles and builds a binary according to linux amd64
scoring: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@
	@echo "\nScoring binary compiled!\n"
