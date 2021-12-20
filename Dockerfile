FROM golang:alpine

WORKDIR /src/app

COPY . .

RUN env GIN_MODE=release go build

CMD ["/src/app/issue-bot"]
