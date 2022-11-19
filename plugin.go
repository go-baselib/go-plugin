package plugin

import (
	"github.com/hashicorp/go-plugin"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASELIB_PLUGIN",
	MagicCookieValue: "go-baselib",
}

var (
	pluginSet     = make(plugin.PluginSet)
	typeURLName   = make(map[string]string)
	typeURLPlugin = make(plugin.PluginSet)
)

func Set() plugin.PluginSet {
	return pluginSet
}

func Register(name string, p plugin.Plugin, tu string) {
	pluginSet[name] = p
	typeURLName[tu] = name
	typeURLPlugin[tu] = p
}

func GetPlugin(typeUrl string) (plugin.Plugin, bool) {
	var p, ok = typeURLPlugin[typeUrl]
	return p, ok
}

func GetName(typeUrl string) string {
	return typeURLName[typeUrl]
}
