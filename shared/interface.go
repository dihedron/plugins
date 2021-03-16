package shared

import "github.com/hashicorp/go-plugin"

// Handshake is used to do a basic handshake between a plugin and a host;
// if the handshake fails, a user friendly error is shown. This prevents
// users from executing bad plugins or executing a plugin directory: it is
// a UX feature, not a security feature.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"dns": &dns.HandlerPlugin{},
}

/*
// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"kv_grpc": &KVGRPCPlugin{},
	"kv":      &KVPlugin{},
}
*/
