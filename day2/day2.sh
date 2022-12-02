#!/bin/sh

part1() {
  draws=$(grep -E "A X|B Y|C Z" input.txt | wc -l)
  wins=$(grep -E "B Z|C X|A Y" input.txt | wc -l)
  shapes=$(sed "s/X/1/g;s/Y/2/g;s/Z/3/g;s/[A-C]\s//g" input.txt | paste -sd+ | bc)
  echo $((draws*3 + wins*6 + shapes))
}

part2() {
  draws=$(grep "Y" input.txt | wc -l)
  wins=$(grep "Z" input.txt | wc -l)
  shapes=$(sed "s/A X/3/;s/A Y/1/;s/A Z/2/;s/B X/1/;s/B Y/2/;s/B Z/3/;s/C X/2/;s/C Y/3/;s/C Z/1/" input.txt | paste -sd+ | bc)
  echo $((draws*3 + wins*6 + shapes))
}

part1
part2


