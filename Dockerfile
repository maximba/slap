FROM golang:alpine
WORKDIR /go/src/slap
COPY vendor vendor
COPY *.go ./
RUN CGO_ENABLED=0 go build -o slap

FROM scratch
COPY --from=0 /go/src/slap /
EXPOSE 8080
ENTRYPOINT ["/slap"]
