FROM golang:1.9.2-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/kataras/iris
RUN go get github.com/lib/pq
RUN go get github.com/joho/godotenv
RUN go get github.com/gorilla/securecookie
RUN go get github.com/golang/protobuf/protoc-gen-go

ENV SOURCES /go/src/github.com/marlonpeterc/product-microservice/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}product-microservice
EXPOSE 8090

