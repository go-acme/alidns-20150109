#!/bin/bash -e

SRC_ORG='alibabacloud-go'
DEST_ORG='go-acme'

SRC_REPO_NAME='alidns-20150109'
DEST_REPO_NAME='alidns-20150109'

#LIB_VERSION=v4.5.0
LIB_VERSION=$(curl -s https://api.github.com/repos/${SRC_ORG}/${SRC_REPO_NAME}/releases/latest | jq -r '.tag_name')

SRC_REMOTE="git@github.com:${SRC_ORG}/${SRC_REPO_NAME}.git"
DEST_REMOTE="git@github.com:${DEST_ORG}/${DEST_REPO_NAME}.git"

DEST_BRANCH=modifiedclient

SRC_DIR=$(mktemp -d)
DEST_DIR=$(mktemp -d)

#############

## Fake fork remote

# DEST_REMOTE=$(mktemp -d)
#
# git init -q --bare ${DEST_REMOTE}
#
# DEST_TEMP=$(mktemp -d)
# git clone -q ${DEST_REMOTE} ${DEST_TEMP}
#
# cd ${DEST_TEMP}
# git switch -q -c ${DEST_BRANCH}
# git commit -q -m "Initial empty commit" --allow-empty
# git push -q -u origin ${DEST_BRANCH}
# cd ..
#
# rm -rf ${DEST_TEMP}

## Prepare the fork
# git clone -q --single-branch git@github.com:${DEST_ORG}/${DEST_REPO_NAME}.git /tmp/${DEST_REPO_NAME}
# cd /tmp/${DEST_REPO_NAME}
# git checkout --orphan ${DEST_BRANCH}
# git rm -rf .
# git commit -m "chore: initial empty commit." --allow-empty
# git push origin ${DEST_BRANCH}
# exit 0

#############

## Clone original repository

rm -rf ${SRC_DIR}

git clone -c advice.detachedHead=false -q --branch ${LIB_VERSION} --single-branch --depth 1 "${SRC_REMOTE}" ${SRC_DIR}

rm -rf ${SRC_DIR}/.git

## Clone destination repository

rm -rf ${DEST_DIR}

git clone -q --branch ${DEST_BRANCH} --single-branch ${DEST_REMOTE} ${DEST_DIR}

cd ${DEST_DIR}

## Remove all files
git rm -f -r --ignore-unmatch '*'

## Copy the code from the sources
cp -r ${SRC_DIR}/. .

## Change module name
go mod edit -module github.com/${DEST_ORG}/${DEST_REPO_NAME}/v5

## Convert the code
sed -E '

# --- Transform receiver (client *Client) to parameter ---

s|\(client \*Client\) ([^(]+)\(|\1(client *Client, |

# --- Transform method call to function call ---

s|\bclient\.([a-zA-Z0-9]+)\(|\1(client,|

# --- Fixes ---

# CallApi must be a method
s|CallApi\(client,|client.CallApi(|

# Init method and NewClient constructor
s|_err = CheckConfig\(client,config\)|_err = client.CheckConfig(config)|
s|func Init\(client \*Client, |func (client *Client) Init(|
s|err := Init\(client,|err := client.Init(|

' client/client.go > client/modifiedclient.go

rm client/client.go

sed -E '

# --- Transform receiver (client *Client) to parameter ---

s|\(client \*Client\) ([^(]+)\(ctx context.Context, |\1(ctx context.Context, client *Client, |

' client/client_context_func.go > client/modifiedclient_context_func.go

rm client/client_context_func.go

## Check compilation
go mod tidy
golangci-lint fmt -Egofmt
go build ./client/
rm go.sum

## Commit and Push

git add .
git commit -q -m "feat: update to ${LIB_VERSION}"

git push -q origin ${DEST_BRANCH}
git tag ${LIB_VERSION}
git push -q origin ${LIB_VERSION}

cd ..

rm -rf ${SRC_DIR}
rm -rf ${DEST_DIR}

##########################

# echo ${DEST_DIR}
#
# rm -rf ${DEST_REMOTE}
#
# cd /home/ldez/sources/go-acme/lego
#
# go mod edit -dropreplace github.com/alibabacloud-go/alidns-20150109/v5
# go mod edit -replace github.com/alibabacloud-go/alidns-20150109/v5=${DEST_DIR}
# go mod tidy
