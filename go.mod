module github.com/nspcc-dev/neo-go

replace github.com/nspcc-dev/dbft => github.com/ixje/neo-go-legacy-dbft v0.0.0-20250115175722-57f3662027db

require (
	github.com/Workiva/go-datastructures v1.0.50
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/alicebob/miniredis v2.5.0+incompatible
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/dgraph-io/badger/v2 v2.0.3
	github.com/go-redis/redis v6.10.2+incompatible
	github.com/gomodule/redigo v1.9.2 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/mr-tron/base58 v1.1.2
	github.com/nspcc-dev/neofs-crypto v0.2.3
	github.com/nspcc-dev/rfc6979 v0.2.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1
	github.com/spaolacci/murmur3 v1.1.0
	github.com/stretchr/testify v1.8.4
	github.com/syndtr/goleveldb v0.0.0-20180307113352-169b1b37be73
	github.com/urfave/cli v1.20.0
	github.com/yuin/gopher-lua v1.1.1 // indirect
	go.etcd.io/bbolt v1.3.4
	go.uber.org/atomic v1.4.0
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0
	golang.org/x/tools v0.0.0-20180318012157-96caea41033d
	gopkg.in/yaml.v2 v2.2.4
)

go 1.13
