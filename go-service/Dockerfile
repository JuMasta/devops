FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0  go build -ldflags="-w -s" -o /app/service 

FROM scratch 

WORKDIR /

COPY --from=build /app/service /

EXPOSE 8080

EXPOSE 8081

CMD [ "/service" ]
