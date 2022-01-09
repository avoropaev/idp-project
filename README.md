# IDP-PROJECT

## Run

```bash
make init
```

## Try

```bash
curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"query": "{hashcode (code:0)}"}' \
	http://localhost:8080/graphql
```

```bash
curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"jsonrpc": "2.0","method": "hash.calc","params": {"code": 1},"id": 123}' \
	http://localhost:8080/rpc
```

```bash
curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"jsonrpc": "2.0","method": "guid.generate","params": {"code": 2},"id": 123}' \
	http://localhost:8080/rpc
```

## Links
[http://localhost:8080/graphql/playground](http://localhost:8080/graphql/playground)

[http://localhost:8080/rpc](http://localhost:8080/rpc)