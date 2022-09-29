FROM golang:1.18.1-bullseye as build

ENV PACKAGES bash curl

WORKDIR /build

RUN apt update
RUN apt install -y $PACKAGES

RUN curl https://get.ignite.com/cli?type=script | bash && chown root:root ignite

WORKDIR /code
COPY . /code/

RUN /build/ignite chain build

FROM debian:10.11-slim as node
WORKDIR /validator
VOLUME [ "/nodes/node1/" ]
VOLUME [ "/nodes/node2/" ]
COPY --from=build  /go/bin/xchaind /validator/

# RUN chmod 777 /validator/xchaind
ENTRYPOINT ["/validator/xchaind",  "start", "--home", "/nodes/node1"]   


