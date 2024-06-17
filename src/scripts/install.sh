#!/bin/bash
# https://circleci.com/docs/orbs-best-practices/#accepting-parameters-as-strings-or-environment-variables
PATH=$(circleci env subst "${PARAM_PATH}")

go build -o "${PATH}commenter"
chmod +x "${PATH}commenter"
