.PHONY: test clean

test: deploy.done
	curl -fsSL -d '{"new_exposure_summary":{"date_received": 1597482000,"timezone_offset": 32400,"seq_no_in_day": 1,"attenuation_durations": {"low": 900, "medium": 0, "high": 0},"matched_key_count": 1,"days_since_last_exposure": 1,"maximum_risk_score": 1,"risk_score_sum": 1}}' -H "Content-Type: application/json" -X POST "$$(terraform output url)"

clean:
	terraform destroy
	rm -f init.done deploy.done scoring.zip scoring

init.done:
	terraform init
	touch $@

deploy.done: init.done main.tf scoring.zip
	terraform apply
	touch $@

scoring.zip: scoring
	zip $@ $<

scoring: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@
