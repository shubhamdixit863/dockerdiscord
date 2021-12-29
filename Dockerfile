FROM golang:1.17

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/shubhamdixit863/discordgo

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
ENV token=OTIzODMzMjA2MzYzNTM3NDI5.YcVwuA.QOyVM2afeV-nGQPGGNsGlQa1ShA

# Run the executable
CMD ["discordgo"]