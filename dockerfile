# RUN go run main.go
FROM golang:1.20 

# Remove other microservices' folders

# RUN find /app/src/app/ -mindepth 1 -maxdepth 1 -type d ! -name $MICROSERVICE -exec rm -r {} \;

WORKDIR /app

COPY . /app/stock_broker_application


ARG MICROSERVICE="authentication"
ARG PORT=8080

# RUN rm -rf /stock_broker_application/src/app/!$MICROSERVICE

WORKDIR /app/stock_broker_application

RUN  echo "MICROSERVICE is set to: $MICROSERVICE"


WORKDIR /app/stock_broker_application/src/app

# RUN ls | grep -v "^$MICROSERVICE$" | xargs rm -rf

WORKDIR /app/stock_broker_application/src/app/$MICROSERVICE

RUN go install .


# # SWAG installation
# # RUN curl -o /usr/local/bin/Swag https://github.com/linuxserver/docker-Swag/releases/latest/download/Swag && chmod +x /usr/local/bin/Swag

# # Build the Go app
RUN go build -o main .
# # Swagger update
# # RUN /usr/local/bin/Swag init -g main.go
EXPOSE $PORT

CMD ["./main"]
