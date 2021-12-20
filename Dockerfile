FROM golang:alpine

WORKDIR /src/app

COPY . .

RUN env GIN_MODE=release go build
RUN cp issue-bot /usr/bin/

CMD ["/usr/bin/issue-bot"]
