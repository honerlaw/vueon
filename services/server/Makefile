.PHONY: build docker clean deploy destroy docker-deploy test

TERRAFORM_SRC_PATH=../../terraform/src/server/
TERRAFORM_VAR_FILE=../../terraform/env/server/example.tfvars

TAGS=$(shell aws ecr describe-images --repository-name vueon --output json --query 'sort_by(imageDetails,& imagePushedAt)[*].imageTags[*]')
LATEST_TAG=$(shell echo $(TAGS) | jq '.[length - 1] | .[length - 1] | tonumber')
TAG=$$(( $(LATEST_TAG) + 1 ))
URL=106480132517.dkr.ecr.us-east-1.amazonaws.com/vueon:$(TAG)
LATEST_URL=106480132517.dkr.ecr.us-east-1.amazonaws.com/vueon:$(LATEST_TAG)

clean:
	rm -rf bin

test:
	go clean -testcache
	go test ./core/repository/
	go test ./user/service/
	go test ./payment/service/

build:
	cd ./src/server/ && go get -d -v
	go build -o bin/server ./src/server/main.go

docker:
	eval $(shell aws ecr get-login --no-include-email --region=us-east-1)
	docker build -t $(URL) .
	docker push $(URL)

docker-deploy: docker deploy

deploy:
	terraform init $(TERRAFORM_SRC_PATH)
	terraform apply -auto-approve -var-file=$(TERRAFORM_VAR_FILE) -var="image=$(LATEST_URL)" $(TERRAFORM_SRC_PATH)

destroy:
	terraform init $(TERRAFORM_SRC_PATH)
	terraform destroy -auto-approve -var-file=$(TERRAFORM_VAR_FILE) -var="image=$(LATEST_URL)" $(TERRAFORM_SRC_PATH)

run-local:
	go run src/server/main.go