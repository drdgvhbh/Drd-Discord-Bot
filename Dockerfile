FROM golang:1.11
ARG gosource=/go/src/
WORKDIR ${gosource}
ARG root=./
ARG directory=drdgvhbh/discordbot/
RUN mkdir -p ${directory}
COPY . ${directory}
WORKDIR ${gosource}${directory}
RUN go get github.com/google/go-cloud/wire/cmd/wire
RUN make clean
RUN make build
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait
CMD /wait && ./bin/drdbot