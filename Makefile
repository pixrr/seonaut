.PHONY: run test linux docker

run:
	go run cmd/server/main.go

test:
	go test github.com/mnlg/seonaut/test

docker:
	docker run -p 6306:3306 --name crawler-mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

linux:
	GOOS=linux GOARCH=amd64 go build -o build/seonaut cmd/server/main.go
