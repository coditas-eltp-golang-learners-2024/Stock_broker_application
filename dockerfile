# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY ./ /app

RUN pwd

#List the files in the directory
RUN ls -a

# Set the environment variable
ARG MICROSERVICE=authentication
ARG PORT= 8080


# Remove other microservices' folders
#RUN find /app/src/app/ -mindepth 1 -maxdepth 1 -type d ! -name "$MICROSERVICE" -exec rm -r {} \;


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
WORKDIR /app/src/app/$MICROSERVICE

# Copy go mod and sum files
RUN go mod tidy

RUN pwd
RUN ls -a /app/src/app/$MICROSERVICE


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go install .


#Expose port to outside world
EXPOSE $PORT

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
