.PHONY: deploy

export KUBERNETES_CONFIG := ${HOME}/.kube/config
export WATCH_NAMESPACE := default

web-run:
	cd terminator; yarn start &
	python terminator/manage.py runserver 0.0.0.0:8000

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
