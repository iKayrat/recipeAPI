FROM golang:1.18

LABEL stage=builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /app

ADD go.mod ./
ADD go.sum ./
RUN go mod download
COPY ./ ./

# RUN go install recipeApi


# Build the Go application
RUN go build -ldflags="-s -w" -o ./recipeAPI cmd/main.go

CMD /recipeAPI

CMD ["./recipeAPI"]

# EXPOSE 8000