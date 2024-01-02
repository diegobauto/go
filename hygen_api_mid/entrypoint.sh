
#!/usr/bin/env bash

set -e
set -u
set -o pipefail

#if [ -n "${PARAMETER_STORE:-}" ]; then
#  export HYGEN_API_MID_PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}hygen_api_mid/db/username --output text --query Parameter.Value)"
#  export HYGEN_API_MID_PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/hygen_api_mid/db/password --output text --query Parameter.Value)"

exec ./main "$@"