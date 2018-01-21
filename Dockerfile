FROM centos:7

RUN rm /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY nest /opt/

COPY tmpl/ /opt/tmpl/

VOLUME /opt/static

EXPOSE 80

WORKDIR /opt

ENTRYPOINT ["./nest"]
