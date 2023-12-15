- follow [this blog](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin#h-goals)

# Preparation

``` shell
# project directory
cd your_work_space
mkdir go-web-application-with-gin && cd "$_"

# go mod
go mod init github.com/your_user_name/your_path
# we use gin here
go get -u github.com/gin-gonic/gin
```

# developing

1. router for index.html 

- create `template/index.html`
- create `main.go`
- add `gin engine` usage codes

then run app
``` shell
go build -o app && ./app
# default listen in http://localhost:8080/
```

