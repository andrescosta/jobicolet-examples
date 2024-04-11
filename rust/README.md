Dependencies WSL:

- Jobico for K8s
- Rust
- wasm-pack
- Docker

## curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
## cc: sudo apt install build-essential
## cargo install wasm-pack
## rustup target add wasm32-unknown-unknown

# Deploying Jobico

git clone https://github.com/andrescosta/jobicok8s.git
cd jobicok8s
make kind
wait certificates
make deploy-all

git clone https://github.com/andrescosta/jobicolet-examples

cd jobicolet-examples/rust/greet
make release
make test
make logs

# Test:

1- Compile

make build

2- Upload to K8s

make docker

3- Create a Job 

make job

4- Test

make test

5- Logs

make logs
