# 如果go run .



### 有ERROR



```
go run .
../github.com/go-playground/validator/baked_in.go:22:2: cannot find package "golang.org/x/crypto/sha3" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/golang.org/x/crypto/sha3 (from $GOROOT)
	/Users/tsengyenchi/go/src/golang.org/x/crypto/sha3 (from $GOPATH)
../github.com/go-playground/validator/baked_in.go:23:2: cannot find package "golang.org/x/text/language" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/golang.org/x/text/language (from $GOROOT)
	/Users/tsengyenchi/go/src/golang.org/x/text/language (from $GOPATH)
../github.com/gin-gonic/gin/binding/toml.go:12:2: cannot find package "github.com/pelletier/go-toml/v2" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/github.com/pelletier/go-toml/v2 (from $GOROOT)
	/Users/tsengyenchi/go/src/github.com/pelletier/go-toml/v2 (from $GOPATH)
../github.com/gin-gonic/gin/binding/msgpack.go:14:2: cannot find package "github.com/ugorji/go/codec" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/github.com/ugorji/go/codec (from $GOROOT)
	/Users/tsengyenchi/go/src/github.com/ugorji/go/codec (from $GOPATH)
../github.com/gin-gonic/gin/binding/protobuf.go:12:2: cannot find package "google.golang.org/protobuf/proto" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/google.golang.org/protobuf/proto (from $GOROOT)
	/Users/tsengyenchi/go/src/google.golang.org/protobuf/proto (from $GOPATH)
../github.com/gin-gonic/gin/binding/yaml.go:12:2: cannot find package "gopkg.in/yaml.v3" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/gopkg.in/yaml.v3 (from $GOROOT)
	/Users/tsengyenchi/go/src/gopkg.in/yaml.v3 (from $GOPATH)
../github.com/gin-gonic/gin/logger.go:14:2: cannot find package "github.com/mattn/go-isatty" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/github.com/mattn/go-isatty (from $GOROOT)
	/Users/tsengyenchi/go/src/github.com/mattn/go-isatty (from $GOPATH)
../golang.org/x/net/idna/idna10.0.0.go:25:2: cannot find package "golang.org/x/text/secure/bidirule" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/golang.org/x/text/secure/bidirule (from $GOROOT)
	/Users/tsengyenchi/go/src/golang.org/x/text/secure/bidirule (from $GOPATH)
../golang.org/x/net/idna/idna10.0.0.go:26:2: cannot find package "golang.org/x/text/unicode/bidi" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/golang.org/x/text/unicode/bidi (from $GOROOT)
	/Users/tsengyenchi/go/src/golang.org/x/text/unicode/bidi (from $GOPATH)
../golang.org/x/net/idna/idna10.0.0.go:27:2: cannot find package "golang.org/x/text/unicode/norm" in any of:
	/usr/local/Cellar/go/1.21.3/libexec/src/golang.org/x/text/unicode/norm (from $GOROOT)
	/Users/tsengyenchi/go/src/golang.org/x/text/unicode/norm (from $GOPATH)

```

## 解法

```
go get golang.org/x/crypto/sha3
go get golang.org/x/text/language
....
..
.

```


半自動建立laravel
