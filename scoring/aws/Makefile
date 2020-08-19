.PHONY: test clean

test: deploy.done
	# curl -fsSL -D - "$$(terraform output url)?name=Lambda"

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
