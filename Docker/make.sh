#!/usr/bin/env bash
echo "*************************************************"
echo "*    Download binary & resources for secpass    *"
echo "*  and generate Dockerfile & docker-compose.yml *"
echo "*                     @2019                     *"
echo "*************************************************"

# check version variable
if [ -z "$1" ]; then
   echo "You should provide a version variable."
   echo -e 'For example: \033[1m ./make.sh v1.0.3\033[0m'
   echo "Exiting...."
   exit 1
fi

# make Docker file
echo -e "\033[1mMake Dockerfile:\033[0m"
echo "docker build -t sptnk/secpass:latest ." > ./1-build.sh && chmod +x ./1-build.sh

echo -e '# Version: 0.0.1
FROM alpine
MAINTAINER Aleksey Ovchinnikov <alexovchinnicov@google.com>

RUN mkdir -p /opt/secpass

WORKDIR /opt/secpass

COPY ./resources.zip /opt/secpass
COPY ./secpass-linux-amd64 /opt/secpass

RUN cd /opt/secpass && unzip resources.zip && rm resources.zip

RUN sed -i "s/127.0.0.1/secpassdb/g" /opt/secpass/resources/config.json
RUN sed -i "s/5432/12215/g" /opt/secpass/resources/config.json

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 2215

ENTRYPOINT ["/opt/secpass/secpass-linux-amd64"]
' > Dockerfile
echo ok

# make Docker-compose file
echo -e "\033[1mMake Docker-compose file:\033[0m"
echo "docker-compose up --build -d" > ./2-build-compose.sh && chmod +x ./2-build-compose.sh
echo -e 'version: "3.3"
services:
    secpassdb:
        image: postgres:alpine
        container_name: secpassdb
        network_mode: bridge
        volumes:
          - "/var/lib/secpass/data:/var/lib/postgresql/data"
        environment:
          - POSTGRES_USER=secpass
          - POSTGRES_PASSWORD=secpass
          - POSTGRES_DB=secpass
        ports:
          - 127.0.0.1:12255:5432
        restart: always

    secpass:
        depends_on:
            - secpassdb
        links:
            - secpassdb
        image: sptnk/secpass
        container_name: secpass
        network_mode: bridge
        ports:
            - 2215:2215
' > docker-compose.yml
echo ok

# download using wget utility
echo -e "\033[1mDownload:\033[0m"
wget -q --show-progress https://github.com/ltheinrich/secpass/releases/download/$1/resources.zip
wget -q --show-progress https://github.com/ltheinrich/secpass/releases/download/$1/secpass-linux-amd64 && chmod +x ./secpass-linux-amd64

# list files
echo -e "\033[1mNew Files:\033[0m"
ls -1

# notes
echo -e "\033[1mNotes:\033[0m"
echo -e 'Run the \033[1m./1-build.sh && ./2-build-compose.sh\033[0m files one by one.'
echo -e 'Then open URL \033[1mhttps://{{dockerhost}}:2215\033[0m in you browser.'
echo ""
echo -e "\033[1mGood luck!\033[0m"
echo ""
