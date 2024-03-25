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

3. [✨ index list all articles](https://github.com/evisx/golang-demos/commit/96914404d867c2fb5e8a79bef9544f9226fbefb6)

- render payload

## unit test router handler

1. [✨ unitest showIndePage](https://github.com/evisx/golang-demos/commit/a8efadb2b3acdca7b070ec8c73bd682cbd0a94b6)
- test router handler by http requesting (use testing, httptest) for showIndexPage
    - extract getRuouterWithInitialized from main
    - statusOK: should https status ok
    - pageOK: should include some special html tag
    - simulate http request and response by router serving http

2. extract common test helper function
- [♻ extract common test helper](https://github.com/evisx/golang-demos/commit/ae4d0f4be0f6666bc50ccc24da1a27c09ddfc01a)
- [♻ extract simpleHTTPResponse](https://github.com/evisx/golang-demos/commit/e57f4177ffb11d08ccfd7c3461a82fbdd3520fc2)

## article view 

1. display single simple article
- [✨ article view](https://github.com/evisx/golang-demos/commit/29843844dfde8c005751f71dd536b9a185c4f661)
- [♿ view article by id](https://github.com/evisx/golang-demos/commit/654d5531b4d987eef639741d27e315062dd3299e)

## more response format

- [✨ response format by requset header of Accept](https://github.com/evisx/golang-demos/commit/1c560dfbb2790d30d1deb3d0343656c547af48c3)
    1. get request header: Accept
    2. render right format for request
        - `application/json` will render JSON
        - `application/xml` will render XML
        - otherwise default render HTML
   
``` shell
curl -X GET -H "Accept: application/json" http://localhost:8080/

curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1
```
## Reference

- Mostly follow [this blog](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin#h-goals)
