# Installation

First you will need to clone this repo.

## Run with docker

You need to provide a running mongodb server.

	docker run --name re4eee --link MY_MONGODB -d -p 80:40000 -p 50000:50000 -e "dbHost=MY_MONGODB" -e "dbPassword=MY_PASSWORD" -e "dbUser=Re4EEE" re4eee

## Run with docker-compose
The docker-compose script will set up a mongodb automatically.
Change `ROOT_PASS` and `USER_PASS` in `docker-compose.yml` to something more safe.

	docker-compose up -d --build
