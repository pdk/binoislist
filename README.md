# binoislist
A template and bash script for generating types of list of things

This is just a simple example of one way to use generators to create "generic"
code.

Here are some examples that can be triggered with `go generate ./...`:

```
//go:generate $GOPATH/src/github.com/pdk/binoislist/make-binois-list.sh animal_list.go animal Animal AnimalsList Animal{}
//go:generate $GOPATH/src/github.com/pdk/binoislist/make-binois-pointer-list.sh user_list.go models User UserList
//go:generate $GOPATH/src/github.com/pdk/binoislist/make-binois-pointer-list.sh account_list.go models Account AccountList
```

In the subfolders `animal` and `models` there are examples of generated code.
Since this "library" is mainly meant to be used as a template, there are no real
executable programs here, except for the two bash scripts that use `sed` to
generate the target `.go` files.