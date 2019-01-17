FROM golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY parseAmount.go config.json findRange.py ./
RUN go build -o libparseAmount.so -buildmode=c-shared parseAmount.go

FROM python:3
WORKDIR /usr/src/app
COPY --from=0 /usr/src/app .
#RUN python findRange.py -c ./config.json -i "SGD 12.34"