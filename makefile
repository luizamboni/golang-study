
build:
	docker build -t golang-study .

run:
	docker run \
	-v $(shell pwd):/apps/ \
	golang-study \
	go run /apps/13-time/main.go