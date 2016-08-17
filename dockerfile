# Use a Ubuntu server as basis:
FROM ubuntu:14.04

# Set the gopath in order to use Go:
ENV GOPATH /go

# Update the operating system and install base tools e.g. Git:
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y git wget zip && \
	
	# Install the MongoDB clients:
	apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv EA312927 && \
	echo "deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.2.list && \
	apt-get update && \
	apt-get install -y mongodb-org-tools mongodb-org-shell && \

	# Create the Go workspace:
	mkdir /go && \
	mkdir /go/src && \
	mkdir /go/bin && \
	mkdir /go/pkg && \
	cd /go && \

	# Install current Go:
	wget --no-check-certificate -O go.tar.gz https://storage.googleapis.com/golang/go1.5.3.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go.tar.gz && \
	rm go.tar.gz

# Create the project directories:	
RUN cd /go && \
	mkdir src/github.com && \
	mkdir src/github.com/SommerEngineering && \
	mkdir src/github.com/SommerEngineering/Re4EEE

# Define some variables that the user can adjust for every container:
ENV dbHost=127.0.0.1 \
	dbPassword= \
	dbUser=Re4EEE \
	projectName=Re4EEE

# Set up the project:
RUN export PATH=$PATH:/usr/local/go/bin && \

	# Install libraries for Re4EEE and Ocean:
	cd /go/src/github.com/SommerEngineering/Re4EEE && \

	# Ocean:
	go get github.com/SommerEngineering/Ocean && \

	# Generator for UUIDs:
	go get github.com/twinj/uuid && \
	
	# Database driver:
	go get gopkg.in/mgo.v2 

# Insert all files from the repo (but from the current directory, not from Git):
ADD / /go/src/github.com/SommerEngineering/Re4EEE/	
	
# Compile and Setup
RUN	export PATH=$PATH:/usr/local/go/bin && \
	cd /go/src/github.com/SommerEngineering/Re4EEE && \

	# Compile the Re4EEE:
	go install && \

	# Copy the final binary and the runtime scripts to the home folder:
	cp /go/bin/Re4EEE /home && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/run.sh /home/run.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/setConfiguration.sh /home/setConfiguration.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/configureCustomerDB.sh /home/configureCustomerDB.sh && \
	cp /go/src/github.com/SommerEngineering/Re4EEE/uploadStaticData.sh /home/uploadStaticData.sh && \

	# Zip static data and move them to the home folder:
	cd staticFiles && \
	zip -r /home/staticFiles.zip . && \
	cd ../templates && \
	zip -r /home/templates.zip . && \
	cd ../web && \
	zip -r /home/web.zip . && \
	cd .. && \

	# Uninstall Git and tools:
	apt-get remove -y git wget zip && \
	apt-get autoremove -y && \

	# Delete Go:
	rm -r -f /usr/local/go && \

	# Delete the entire Go workspace:
	cd /home && \
	rm -r -f /go && \

	# Create the configuration file:
	touch /home/configuration.json && \
	touch /home/project.name && \

	# Make the scripts executable:
	chmod 0777 /home/run.sh && \
	chmod 0777 /home/setConfiguration.sh && \
	chmod 0777 /home/configureCustomerDB.sh && \
	chmod 0777 /home/uploadStaticData.sh && \
	chmod 0777 /home/Re4EEE && \
	chmod 0666 /home/configuration.json && \
	chmod 0666 /home/project.name

# Run anything below as nobody:
USER nobody

# Re4EEE provides HTTP by port 40000 and the admin interface on port 50000:
EXPOSE 40000 50000

# Define the working directory:
WORKDIR /home

# The default command to run, if a container starts:
CMD ./run.sh
