for f in *.svg; do
	inkscape -z -w 64 -h 64 $f -e $(echo $f | sed 's/.svg/.png/g')
done
