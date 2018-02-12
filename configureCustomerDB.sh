#!/bin/bash
echo "Configure the customerDB and loggingDB..."

# CustomerDB hostname and port:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBHost'},{\$set: {Value: '${dbHost}'}});"

# CustomerDB database name:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBDatabase'},{\$set: {Value: 'ejob-o-mat'}});"

# CustomerDB username:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBUsername'},{\$set: {Value: '${dbUser}'}});"

# CustomerDB password:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'CustomerDBPassword'},{\$set: {Value: '${dbPassword}'}});"

# LoggingDB hostname and port:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBHost'},{\$set: {Value: '${dbHost}'}});"

# LoggingDB database name:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBDatabase'},{\$set: {Value: 'ejob-o-mat'}});"

# LoggingDB username:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBUsername'},{\$set: {Value: '${dbUser}'}});"

# LoggingDB password:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogDBPassword'},{\$set: {Value: '${dbPassword}'}});"

# Activate and bind the admin interface to all network cards to port 50000:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'AdminWebServerBinding'},{\$set: {Value: ':50000'}});"
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'AdminWebServerEnabled'},{\$set: {Value: 'true'}});"

# Set the public web server's port to 80:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'PublicWebServerPort'},{\$set: {Value: '40000'}});"

# Configure the logging devices:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogUseConsoleLogging'},{\$set: {Value: 'true'}});"
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'LogUseDatabaseLogging'},{\$set: {Value: 'true'}});"

# Configure site verification token:
mongo ejob-o-mat --host ${dbHost} --username ${dbUser} --password ${dbPassword} --eval "db.Configuration.update({Name: 'SiteVerificationToken'},{\$set: {Value: '${SITE_VERIFICATION_TOKEN}'}});"

echo "Configure the customerDB and loggingDB... done."