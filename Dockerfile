from golang:1.13.8 as builder

COPY  . ./src/app
WORKDIR /go/src/app
RUN mkdir bin/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/ifood-oh

FROM alpine:3.8
COPY --from=builder /go/src/app/bin/ifood-oh /app/ifood-oh
WORKDIR /app
CMD /app/ifood-oh