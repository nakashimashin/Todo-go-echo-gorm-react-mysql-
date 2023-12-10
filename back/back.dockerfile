FROM golang:1.21


WORKDIR /back

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]

EXPOSE 8080