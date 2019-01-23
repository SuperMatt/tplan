DOCKER_TAG=tplan_build
DOCKER_NAME=tplan_build

itest_%: itest

itest: build_docker package

build_docker:
	[ -d ./dist ] || mkdir ./dist
	docker build -t $(DOCKER_TAG) -f Docker/Dockerfile .

package:
	docker run --name $(DOCKER_NAME) -t -v  $(CURDIR):/work:rw $(DOCKER_TAG):latest /build.sh

clean:
	rm -rf bin dist
	rm -rf go
	docker stop $(DOCKER_NAME)
	docker rm $(DOCKER_NAME)
	docker rmi $(DOCKER_TAG)