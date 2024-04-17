# if building manually first run:
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go
# and then build Docker image

FROM scratch
COPY manager /
ENTRYPOINT ["/manager"]
