// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 574fb16d86
// Version Date: 2019年 04月 12日 星期五 00:42:59 UTC

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"context"
	"google.golang.org/grpc"

	"github.com/pkg/errors"

	// This Service
	pb "github.com/fanioc/go-poetryminapp/services/mini-app"
	"github.com/fanioc/go-poetryminapp/services/mini-app/minapp-service/svc/client/cli/handlers"
	grpcclient "github.com/fanioc/go-poetryminapp/services/mini-app/minapp-service/svc/client/grpc"
	httpclient "github.com/fanioc/go-poetryminapp/services/mini-app/minapp-service/svc/client/http"
)

var (
	_ = strconv.ParseInt
	_ = strings.Split
	_ = json.Compact
	_ = errors.Wrapf
	_ = pb.RegisterMinAppServer
)

func main() {
	os.Exit(submain())
}

type headerSeries []string

func (h *headerSeries) Set(val string) error {
	const requiredParts int = 2
	parts := strings.SplitN(val, ":", requiredParts)
	if len(parts) != requiredParts {
		return fmt.Errorf("value %q cannot be split in two; must contain at least one ':' character", val)
	}
	parts[1] = strings.TrimSpace(parts[1])
	*h = append(*h, parts...)
	return nil
}

func (h *headerSeries) String() string {
	return fmt.Sprintf("%v", []string(*h))
}

// submain exists to act as the functional main, but will return exit codes to
// the actual main instead of calling os.Exit directly. This is done to allow
// the defered functions to be called, since if os.Exit where called directly
// from this function, none of the defered functions set up by this function
// would be called.
func submain() int {
	var headers headerSeries
	flag.Var(&headers, "header", "Header(s) to be sent in the transport (follows cURL style)")
	var (
		httpAddr = flag.String("http.addr", "", "HTTP address of addsvc")
		grpcAddr = flag.String("grpc.addr", ":5040", "gRPC (HTTP) address of addsvc")
	)

	// The addcli presumes no service discovery system, and expects users to
	// provide the direct address of an service. This presumption is reflected in
	// the cli binary and the the client packages: the -transport.addr flags
	// and various client constructors both expect host:port strings.

	fsCheckUserSession := flag.NewFlagSet("checkusersession", flag.ExitOnError)

	fsGetUserConfig := flag.NewFlagSet("getuserconfig", flag.ExitOnError)

	fsGetUserInfo := flag.NewFlagSet("getuserinfo", flag.ExitOnError)

	fsLogin := flag.NewFlagSet("login", flag.ExitOnError)

	fsSetUserConfig := flag.NewFlagSet("setuserconfig", flag.ExitOnError)

	fsUpdateUserInfo := flag.NewFlagSet("updateuserinfo", flag.ExitOnError)

	var (
		flagSessionCheckUserSession = fsCheckUserSession.String("session", "", "")
		flagSessionGetUserInfo      = fsGetUserInfo.String("session", "", "")
		flagUserIdGetUserInfo       = fsGetUserInfo.Int("userid", 0, "")
		flagSessionUpdateUserInfo   = fsUpdateUserInfo.String("session", "", "")
		flagSessionGetUserConfig    = fsGetUserConfig.String("session", "", "")
		flagUserIdGetUserConfig     = fsGetUserConfig.Int("userid", 0, "")
		flagSessionSetUserConfig    = fsSetUserConfig.String("session", "", "")
		flagCodeLogin               = fsLogin.String("code", "", "")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Subcommands:\n")
		fmt.Fprintf(os.Stderr, "  %s\n", "checkusersession")
		fmt.Fprintf(os.Stderr, "  %s\n", "getuserconfig")
		fmt.Fprintf(os.Stderr, "  %s\n", "getuserinfo")
		fmt.Fprintf(os.Stderr, "  %s\n", "login")
		fmt.Fprintf(os.Stderr, "  %s\n", "setuserconfig")
		fmt.Fprintf(os.Stderr, "  %s\n", "updateuserinfo")
	}
	if len(os.Args) < 2 {
		flag.Usage()
		return 1
	}

	flag.Parse()

	var (
		service pb.MinAppServer
		err     error
	)

	if *httpAddr != "" {
		service, err = httpclient.New(*httpAddr, httpclient.CtxValuesToSend(headers...))
	} else if *grpcAddr != "" {
		conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while dialing grpc connection: %v", err)
			return 1
		}
		defer conn.Close()
		service, err = grpcclient.New(conn, grpcclient.CtxValuesToSend(headers...))
	} else {
		fmt.Fprintf(os.Stderr, "error: no remote address specified\n")
		return 1
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("No 'method' subcommand provided; exiting.")
		flag.Usage()
		return 1
	}

	ctx := context.Background()
	for i := 0; i < len(headers); i += 2 {
		ctx = context.WithValue(ctx, headers[i], headers[i+1])
	}

	switch flag.Args()[0] {

	case "checkusersession":
		fsCheckUserSession.Parse(flag.Args()[1:])

		SessionCheckUserSession := *flagSessionCheckUserSession

		request, err := handlers.CheckUserSession(SessionCheckUserSession)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.CheckUserSession: %v\n", err)
			return 1
		}

		v, err := service.CheckUserSession(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.CheckUserSession: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(SessionCheckUserSession)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "getuserconfig":
		fsGetUserConfig.Parse(flag.Args()[1:])

		SessionGetUserConfig := *flagSessionGetUserConfig
		UserIdGetUserConfig := int32(*flagUserIdGetUserConfig)

		request, err := handlers.GetUserConfig(SessionGetUserConfig, UserIdGetUserConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.GetUserConfig: %v\n", err)
			return 1
		}

		v, err := service.GetUserConfig(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.GetUserConfig: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(SessionGetUserConfig, UserIdGetUserConfig)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "getuserinfo":
		fsGetUserInfo.Parse(flag.Args()[1:])

		SessionGetUserInfo := *flagSessionGetUserInfo
		UserIdGetUserInfo := int32(*flagUserIdGetUserInfo)

		request, err := handlers.GetUserInfo(SessionGetUserInfo, UserIdGetUserInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.GetUserInfo: %v\n", err)
			return 1
		}

		v, err := service.GetUserInfo(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.GetUserInfo: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(SessionGetUserInfo, UserIdGetUserInfo)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "login":
		fsLogin.Parse(flag.Args()[1:])

		CodeLogin := *flagCodeLogin

		request, err := handlers.Login(CodeLogin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.Login: %v\n", err)
			return 1
		}

		v, err := service.Login(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.Login: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(CodeLogin)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "setuserconfig":
		fsSetUserConfig.Parse(flag.Args()[1:])

		SessionSetUserConfig := *flagSessionSetUserConfig

		request, err := handlers.SetUserConfig(SessionSetUserConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.SetUserConfig: %v\n", err)
			return 1
		}

		v, err := service.SetUserConfig(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.SetUserConfig: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(SessionSetUserConfig)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	case "updateuserinfo":
		fsUpdateUserInfo.Parse(flag.Args()[1:])

		SessionUpdateUserInfo := *flagSessionUpdateUserInfo

		request, err := handlers.UpdateUserInfo(SessionUpdateUserInfo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling handlers.UpdateUserInfo: %v\n", err)
			return 1
		}

		v, err := service.UpdateUserInfo(ctx, request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error calling service.UpdateUserInfo: %v\n", err)
			return 1
		}
		fmt.Println("Client Requested with:")
		fmt.Println(SessionUpdateUserInfo)
		fmt.Println("Server Responded with:")
		fmt.Println(v)

	default:
		flag.Usage()
		return 1
	}

	return 0
}
