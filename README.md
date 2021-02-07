# batumbu backend test

## Prerequisite:

1. This project requires Golang to run, view installation instruction here https://golang.org/doc/install
2. set up golang to local environment, view the complete instruction here https://github.com/golang/go/wiki/SettingGOPATH or here https://www.tutorialspoint.com/go/go_environment.htm
3. Make sure Docker and Docker compose is already installed in your local computer as this step is required to run this project on Docker
4. Install race-the-web(https://github.com/TheHackerDev/race-the-web) on your local computer, this tools is required to perform load test on this project. to install, run on terminal: go get github.com/TheHackerDev/race-the-web

## How to run:

This project depedency is managed by go mod, therefore you can put it anywhere other than GO-PATH directory.
steps to run using docker-compose:
    1. cd to project root directory
    2. run on terminal : 
        docker-compose up --build -d
    or by using Makefile, run on terminal : 
        make docker-run
    3. verify the container is running by run:
        docker ps
    4. Execute the call:
        curl localhost:8080/invest/a/5
    5. To Stop run on terminal:
        make stop

## How to test:

   a. Unit testing
    to perform unit test with coverage info run on terminal:
        make test
    b. Load testing
    verify race-the-web is installed on your local computer and before performing this test you need to verify that this service is running and ready to accept request, then run on terminal:
        make loadtest
    

## API

    Available endpoints on this project
    a. to view invest retutn                   -> [GET] {host}/invest/:invest_type/:period 
        eg : ocalhost:8080/invest/a/5

