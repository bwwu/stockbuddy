FROM golang:1.20

WORKDIR $GOPATH/src/github.com/bwwu/stockbuddy
COPY . .

RUN go build .

# Set this via docker build --build-arg secret=<VALUE>
ARG secret
ENV STOCKBUDDY_PASSWORD=$secret

CMD ["./stockbuddy", "--mail_to", "brandonwu23@gmail.com"]