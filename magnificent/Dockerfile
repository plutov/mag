FROM python:3.8

WORKDIR /go/src/github.com/plutov/mag/magnificent

COPY . .

RUN pip install twisted

EXPOSE 12345

ENTRYPOINT ["python", "-u", "server.py"]