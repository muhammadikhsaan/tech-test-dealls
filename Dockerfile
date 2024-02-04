FROM golang:1.21

# Set environtment value 
ENV PORT 8111
ENV MODE RELEASE
ENV REQUEST_TIMEOUT 60
ENV CREDENTIAL_PATH ../app/.credentials
ENV COOKIE_SECURE true
ENV HASH_COST 8

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./
RUN go mod tidy
RUN go work sync

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/service dealls.test/cmd

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE ${PORT}/tcp

# Remove Unused file
RUN rm -rf /docker-compose.yaml
RUN rm -rf /README.md

# Run
CMD ["/build/service"]