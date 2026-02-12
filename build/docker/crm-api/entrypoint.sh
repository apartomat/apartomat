#!/bin/sh
set -e

if [ "$1" = "migration" ]; then
  shift
  exec migration "$@"
fi

exec crm-api "$@"
