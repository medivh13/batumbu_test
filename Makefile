BINARY=linkaja

test: 
	go test -v -cover -covermode=atomic ./...

build:
	go build -o batumbu_test

run:
	./batumbu_test

build-run:
	make build
	make run

loadtest:
	race-the-web test.toml	

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t batumbu_test .

docker-run:
	docker-compose up --build -d

stop:
	docker-compose down