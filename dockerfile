FROM golang:1.22

WORKDIR /user/app_go/product

COPY . .

ENV CONFIG_PATH="./config/config.yml"

RUN cd /user/app_go/product && go mod download && go build ./cmd/product/main.go

CMD [ "./main" ]

EXPOSE 8081
