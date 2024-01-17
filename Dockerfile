FROM golang:1.20.6-alpine3.18

RUN mkdir /online-store
WORKDIR /online-store

COPY go.mod /online-store

COPY . /online-store 

RUN go mod tidy 

RUN go build /online-store/main.go

EXPOSE 3000

CMD [ "/online-store/main" ]