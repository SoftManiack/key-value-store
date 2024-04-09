FROM golang:1.16

# Copy the source files from the host
COPY . .

# Set the working directory to the same place we copied the code
WORKDIR /src

RUN go install -v ./...

EXPOSE 8080

