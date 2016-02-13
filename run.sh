#!/bin/bash

# Set the ConfigurationDB settings:
echo "Try to write the configuration file:"
./setConfiguration.sh

# Try to upload the static data:
echo "Try to upload the static files in to the database:"
./uploadStaticData.sh

# Try to start the app:
./Re4EEE

OUT=$?
if [ $OUT -eq 1 ];then
   echo "Error while connecting to the configuration database. Is the dbHost parameter (${dbHost}) correct?"
elif [ $OUT -eq 2 ]; then
   echo "Cannot connect to the customer database: Reconfigure the configuration database now..."
   
   # Configure the customerDB and loggingDB:
   echo "Try to configure the customerDB and loggingDB:"
   ./configureCustomerDB.sh
   
   # Try to upload the static data again:
   echo "Try to upload the static files in to the database:"
   ./uploadStaticData.sh
   
   # Start the app:
   ./Re4EEE
fi