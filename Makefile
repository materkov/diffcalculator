VERSION=$(git rev-parse HEAD)
SERVICE=diffcalculator
TAG=563473344515.dkr.ecr.eu-central-1.amazonaws.com/$SERVICE:$VERSION

DEP_VERSION=0.4.1
BUILD_PATH=/go/src/github.com/materkov/diffcalculator

ci_build:
	curl -L -s https://github.com/golang/dep/releases/download/v$(DEP_VERSION)/dep-linux-amd64 -o $(GOPATH)/bin/dep
	chmod +x $(GOPATH)/bin/dep
	dep ensure
	
	go build -o diffcalculator cmd/main.go

	docker build . -t $(TAG)
	docker push $(TAG)

docker_build:
	docker run --rm -v $(PWD):$(BUILD_PATH) -w $(BUILD_PATH) golang:1.10.0 make ci_build
