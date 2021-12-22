FROM golang

RUN useradd --no-log-init --create-home --user-group --uid 1000 test-user

USER 1000:1000

COPY . .

RUN go build -o app main.go

ENTRYPOINT ./app
