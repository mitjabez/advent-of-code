#!/bin/sh
echo "Elf with most calories:"
paste -sd+ input.txt | sed s/++/\\n/g | bc | sort -nr | head -n1

echo "Cals of top 3 Elves:"
paste -sd+ input.txt | sed s/++/\\n/g | bc | sort -nr | head -n3 | paste -sd+ | bc
