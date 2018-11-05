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
RUN sed -i 's/DB_HOST=[a-z|A-Z]*/DB_HOST=postgres/g' .env
RUN sed -i 's/DB_PORT=[0-9]*/DB_PORT=5432/g' .env
CMD ./exec/wait && make run
