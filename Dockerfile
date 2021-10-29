FROM centos

MAINTAINER 972568589@qq.com

WORKDIR /home

add cmd/main /home

CMD ["/home/main"]

EXPOSE 8080