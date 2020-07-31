#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname $0)
echo ${SCRIPT_DIR}
cd ${SCRIPT_DIR} && cd ../src/adapter/web/gqlmodel

# https://github.com/vektah/dataloaden
#go get -u github.com/vektah/dataloaden@v0.3.0
dataloaden TextContentLoader int64 []*github.com/sky0621/wht/adapter/web/gqlmodel.TextContent
dataloaden ImageContentLoader int64 []*github.com/sky0621/wht/adapter/web/gqlmodel.ImageContent
dataloaden VoiceContentLoader int64 []*github.com/sky0621/wht/adapter/web/gqlmodel.VoiceContent
dataloaden MovieContentLoader int64 []*github.com/sky0621/wht/adapter/web/gqlmodel.MovieContent
