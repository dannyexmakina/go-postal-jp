FROM centos:7

RUN yum install -y epel-release wget zip unzip git bind-utils which

WORKDIR /usr/local/src
RUN wget https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.10.1.linux-amd64.tar.gz

RUN yum install -y gcc sqlite

RUN mkdir -p /root/cli /var/www/html /var/go/src /var/go/pkg /var/go/bin

ENV GOPATH /var/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go get -u github.com/golang/dep/cmd/dep

COPY ./root/run /root/
#RUN chmod 755 /root/run

VOLUME ["/var/www/html/app","/var/go/src","/root/cli"]


#EXPOSE 80 8080 8000

ENTRYPOINT ["/root/run"]
