FROM golang:1.13

WORKDIR /herd
COPY . .
RUN sed -e /provider/d -i go.mod
RUN make HERD_TAGS=no_third_party
