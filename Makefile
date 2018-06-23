.PHONY: deploy

export KUBERNETES_CONFIG := ${HOME}/.kube/config
export WATCH_NAMESPACE := default

web-run:
	cd webserver && \
	yarn start && \
	python manage.py runserver 0.0.0.0:8000

web-build:
	cd webserver && \
	docker build -f Dockerfile.dj -t web-ops:latest . && \
	docker build -f Dockerfile.fr -t web-front:latest . && \
	docker tag web-ops:latest knabben/web-ops:latest && \
	docker tag web-front:latest knabben/web-front:latest

web-push:
	docker push knabben/web-ops:latest
	docker push knabben/web-front:latest

web-deploy:
	kubectl create -f webserver/deploy/webserver.yaml

web-clean:
	kubectl delete -f webserver/deploy/webserver.yaml

generate:
	operator-sdk generate k8s

build:
	cd term-operator && operator-sdk build knabben/ops

ops-deploy:
	@make build
	kubectl apply -f term-operator/deploy/rbac.yaml && \
	kubectl apply -f term-operator/deploy/crd.yaml && \
	kubectl apply -f term-operator/deploy/operator.yaml

ops-clean:
	cd term-operator && \
	kubectl delete -f deploy/rbac.yaml && \
	kubectl delete -f deploy/crd.yaml && \
	kubectl delete -f deploy/operator.yaml

ops-run:
	@make build
	cd term-operator && tmp/_output/bin/term-operator
