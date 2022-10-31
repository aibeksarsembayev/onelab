run-nodocker:
	go run ./cmd

build:
	docker build -t icrudtmpl:multistage -f Dockerfile.multistage .
	docker image prune

run:
	docker run -d -p 8080:9090 --rm --name ccrudtmpl icrudtmpl:multistage

stop:
	docker stop ccrudtmpl

postmanCLI:
	curl -o- "https://dl-cli.pstmn.io/install/linux64.sh" | sh
	postman login --with-api-key PMAK-635fb2d50ad15f15ff8a27d9-8e36333770a22f935ed56a9d818d7585f0
	postman collection run 19375706-65d78971-c745-4072-a409-5bc668b66a00

.DEFAULT_GOAL := build
