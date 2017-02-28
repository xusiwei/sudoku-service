#!/bin/sh

if [ !$1 ]; then
	echo "Usage: $0 level year mon day"
fi

level=$1
day=$2
mon=$3
year=$4

echo "$level $year/$mon/$day"

if [ ! $level ]; then
	level=0
fi

if [ ! $year ]; then
	year=`date +%Y`
fi

if [ ! $mon ]; then
	mon=`date +%m`
fi

if [ ! $day ]; then
	day=`date +%d`
fi

echo "$level $year/$mon/$day"

UA="Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36"
URL="http://sudokupuzzle.org/online2.php?nd=$level&y=$year&m=$mon&d=$day"

echo "URL: $URL"
curl --user-agent "$UA" -s "$URL" | grep 'tmda=' | sed "s/.*tmda='//g" | cut -c -81
