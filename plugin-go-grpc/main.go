package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dihedron/plugins/log"
	"github.com/dihedron/plugins/shared"
	"github.com/hashicorp/go-plugin"
	"go.uber.org/zap"
)

// Here is a real implementation of KV that writes to a local file with
// the key name and the contents are the value of the key.
type KV struct{}

func (KV) Put(key string, value []byte) error {
	value = []byte(fmt.Sprintf("%s\n\nWritten from plugin-go-grpc", string(value)))
	return ioutil.WriteFile("kv_"+key, value, 0644)
}

func (KV) Get(key string) ([]byte, error) {
	return ioutil.ReadFile("kv_" + key)
}

func main() {
	defer zap.L().Sync()
	zap.L().Info("gRPC plugin starting...")
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"kv": &shared.KVGRPCPlugin{Impl: &KV{}},
		},
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     log.NewHCLogAdapter(nil),
	})
}
