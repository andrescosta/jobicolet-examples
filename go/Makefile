.PHONY: install-job uninstall-job install-file logs

install-file:
	docker cp greet/greet.wasm jobico-control-plane:/data/volumes/pv1/wasm/greet-go.wasm

install-job:
	docker apply -f greet/k8s-greet.yml

uninstall-job:
	docker apply -f greet/k8s-greet.yml

test:
	greet/ev-k8s.sh

logs:
	kubectl logs -f -l='app=exec, owner=job-for-go'