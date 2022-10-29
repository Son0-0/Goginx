# Goginx

<img width="600" alt="스크린샷 2022-10-29 오후 4 10 45" src="https://user-images.githubusercontent.com/81317358/198818960-750ee8fc-6dec-4608-939b-2c9a39c2f121.png">

```
.
├── README.md
├── go.mod
├── handlers
│   └── reverseProxy.go
├── main.go
├── target1.go
└── target2.go
1 directories, 6 files
```

## Description
- Reverse Proxy 기능 (L7 Load Balancing)
    - 현재 간단하 REST API 정도만 동작
    - Socket 통신, Redirect 등 동작 X

## Installation
1. 레포지토리 클론  
```git clone https://github.com/Son0-0/Goginx.git```
2. Goginx 실행 (Settings의 포트 포워딩 완료 후 실행)
    - 2.1 컴파일 후 실행
        - ```go build main.go && ./main```
    - 2.2 실행
        - ```go run main.go```

## Settings

```
Server: AWS EC2 t2.micro (Ubuntu 20.04 LTS) 기준
```

1. 443 포트 8443 포트로 포트포워딩  
```sudo iptables -t nat -A PREROUTING -p tcp --dport 443 -j REDIRECT --to-port 8443```
2. Reverse Proxy 설정

``` go
// main.go

// localhost:8080/target1/xxx 경로로 들어올 경우 8081 포트로 Request를 보내줌
target1Handler := &handlers.PortNumHandler{PortNum: "8081"}
http.HandleFunc("/target1/", target1Handler.Handler)

// localhost:8080/target2/xxx 경로로 들어올 경우 8082 포트로 Request를 보내줌
target2Handler := &handlers.PortNumHandler{PortNum: "8082"}
http.HandleFunc("/target2/", target2Handler.Handler)
```

Nginx 설정 파일은 다음과 같이 설정할 수 있다.

``` nginx
server {
  listen       443 ssl;
  server_name  example.com;
    ssl_certificate cert.pem
    ssl_certificate_key key.pem

  location /target1 {
    proxy_pass http://localhost:8081;
  }

  location /target2 {
    proxy_pass http://localhost:8082;
  }
}
```

## Test

1. origin server 실행  
- ```go run target1.go```
- ```go run target2.go```

2. Goginx 실행  
```go run main.go```

Request via curl
```
curl -i localhost:8080/target1/home
```

Response
```
HTTP/1.1 200 OK
Server: Goginx (Linux)
Date: Sat, 29 Oct 2022 06:57:32 GMT
Content-Length: 23
Content-Type: text/plain; charset=utf-8

target1 server response%
```
