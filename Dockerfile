FROM golang:latest

RUN mkdir -p /go/src/final
RUN mkdir -p /opt/bin/
RUN mkdir -p /opt/bin/gw
RUN mkdir -p /opt/bin/news
RUN mkdir -p /opt/bin/comment
RUN mkdir -p /opt/bin/censor
RUN mkdir -p /opt/logs
RUN go env -w GO111MODULE=auto

WORKDIR /go/src/final
ADD apigw/  ./apigw/
ADD internal/ ./internal/
ADD censor/ ./censor/
ADD comment/ ./comment/
ADD news/ ./news/
ADD go.mod .
#ADD --chmod=755  run_all.sh  /opt/bin/
RUN go mod  tidy
# RUN ls -al ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build   -o /opt/bin/gw/gwcmd  ./apigw
ADD  apigw/config.json /opt/bin/gw/
ADD  apigw/webapp /opt/bin/gw/webapp
ADD  apigw/docs /opt/bin/gw/docs
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build   -o /opt/bin/news/newscmd   ./news/cmd
ADD  news/cmd/config.json /opt/bin/news/
RUN  GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build   -o /opt/bin/comment/commentcmd   ./comment/cmd
ADD  comment/cmd/config.json /opt/bin/comment/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0   go build   -o /opt/bin/censor/censorcmd   ./censor
ADD  censor/config.json /opt/bin/censor
#RUN ls -al /opt/bin/
# ENTRYPOINT  [ "/bin/bash -c /opt/bin/run_all.sh"]  
# CMD ["bash", "/opt/bin/run_all.sh"]
# EXPOSE 80 11200 11001 12035


 FROM alpine:latest
LABEL version="1.0.0"
LABEL maintainer="Taradaev Andrey<ataradaev@gmail.com>"
ENV CENSORDB=postgres://<user>:<password>@<host>:<port>/<DBcensor>
ENV COMMENTDB=postgres://<user>:<password>@<host>:<port>/<DBcomments>
ENV NEWSDB=postgres://<user>:<password>@<host>:<port>/<BDnews>
WORKDIR /opt/
RUN mkdir -p /opt/gw
RUN mkdir -p /opt/news
RUN mkdir -p /opt/comment
RUN mkdir -p /opt/censor
RUN mkdir -p /opt/logs
COPY --from=0  /opt/bin/  /opt
#COPY --from=0 /opt/bin/news/  /opt/news
#COPY --from=0 /opt/bin/comment/  /opt/comment
#COPY --from=0 /opt/bin/censor/  /opt/censor
#COPY --from=0 /opt/bin/censor/  /opt/censor
ADD --chmod=755  run_all.sh  /opt/
RUN chmod +x /opt/news/newscmd
CMD ["sh", "/opt/run_all.sh"]
EXPOSE 80 11200 11001 12035