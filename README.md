# Installation

First you will need to clone this repo and set up a mongodb server.

## Run with docker

	docker run --name re4eee --link MY_MONGODB -d -p 80:40000 -p 50000:50000 -e "dbHost=MY_MONGODB" -e "dbPassword=MY_PASSWORD" -e "dbUser=Re4EEE" re4eee

## Run with docker-compose
Adjust the environment variables in `docker-compose.yml` to properly reflect your database setup.

	docker-compose up -d --build
