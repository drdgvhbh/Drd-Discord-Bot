FROM golang:1.11
ARG gosource=/go/src/
WORKDIR ${gosource}
ARG root=./
ARG directory=drdgvhbh/discordbot/
RUN mkdir -p ${directory}
COPY . ${directory}
WORKDIR ${gosource}${directory}
RUN make clean
RUN make build
CMD ./exec/wait && make run
