#!/bin/sh
paste -sd+ input.txt | sed s/++/\\n/g | bc | nl | sort -k2 -nr | head -n1
