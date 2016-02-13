#!/bin/bash
echo "Configure the customerDB and loggingDB..."

# CustomerDB hostname and port:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBHost'},{\$set: {Value: '${dbHost}'}});"

# CustomerDB database name:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBDatabase'},{\$set: {Value: 'Re4EEE'}});"

# CustomerDB username:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBUsername'},{\$set: {Value: '${dbUser}'}});"

# CustomerDB password:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBPassword'},{\$set: {Value: '${dbPassword}'}});"

# LoggingDB hostname and port:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBHost'},{\$set: {Value: '${dbHost}'}});"

# LoggingDB database name:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBDatabase'},{\$set: {Value: 'Re4EEE'}});"

# LoggingDB username:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBUsername'},{\$set: {Value: '${dbUser}'}});"

# LoggingDB password:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBPassword'},{\$set: {Value: '${dbPassword}'}});"

# Activate and bind the admin interface to all network cards to port 50000:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'AdminWebServerBinding'},{\$set: {Value: ':50000'}});"
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'AdminWebServerEnabled'},{\$set: {Value: 'true'}});"

# Set the public web server's port to 80:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'PublicWebServerPort'},{\$set: {Value: '40000'}});"

# Configure the logging devices:
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogUseConsoleLogging'},{\$set: {Value: 'true'}});"
mongo Re4EEE --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogUseDatabaseLogging'},{\$set: {Value: 'true'}});"

echo "Configure the customerDB and loggingDB... done."