# Start from the latest golang base image
FROM golang:1.21.3

LABEL author="Mauricio Milano"
LABEL description="jobsity challenge"
LABEL version="1.0"


WORKDIR /app/
# COPY go.mod ./go.mod
COPY go.sum go.sum ./ 

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV APP_POSTGRES_HOST='postgres'
ENV APP_POSTGRES_PORT='5432'
ENV APP_POSTGRES_DB='chat'
ENV APP_POSTGRES_USER='app'
ENV APP_POSTGRES_PASS='APP'
ENV APP_PORT="3000"

ENV PUBLISHER_QUEUE='pqueue'
ENV RECEIVER_QUEUE='rqueue'
ENV RMQ_HOST='rabbitmq'
ENV RMQ_USERNAME='admin'
ENV RMQ_PASSWORD='admin'
ENV RMQ_PORT='5672'
ENV JWT_SECRET='password_very_strong'
ENV JWT_TTL='10000'



RUN go mod download

EXPOSE 3000

CMD ["go","run","main.go"]