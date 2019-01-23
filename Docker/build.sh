#! /bin/bash

set -xe

apt-get update && apt-get install -y software-properties-common git
add-apt-repository -y ppa:gophers/archive && apt-get update && apt-get install -y golang-1.10-go
gem install fpm

cd /work/
PATH="${PATH}:/usr/lib/go-1.10/bin"
export GOPATH=/work/go
mkdir -p /work/go/src/github.com/hashicorp/terraform
git clone --single-branch --branch v${VERSION} https://github.com/hashicorp/terraform/ /work/go/src/github.com/hashicorp/terraform
 
go test ./tplan
go build -o bin/tplan-${VERSION} ./tplan

cd /work/dist
fpm -s dir -t deb --deb-no-default-config-files --name tplan --prefix=/usr/bin \
    --version ${VERSION} --iteration ${ITERATION} ../bin

apt install /work/dist/tplan_${VERSION}-${ITERATION}_amd64.deb
which tplan-${VERSION}
tplan-${VERSION} -version