FROM --platform=linux/amd64 golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
COPY main.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-server main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/go-server /go-server
EXPOSE 8080
CMD ["./go-server"]
