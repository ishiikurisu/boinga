default: test

test:
	go test

wasm:
	GOARCH=wasm GOOS=js go build -o lib.wasm ./main
