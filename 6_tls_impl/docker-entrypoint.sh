#!/bin/bash
set -e

cp ca.crt certs

exec "$@"
