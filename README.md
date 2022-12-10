# querystring
A querystring parsing and stringifying  go library.  Golang序列化和解析querystring库.



## Features
- [x] stringify any to querystring
- [ ] parse querystring


## Using

```go
import (
  "github.com/feiin/querystring"
)

//obj support  struct , map , interface{}
qs, err := querystring.Marshal(obj)
fmt.Printf("qs:%s",qs) //a=b&c=d.....

```
