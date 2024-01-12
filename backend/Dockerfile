FROM golang:1.21-alpine AS sigma-backend

WORKDIR /

# встановлення додаткових інструментів та бібліотек
RUN apk add gcc libc-dev

# встановлення залежностей
COPY go.mod go.sum ./
RUN go mod download

# копіювання основного коду
COPY . ./

# build
RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -a .

# підготовка фінального образу
FROM scratch
EXPOSE 8080
COPY --from=sigma-backend sigma-test .
CMD ["./sigma-test"]