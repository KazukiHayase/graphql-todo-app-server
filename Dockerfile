FROM golang:1.21.0

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest \
  && go install github.com/99designs/gqlgen@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

CMD ["air", "-c", ".air.toml"]
