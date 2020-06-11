FROM alpine:3.10.3

RUN apk --no-cache add python3 py-pip ca-certificates

COPY . /app

WORKDIR /app

RUN pip3 install -r requirements.txt

ENV FLASK_APP=src/server

EXPOSE 80

ENTRYPOINT [ "flask" ]

CMD [ "run", "-h", "0.0.0.0", "-p", "80" ]