FROM golang:1.21


WORKDIR /back

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN apt-get update && apt-get install -y lsof

EXPOSE 8080

CMD ["go", "run", "main.go"]
