#!/bin/bash
set -euo pipefail

LITESTREAM_PATH="${LITESTREAM_PATH:-gcs://tmcdev-misc/moderncrud.db}"

if [ "${ENABLE_LITESTREAM:-1}" != "" ]; then
  set -x
  litestream restore -o db.sqlite "${LITESTREAM_PATH}"
  exec litestream replicate -exec "moderncrud-server -v" ./db.sqlite "${LITESTREAM_PATH}" 
else
  exec moderncrud-server
fi
