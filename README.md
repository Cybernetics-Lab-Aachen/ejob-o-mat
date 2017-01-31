# About
This repository contains all sources for the e-learning recommendation system elearning-finder.net and elearn-o-mat.net. This work is part of the project "ELLI – Excellent Teaching and Learning in Engineering Sciences" and was funded by the federal ministry of education and research, Germany.

# Installation
To set up the project, a clone of the repository is necessary. Afterwards, two options for running the recommendation system are available: Run as Docker container or run as Docker Compose service. In order to run the system as Docker container, a MongoDB database instance must be ready. For the Docker Compose solution, this precondition is not necessary: The necessary database gets started automatically.

## Run with Docker Container
[![Docker Automated buil](https://img.shields.io/docker/automated/jrottenberg/ffmpeg.svg)](https://hub.docker.com/r/elli2/re4eee/)

First of all, the Docker container must be build. In order to do so, perform the following command:

```
docker build -t re4eee .
```

In order to run the system as Docker container, perform the following command:

```
docker run --name re4eee --link MY_MONGODB --restart=always -d -p 80:40000 -p 127.0.0.1:50000:50000 -e "dbHost=MY_MONGODB" -e "dbPassword=MY_PASSWORD" -e "dbUser=Re4EEE" re4eee
```

Here, the MongoDB database can be also run as Docker container or an external database can be used. Please complement the above command with the database’s properties. After starting the recommendation system, the public facing site runs at port `80`. Further, the internal administration site runs on port `50000`, but local to the Docker machine.

## Run with Docker Compose
In order to run the system with Docker Compose, perform the following command from the base directory of this repository:

```
docker-compose up -d --build
```

Please change `ROOT_PASS` and `USER_PASS` in `docker-compose.yml` locally to provide a safe password. After starting the recommendation system, the public facing site runs at port `80`. Further, the internal administration site runs on port `50000`, but local to the Docker machine.

## Updating
After pulling the latest version the update process differs on whether Docker or Docker Compose is used to run the system.
The database does not need to be touched during update.

### Updating with Docker Container
Stop and remove the current container. All data is kept in the database and thus safe.
```
docker rm -f re4eee
```
Afterwards build the new container and start it as described above under Run with Docker Compose.

### Updating with Docker Compose
Updating with Docker Compose is very straightforward and works the same way as installing:
```
docker-compose up -d --build
```

# License
All sources are available under a BSD 2-clause license and can be used for any purposes, also commercial usage. The author does not provide any support.

# Contact
In order to get in contact with the author, please write a mail to thorsten.sommer@ima-zlw-ifu.rwth-aachen.de Please notice, that the author cannot provide support.
