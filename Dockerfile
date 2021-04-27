FROM golang
WORKDIR api
COPY . .
EXPOSE 3000
RUN go build -o bin/api cmd/api/main.go
ENTRYPOINT "bin/api"

