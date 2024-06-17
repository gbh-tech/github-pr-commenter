#!/bin/bash
# https://circleci.com/docs/orbs-best-practices/#accepting-parameters-as-strings-or-environment-variables
PATH=$(circleci env subst "${PARAM_PATH}")
CONTENT=$(circleci env subst "${PARAM_CONTENT}")
ORG=$(circleci env subst "${PARAM_ORG}")
REPO=$(circleci env subst "${PARAM_REPO}")
PULL_ID=$(circleci env subst "${PARAM_PULL_ID}")

echo 'export COMMENT_ID=$(commenter get -o $ORG -r $REPO -p $PULL_ID -c $CONTENT -f $PATH)' >> $BASH_ENV
echo $COMMENT_ID
