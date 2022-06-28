FROM golang:1.18

WORKDIR /go/src/github.com/gweithio/gowlfer/

ADD . .

# Get Deps, run tests and build binary
RUN make all

ENTRYPOINT ["./gowlfer"]
