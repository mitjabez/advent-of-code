#!/bin/bash
set -e

readonly ME=$(basename "$0")
readonly MY_DIR="$(dirname "$0")"

usage() {
  cat <<EOF
Usage: $ME <year> <no of day>
Creates a folder for and Advent of Code day and downloads in the inputs.
AOC_COOKIE environment variable should be set so the script can download the inputs.

Options:
    year:           Year of the assignment.
    no of day:      Day of the assignment (1-25).
EOF
}

log() {
  echo "INFO: $*" 2>&1
}

die() {
  echo "ERROR: $*" 2>&1
  exit 1
}

main() {
  if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
    usage
    exit 0
  fi

  local year="$1"
  local day="$2"

  [[ "$day" =~ ^[0-9][0-9]?$ ]] || die "Day should be a number"
  if [ "$day" -lt 1 ] || [[ "$day" -gt 25 ]]; then
    die "Day should be a number between 1 and 25"
  fi
  [[ "$year" =~ ^[2][0][1-3][0-9]?$ ]] || die "Invalid year"

  day_dir="$MY_DIR/../$year/day$day"
  [ -d "$day_dir" ] && die "Directory $day_dir already exists"

  log "Creating day$day"
  mkdir "$day_dir"
  cp "$MY_DIR/template.go" "$day_dir/main.go"


  base_url="https://adventofcode.com/$year/day/$day"

  log "Downloading input ..."
  curl -sS -H "Cookie: $AOC_COOKIE" "$base_url/input" -o "$day_dir/input"
  log "Downloading input.test ..."
  curl -sS -H "Accept: text/html" -H "Cookie: $AOC_COOKIE" "$base_url" |
    grep -m 1 -A 40 "<pre><code>" |
    grep -m 1 -B 20 "</pre>" |
    sed "s/<\/*pre>//;s/<\/*code>//" |
    grep -v "^$" > "$day_dir/input.test"


}

main "$@"
