FROM alpine:3.5
LABEL maintainer="Jack Toumey"

COPY ./simple-go /app/simple-go

RUN chmod +x /app/simple-go

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/simple-go
