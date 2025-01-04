#! /usr/bin/bash

if [ $# -lt 3 ]; then
    echo "Error: At least 3 arguments are required" >&2
    exit 1
fi

prev=$(pwd)
next=$(dirname "$(realpath "$0")")
cd "$next"

year=$1
day=$2
input=$3
additional_args=${@:4}

path=$(find ${year}/${day}* -maxdepth 0)
cmd="go run ./${path} ${path}/${input} ${additional_args}"
echo -e "\n>> $cmd <<\n"
eval "$cmd"

cd "$prev"
