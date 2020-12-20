#!/bin/bash
ls -al && ls -al ${GOPATH}
mkdir -p ${GOPATH}/src/${LIBPKG}
mkdir -p ${GOPATH}/src/github.com/EC-Release/agent/watcher
mv src/${LIBPKG}/* ${GOPATH}/src/${LIBPKG}
cp -r ./watcher/* ${GOPATH}/src/github.com/EC-Release/agent/watcher
ls -al ${GOPATH}/src/github.com/EC-Release/agent/watcher

#copying packges..
mkdir -p ${GOPATH}/pkg/{linux_amd64_race,linux_amd64,darwin_amd64,windows_amd64,linux_arm}/${LIBPKG}

cp -r ./pkg/linux_amd64/${LIBPKG}/. ${GOPATH}/pkg/linux_amd64/${LIBPKG}/
cp -r ./pkg/darwin_amd64/${LIBPKG}/. ${GOPATH}/pkg/darwin_amd64/${LIBPKG}/
cp -r ./pkg/windows_amd64/${LIBPKG}/. ${GOPATH}/pkg/windows_amd64/${LIBPKG}/
cp -r ./pkg/linux_arm/${LIBPKG}/. ${GOPATH}/pkg/linux_arm/${LIBPKG}/

cp -r ./pkg/linux_amd64_race/${LIBPKG}/. ${GOPATH}/pkg/linux_amd64_race/${LIBPKG}/

chmod -R 755 ./
make
