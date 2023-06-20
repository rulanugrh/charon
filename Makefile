PATH_BINARY=output
BINARY_FILE=${PATH_BINARY}/output

build:
	go build -o ${BINARY_FILE} server.go

run:
	go build -p ${BINARY_FILE} server.go
	./${BINARY_FILE}

clean:
	go clean
	rm ${BINARY_FILE}

runDock:
	docker-compose up -d

cleanDock:
	docker-compose down