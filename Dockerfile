FROM golang:1.16.4-alpine

ENV wallet=/go/src/tools

COPY . $wallet/
RUN apk update && apk add gcc && apk add g++ \
&& cd $wallet/ \
&& go build -o tron-tools

FROM alpine

ENV wallet=/go/src/wallet
COPY --from=0  $wallet/tron-tools /usr/bin
WORKDIR /data

CMD ["tron-tools"]
