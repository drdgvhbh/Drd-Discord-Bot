FROM golang:1.11
WORKDIR /go/src
RUN mkdir -p Drd-Discord-Bot
COPY .env ./Drd-Discord-Bot
COPY bot ./Drd-Discord-Bot/bot
WORKDIR /go/src/Drd-Discord-Bot/bot
RUN go get github.com/google/go-cloud/wire/cmd/wire
RUN make clean
RUN make build
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait
CMD /wait && ./bin/DiscordBot