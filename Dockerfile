FROM golang:1.10 as builder

# install dep
RUN go get github.com/golang/dep/cmd/dep
# create a working directory
WORKDIR /go/src/github.com/br-parser
# copy Gopkg.toml and Gopkg.lock
COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock
# install packages
RUN dep ensure --vendor-only
# copy source code
COPY . .
# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go
RUN go test -v ./...
# use a minimal alpine image
FROM alpine:3.7
# set working directory
WORKDIR /root
RUN mkdir /resources
# copy files from builder
COPY --from=builder /go/src/github.com/br-parser/main .
COPY --from=builder /go/src/github.com/br-parser/resources/ ./resources
# expose port
EXPOSE 8080
# run the binary
CMD ["./main"]

