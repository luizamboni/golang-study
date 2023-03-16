# to run in a docker i have to build for linux
GOOS=linux
#
GOARCH=arm64
# Enables statically linked binaries to make the application more portable. It allows you to use the binary with source images that don't support shared libraries when building your container image.
CGO_ENABLED=1

build:
	go build  -o $(shell pwd)/time $(shell pwd)/14-time/main.go
	docker build -t golang-study --no-cache .

run: build
	docker run \
	-v $(shell pwd):/apps/ \
	golang-study \
	./time

# run: build
# 	docker run \
# 	-v $(shell pwd):/apps/ \
# 	-e TZ=America/Sao_Paulo \
# 	golang-study \
# 	./time