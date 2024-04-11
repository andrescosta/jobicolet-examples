.PHONE: build docker deploy undeploy test logs

release: build docker job

build:
	cargo build --release --target wasm32-unknown-unknown

docker:
	docker cp target/wasm32-unknown-unknown/release/greet.wasm jobico-control-plane:/data/volumes/pv1/wasm

job:
	kubectl apply -f k8s-greet.yml

job-delete: 
	kubectl delete -f k8s-greet.yml

test:
	./ev-k8s.sh

logs:
	kubectl logs -f -l 'app=exec, owner=job-for-rust'
