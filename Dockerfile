FROM golang:1.22 as builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make

FROM alpine
WORKDIR /
COPY --from=builder /workspace/server .
ENTRYPOINT ["/server"]
