FROM --platform=arm64 alpine:3.15

WORKDIR /apps
COPY . /apps/
# ENV TZ=Europe/Amsterdam
ENV TZ=America/Sao_Paulo

RUN apk add tzdata
