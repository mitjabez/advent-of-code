#!/bin/sh
paste -sd+ input.txt | sed s/++/\\n/g | bc | sort -nr | head -n1
paste -sd+ input.txt | sed s/++/\\n/g | bc | sort -nr | head -n3 | paste -sd+ | bc
