#!/bin/bash
echo "Upload the static files in to the database..."

# Add the latest web.zip to the database:
mongofiles --host ${dbHost} --username ${dbUser} --password ${dbPassword} --db Re4EEE --local /home/web.zip --replace put web.zip

# Add the latest staticFiles.zip to the database:
mongofiles --host ${dbHost} --username ${dbUser} --password ${dbPassword} --db Re4EEE --local /home/staticFiles.zip --replace put staticFiles.zip

# Add the latest templates.zip to the database:
mongofiles --host ${dbHost} --username ${dbUser} --password ${dbPassword} --db Re4EEE --local /home/templates.zip --replace put templates.zip

echo "Upload the static files in to the database... done."