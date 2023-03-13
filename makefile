test:
	go test -v -cover ./...

build_run:
	docker build -t test-lending-app . && docker run --name test-lending-app test-lending-app

down:
	docker stop test-lending-app && docker rm test-lending-app && docker rmi -f test-lending-app