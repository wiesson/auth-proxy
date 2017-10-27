## Auth Reverse Proxy

### Install

```
go get github.com/wiesson/auth-proxy
cd $GOPATH/src/github.com/wiesson/auth-proxy
go install
```

### arguments

Example

```
auth-proxy -host https://api.myawesomecompany.com/v2/ -token my-super-secret-token
```

#### -host*

example: `-host https://api.myawesomecompany.com/v2/`

#### -token*

example: `-token my-super-secret-token`
