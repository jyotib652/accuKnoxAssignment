# base go image
# FROM golang:1.18-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o dinerApp ./cmd/api

# RUN chmod +x /app/dinerApp

# # build a tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app/dinerApp /app

# CMD [ "/app/dinerApp" ]


FROM alpine:latest

RUN mkdir /app

COPY dinerApp /app

COPY log.json /log.json

CMD [ "/app/dinerApp" ]