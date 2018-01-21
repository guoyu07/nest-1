build:
	go build

docker build: build
	docker build -t reg.qiniu.com/wolfogre/nest:${version} .
	docekr push -t reg.qiniu.com/wolfogre/nest:${version}