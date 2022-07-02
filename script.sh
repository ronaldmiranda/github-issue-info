#!/usr/bin/env bash

# https://github.com/google/material-design-icons

ORG=${1:?"Organization is required . Usage: ${0} <org> <repo>"}
REPO=${2:?"Repository is required . Usage: ${0} <org> <repo>"}

curl "https://api.github.com/repos/$ORG/$REPO/issues?state=all&per_page=100&page=1" |
jq -sr '.[] | {
  open: map(select(.state | contains ("open"))) | length,
  closed: map(select(.state | contains ("closed"))) | length
}'
