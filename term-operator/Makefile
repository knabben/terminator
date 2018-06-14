generate:
	operator-sdk generate k8s

build:
	operator-sdk build knabben/ops

.PHONY: deploy
deploy:
	kubectl apply -f deploy/rbac.yaml
	kubectl apply -f deploy/crd.yaml
	kubectl apply -f deploy/operator.yaml

run:
	@make build
	KUBERNETES_CONFIG=/home/vagrant/.kube/config WATCH_NAMESPACE=default tmp/_output/bin/term-operator
