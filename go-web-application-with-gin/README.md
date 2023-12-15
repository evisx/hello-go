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

## basic page and router

1. [✨ router for index\.html](https://github.com/evisx/golang-demos/commit/7bc96b0e28337383ab08f7cc9cfc9a19c22e6ebd)

- create `template/index.html`
- create `main.go`
- add `gin engine` usage codes

then run app
``` shell
go build -o app && ./app
# default listen in http://localhost:8080/
```

2. [♻ extract header and footer template and use bootstrap](https://github.com/evisx/golang-demos/commit/d872026c1c1c93e0a6f1b068933d443c66685ab0)

> header, menu, sidebar, and footer, which will be common across all pages

- extract header and footer base on index.html
- use bootstrap

3. [✨ add navigation menu](https://github.com/evisx/golang-demos/commit/97e9ecdf8c1f9d584248f1d93b5abb5bc63301f6) 

- create menu.html
- embed template at header.html

4. look back, currently we have:

``` text
├── main.go
└── templates
    ├── footer.html
    ├── header.html
    ├── index.html
    └── menu.html
```

## article model

> model, handler and view

1. [♻ extract showIndexPage handler and render](https://github.com/evisx/golang-demos/commit/5f7c097df0f498fc25ac5b9397837d4d1e81e79c)

- extract handler, html render

2. [✨ article model](https://github.com/evisx/golang-demos/commit/d5cb89bbd98ae55c6c2d061a534fadd383d1ab06)

- type article struct
- example articles
- test getAllArticles

3. index for all example articles

4. article view and handler