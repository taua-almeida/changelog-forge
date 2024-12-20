#!/bin/bash

set -e

EXPECTED_FILE="expected_changelog.json"
CHANGED_FILE="changelog.json"

if ! cmp -s "$EXPECTED_FILE" "$CHANGED_FILE"; then
  echo "changelog.json does not match the expected output."
  echo "Please run 'changelog-forge --generate-json' and try again."
  exit 1
fi

echo "changelog.json matches the expected output."
