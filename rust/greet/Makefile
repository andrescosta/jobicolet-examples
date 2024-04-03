## Dependencies WSL ():
## curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
## cc: sudo apt install build-essential
## cargo install wasm-pack
## rustup target add wasm32-unknown-unknown

.PHONE: docker deploy undeploy test logs

docker:
	docker cp target/wasm32-unknown-unknown/release/greet.wasm jobico-control-plane:/data/volumes/pv1/wasm

deploy:
	kubectl apply -f k8s-greet.yml

undeploy: 
	kubectl delete -f k8s-greet.yml

test:
	./ev-k8s.sh

logs:
	kubectl logs -f $(pod)