appname := warmish
version := 0.1-beta
config := config.yml
target := build

sources := $(wildcard *.go)

build = GOOS=$(1) GOARCH=$(2) go build -o ${target}/$(appname)$(3)
tar = cd ${target} && tar -cvzf ${appname}-${version}-$(1)_$(2).tar.gz $(appname)$(3) ${config} && rm $(appname)$(3)
zip = cd ${target} && zip ${appname}-${version}-$(1)_$(2).zip $(appname)$(3) ${config} && rm $(appname)$(3)

.PHONY: all package windows darwin linux clean

all: package windows darwin linux purge

package:
	mkdir -p ${target}
	cp ${config} ${target}/

clean:
	rm -rf ${target}/

purge:
	rm ${target}/${config}

##### LINUX BUILDS #####
linux: build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_386.tar.gz build/linux_amd64.tar.gz

build/linux_386.tar.gz: $(sources)
	$(call build,linux,386,)
	$(call tar,linux,386)

build/linux_amd64.tar.gz: $(sources)
	$(call build,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm.tar.gz: $(sources)
	$(call build,linux,arm,)
	$(call tar,linux,arm)

build/linux_arm64.tar.gz: $(sources)
	$(call build,linux,arm64,)
	$(call tar,linux,arm64)

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_amd64.tar.gz

build/darwin_amd64.tar.gz: $(sources)
	$(call build,darwin,amd64,)
	$(call tar,darwin,amd64)

##### WINDOWS BUILDS #####
windows: build/windows_386.zip build/windows_amd64.zip

build/windows_386.zip: $(sources)
	$(call build,windows,386,.exe)
	$(call zip,windows,386,.exe)

build/windows_amd64.zip: $(sources)
	$(call build,windows,amd64,.exe)
	$(call zip,windows,amd64,.exe)
