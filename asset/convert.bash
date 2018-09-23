s=45
for f in *.svg; do
	inkscape -z -w $s -h $s $f -e $(echo $f | sed 's/.svg/.png/g') &
done
