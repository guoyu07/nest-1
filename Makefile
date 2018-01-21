build:
	go build

docker: build
	docker build -t reg.qiniu.com/wolfogre/nest:${version} .
	docker push reg.qiniu.com/wolfogre/nest:${version}
