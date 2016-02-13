#!/bin/bash
echo "Write the configuration file..."

# Create the project.name file:
echo ${projectName} > project.name

# Create the configuration.json file which points to the configuration database:
echo '{' > configuration.json
echo '"ConfigDBHostname" : "'${dbHost}'",' >> configuration.json
echo '"ConfigDBDatabase" : "Re4EEE",' >> configuration.json
echo '"ConfigDBConfigurationCollection" : "Configuration",' >> configuration.json
echo '"ConfigDBConfigurationCollectionUsername" : "'${dbUser}'",' >> configuration.json
echo '"ConfigDBConfigurationCollectionPassword" : "'${dbPassword}'"' >> configuration.json
echo '}' >> configuration.json

echo "Write the configuration file... done."