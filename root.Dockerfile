FROM golang

COPY . .

RUN go build -o app main.go

ENTRYPOINT ./app
