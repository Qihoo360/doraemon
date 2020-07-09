.PHONY: run-ruleengine run-backend run-frontend

gateway.url = "default"
MAKEFLAGS += --warn-undefined-variables

# Build variables
REGISTRY_URI :=360cloud
RELEASE_VERSION :=$(shell git describe --always --tags)
UI_BUILD_VERSION :=v1.0.0
SERVER_BUILD_VERSION :=v1.0.0

# run module
run-ruleengine:
	cd cmd/rule-engine/ && export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go run main.go --gateway.url=http://$(gateway.url)

run-backend:
	cd cmd/alert-gateway/ && export GO111MODULE=on && export GOPROXY=https://goproxy.cn && go run main.go

run-frontend:
	cd web/ && npm run dev

# release, requiring Docker 17.05 or higher on the daemon and client
build-backend-image:
	@echo "version: $(RELEASE_VERSION)"
	docker build --no-cache -t $(REGISTRY_URI)/alert-gateway:$(RELEASE_VERSION) -f build/backend/Dockerfile .
build-frontend-image:
	@echo "version: $(RELEASE_VERSION)"
	docker build --no-cache -t $(REGISTRY_URI)/doraemon-frontend:$(RELEASE_VERSION) -f build/frontend/Dockerfile .

build-ruleengine-image:
	@echo "version: $(RELEASE_VERSION)"
	docker build --no-cache -t $(REGISTRY_URI)/rule-engine:$(RELEASE_VERSION) -f build/rule-engine/Dockerfile .


push-image:
	docker tag $(REGISTRY_URI)/alert-gateway:$(RELEASE_VERSION) $(REGISTRY_URI)/alert-gateway:latest
	docker push $(REGISTRY_URI)/alert-gateway:$(RELEASE_VERSION)
	docker push $(REGISTRY_URI)/alert-gateway:latest
	docker tag $(REGISTRY_URI)/doraemon-frontend:$(RELEASE_VERSION) $(REGISTRY_URI)/doraemon-frontend:latest
	docker push $(REGISTRY_URI)/doraemon-frontend:$(RELEASE_VERSION)
	docker push $(REGISTRY_URI)/doraemon-frontend:latest
	docker tag $(REGISTRY_URI)/rule-engine:$(RELEASE_VERSION) $(REGISTRY_URI)/rule-engine:latest
	docker push $(REGISTRY_URI)/rule-engine:$(RELEASE_VERSION)
	docker push $(REGISTRY_URI)/rule-engine:latest

release: build-backend-image build-frontend-image build-ruleengine-image push-image
