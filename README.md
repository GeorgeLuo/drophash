golang webservice to store and retrieve messages.

# Approach and Design Decisions
The main.go file is the initialization point. The endpoints are denoted here and the environment directory contains the calculation and storage operations. I find go to be a light-weight and dependable language to use, with robust error handling. I chose mongodb as the database for its availablity in a community edition form and ease of installation, as well capabilities sharding capabilities should we move this project towards production.
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

Direct yourself to your $GOROOT path where your github src are located (~/go/src/github.com/), and cd into the drophash directory. Pull the go dependencies.
```
go get github.com/labstack/echo
go get github.com/dgrijalva/jwt-go
go get github.com/sirupsen/logrus
go get github.com/mongodb/mongo-go-driver/mongo/options
go get github.com/mongodb/mongo-go-driver/mongo
go get github.com/labstack/echo/middleware
go get gopkg.in/mgo.v2/bson
```
Run the program and start the server.
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

#Scaling and Other Future Considerations
Eventually the system will move towards SSL encryption of requests and responses due to the nature of the service. This will add additional overhead/latency. We would have to take a critical look at go implementation in docker environment and assess the validaty of moving forward with this language or docker as an environment. We should look towards distribution of the data and duplicate datasets. In the same line of thinking, we should assess the possibility of bucketing users and or hashing by key and sharding across databases to speed up database access. Depending on the userbase, we should look to availability across regions. Depending on traffic size, we may need to front-end the service with gatekeeping mechanisms, perhaps by number of queries per IP. In addition, we may refactor the service to front-end requests and separate hosts/containers to perform calculations. In addition, we should look to implement logging to a common logger (instead of main out) which is reaped by a central process to aggregate metrics.