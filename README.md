golang webservice to store and retrieve messages.
# Requirements
MongoDB 2.6+
Golang 1.10+
Docker 18.10+

# Configurations
Sample configurations are enumerated in the drophash/config/drophash_local_config.json file. Currently supports modifications to mongodb server host and port. You may also customize the listening port of the drophash service. You must write the file to an appropriate location and denote the path in the DROPHASH_LOCAL_CONFIGURATION_PATH defined in main.go Init().

When running in docker environment, you will need to confirm the DROPHASH_DOCKER_CONFIGURATION_PATH is being used within the main.go Init() function and that you are satisfied with those configurations before building docker image (below).

When local_mem_mode is set to 1, the database connection will not be established, instead local memory will be used to hold entries (see Important section).
database_name is used to configure the mongo database queried. collection_name is the set of data queried within the database. It is recommended tht you familiarize yourself with mongodb convention to successfully stand-up a drophash server.


# Quickstart (running on local machine)
open a command line and run
```
go get github.com/GeorgeLuo/drophash
```
Have a mongo database server running and listening.

Direct yourself to your $GOROOT path where your github src are located (~/go/src/github.com/), and cd into the drophash directory.
```
go run main.go
```

test a message storing POST
```
curl   -X POST   http://localhost:8080/messages   -H 'Content-Type: application/json'   -d '{"message":"hello world"}'
```
test a message retrieving GET
```
curl   http://localhost:8080/messages/[digest value]
```
# Running with Docker
open a command line and run
```
go get github.com/GeorgeLuo/drophash
```
Direct yourself to your $GOROOT path where your github src are located (~/go/src/github.com/), and cd into the drophash directory.

pull the official docker image
```
docker pull golang
```
build the docker image
```
build -t drophash .
```
Have a mongo database server running and listening.

start the drophash container, with port 6060 on the local host directed to port 8080 within the container
```
docker run --publish 6060:8080 --name test drophash
```

You now have a docker container that is listening on localhost:6060 and functions like the localized server.

clean your environment with the following when you wish to terminate the container
```
docker stop test
docker rm $(docker ps -a -f status=exited -q)
```
# Important
If database configuration is not set, the service will default to in-memory mode. This is not advisable in a distributed environment across multiple hosts, there will be no guarantee of correct behavior and data will be unrecoverable after the process is terminated.