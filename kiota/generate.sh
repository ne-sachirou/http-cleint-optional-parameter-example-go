#!/bin/bash
set -eux
docker run -v "$(pwd):/app/output" -v "$(pwd)/../openapi.yaml:/app/openapi.yaml" mcr.microsoft.com/openapi/kiota generate -l Go -n kiota/client -o client
