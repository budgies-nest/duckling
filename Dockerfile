FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .

RUN <<EOF
go mod download
go mod tidy 
go build -o duckling
EOF

FROM scratch


WORKDIR /app
COPY --from=builder /app/duckling .
CMD ["./duckling"]