
NS = quay.io/ipedrazas
VERSION ?= latest

REPO = events-api
REPOPROXY = events-web
RELEASE = dev

.PHONY: build push release

build:
	docker build -t $(NS)/$(REPO):$(VERSION) ./server
	docker build -t $(NS)/$(REPOPROXY):$(VERSION) ./client

push:
	docker push $(NS)/$(REPO):$(VERSION)
	docker push $(NS)/$(REPOPROXY):$(VERSION)

init-release: 
	helm init
	helm dependency update ./charts/events

release: build push
	helm upgrade --install $(RELEASE) charts/events --debug --set image.tag=$(VERSION),api.tag=$(VERSION),db.release=wonderful-lemur

default: build