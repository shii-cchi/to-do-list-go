FROM golang:1.22.5-alpine
RUN apk add --no-cache make
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
ARG DB_URI
ENV DB_URI=$DB_URI
WORKDIR to-do-list-go
COPY . .
RUN go mod download
RUN go build -o todo_server cmd/main.go
CMD ["sh", "-c", "goose -dir internal/database/migrations postgres ${DB_URI} up && ./todo_server"]