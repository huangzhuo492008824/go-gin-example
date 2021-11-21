FROM scratch

WORKDIR $GOPATH/src/github.com/huangzhuo492008824/go-gin-example
COPY ./go-gin-example $GOPATH/src/github.com/huangzhuo492008824/go-gin-example/
COPY ./conf/app.ini $GOPATH/src/github.com/huangzhuo492008824/go-gin-example/conf/app.ini

EXPOSE 8000
CMD ["./go-gin-example"]