#!/bin/bash
# https://circleci.com/docs/orbs-best-practices/#accepting-parameters-as-strings-or-environment-variables
PATH=$(circleci env subst "${PARAM_PATH}")
CONTENT=$(circleci env subst "${PARAM_CONTENT}")
ORG=$(circleci env subst "${PARAM_ORG}")
REPO=$(circleci env subst "${PARAM_REPO}")
COMMENT_ID=$(circleci env subst "${PARAM_COMMENT_ID}")

commenter update -o $ORG -r $REPO -i $COMMENT_ID -c $CONTENT -f $PATH
