.PHONY: deploy

export PATH := ${PATH}:${PWD}/webserver/node_modules/.bin
export KUBERNETES_CONFIG := ${HOME}/.kube/config
export WATCH_NAMESPACE := default


## Front-end targets
frontend-build:
	docker build -f docker/Dockerfile.frontend -t web-front:latest webserver/ && \
	docker tag web-front:latest knabben/web-front:latest

frontend-push:
	@make frontend-build
	docker push knabben/web-front:latest

local-frontend-run:
	yarn --cwd webserver start

## Back-end targets
backend-build:
	docker build -f docker/Dockerfile.backend -t web-backend:latest backend && \
	docker tag web-backend:latest knabben/web-backend:latest

backend-push:
	@make backend-build
	docker push knabben/web-backend:latest

local-backend-run:
	cd backend; pipenv run python manage.py runserver 0.0.0.0:8092

## Web assets Production deploy
web-deploy:
	kubectl create -f k8s/web-deploy.yaml

web-clean:
	kubectl delete -f k8s/web-deploy.yaml

## Operator targets
local-ops-generate:
	cd term-operator && operator-sdk generate k8s

local-ops-build:
	cd term-operator && GOOS=darwin GOARCH=amd64 operator-sdk build knabben/ops

local-ops-run:
	@make local-ops-build
	@make ops-deploy
	kubectl apply -f term-operator/deploy/cr.yaml
	cd term-operator && TELEMETRY_HOST=localhost:8092 tmp/_output/bin/term-operator

ops-build:
	cd term-operator && GOOS=linux GOARCH=amd64 operator-sdk build knabben/ops

ops-push:
	docker push knabben/ops:latest

ops-deploy:
	kubectl apply -f term-operator/deploy/rbac.yaml && \
	kubectl apply -f term-operator/deploy/crd.yaml

ops-deploy-op:
	@make ops-build
	kubectl apply -f term-operator/deploy/operator.yaml

ops-clean:
	rm term-operator/tmp/_output/bin/term-operator && \
	cd term-operator && \
	kubectl delete -f deploy/rbac.yaml && \
	kubectl delete -f deploy/crd.yaml && \
	kubectl delete -f deploy/operator.yaml
