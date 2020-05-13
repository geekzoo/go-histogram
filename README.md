# go-histogram

1. Construct histrogram of time seriese data using golang.
1. Graph increase and decreased values or likeness.

Usage: `cat mydata.txt | go run histrogram.go`

Build: `go build`

Run: `cat mydata.txt | ./go-histogram`

Other: `for i in {000000..999999}; do echo "${i} $(( $RANDOM % 240 ))"; done | go run histrogram.go`

data example

```data
2020/01/28 2
2020/01/29 16
2020/01/30 72
2020/02/01 10
```

```data
2020/01/28 ∎∎
2020/01/29 ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
2020/01/30 ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
2020/02/01 ∎∎∎∎∎∎∎∎∎∎
```

- [x] read data
- [x] graph data
- [ ] Add color to >=< to mesure changes by day