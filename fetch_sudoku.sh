#!/bin/sh

if [ $1 ]; then
	level=$1
else
	level=0
fi
year=`date +%Y`
mon=`date +%m`
day=`date +%d`

UA="Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36"
URL="http://sudokupuzzle.org/online2.php?nd=$level&y=$year&m=$mon&d=$day"

echo "URL: $URL"
curl --user-agent "$UA" -s "$URL" | grep 'tmda=' | sed "s/.*tmda='//g" | cut -c -81
