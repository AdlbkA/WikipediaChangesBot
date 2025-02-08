FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o wikipedia_bot ./
RUN ls -l wikipedia_bot
RUN chmod +x wikipedia_bot

EXPOSE 8080

CMD ["./wikipedia_bot"]