#!/bin/bash

problem=$1

if [ ! $# -eq 1 ]; then
  echo "usage: $0 problem"
  exit 1
fi

if [ ! -e $problem ]; then
  mkdir -p $problem
  cp ../templates/template.go ./$problem/main.go
else
  echo "$problem is already exists."
fi
