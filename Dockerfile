FROM golang:alpine as builder
WORKDIR /go/src/slap
COPY vendor vendor
COPY *.go ./
RUN CGO_ENABLED=0 go build -o slap

FROM scratch
COPY --from=builder /go/src/slap /
EXPOSE 8080
ENTRYPOINT ["/slap"]
