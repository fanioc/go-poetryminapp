// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 574fb16d86
// Version Date: 2019年 04月 12日 星期五 00:42:59 UTC

package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/fanioc/go-poetryminapp/services/mini-app/minapp-service/svc/server"
)

// Config will be populated by ENV vars on initilization
// flags will overwrite ENV vars in Config on flag.Parse()
var Config server.Config

func init() {
	flag.StringVar(&Config.DebugAddr, "debug.addr", ":5060", "Debug and metrics listen address")
	flag.StringVar(&Config.HTTPAddr, "http.addr", ":5050", "HTTP listen address")
	flag.StringVar(&Config.GRPCAddr, "grpc.addr", ":5040", "gRPC (HTTP) listen address")

	// Use environment variables, if set. Flags have priority over Env vars.
	if addr := os.Getenv("DEBUG_ADDR"); addr != "" {
		Config.DebugAddr = addr
	}
	if port := os.Getenv("PORT"); port != "" {
		Config.HTTPAddr = fmt.Sprintf(":%s", port)
	}
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		Config.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		Config.GRPCAddr = addr
	}
}
