#! /bin/bash

set -e

out=game/levels.go
echo "package game; var AllLevels = []string{" > $out

for f in levels/*.level; do
	echo -n "\`" >> $out
	cat $f >> $out
	echo "\`," >> $out
done;

echo "}" >> $out
gofmt -w $out
