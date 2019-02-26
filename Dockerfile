FROM golang

ADD . /go/src/github.com/GeorgeLuo/drophash

# Install deps
RUN go get github.com/labstack/echo
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/sirupsen/logrus
RUN go get github.com/mongodb/mongo-go-driver/mongo/options
RUN go get github.com/mongodb/mongo-go-driver/mongo
RUN go get github.com/labstack/echo/middleware
RUN go get gopkg.in/mgo.v2/bson

# Build 
RUN go install github.com/GeorgeLuo/drophash

# Set default run command
ENTRYPOINT /go/bin/drophash

# Expose port 8080.
EXPOSE 8080