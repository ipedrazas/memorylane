
NS = quay.io/ipedrazas
VERSION ?= latest

REPO = events-api
REPOPROXY = events-web
REPO_MIGRATION = events-migration
RELEASE = dev

.PHONY: build push release

build:
	docker build -t $(NS)/$(REPO):$(VERSION) ./server
	docker build -t $(NS)/$(REPOPROXY):$(VERSION) ./client
	docker build -t $(NS)/$(REPO_MIGRATION):$(VERSION) ./jobs

push:
	docker push $(NS)/$(REPO):$(VERSION)
	docker push $(NS)/$(REPOPROXY):$(VERSION)
	docker push $(NS)/$(REPO_MIGRATION):$(VERSION)

init-release: build push 
	helm init
	helm repo update
	helm dependency update ./charts/events
	helm upgrade --install $(RELEASE) charts/events --debug --set image.tag=$(VERSION),db.migration=true

release: build push
# we specify the postgresPassword to avoid helm re-generating the value:
#
#	{{ if .Values.postgresPassword }}
# 	postgres-password:  {{ .Values.postgresPassword | b64enc | quote }}
# 	{{ else }}
# 	postgres-password: {{ randAlphaNum 10 | b64enc | quote }}
# 	{{ end }}
	helm upgrade --install $(RELEASE) charts/events --debug --set image.tag=$(VERSION),postgresql.postgresPassword=$(PGPASSWORD)


default: build