package server

import (
	"flag"
	"fmt"
	rc "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/common"
	"golang.conradwood.net/go-easyops/rpc"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"sync"
)

var (
	rpcclient  rc.RPCInterceptorServiceClient
	gettingrpc = false
	rpclock    sync.Mutex
	debug_auth = flag.Bool("ge_debug_auth", false, "debug grpc authentication")
)

func initrpc() error {
	if gettingrpc {
		return fmt.Errorf("[go-easyops] (auth) RPCInterceptor unavailable")
	}
	if rpcclient != nil {
		return nil
	}
	rpclock.Lock()
	defer rpclock.Unlock()
	gettingrpc = true
	if rpcclient != nil {
		return nil
	}
	if rpcclient == nil {
		rpcclient = rc.NewRPCInterceptorServiceClient(client.Connect("rpcinterceptor.RPCInterceptorService"))
	}
	gettingrpc = false
	return nil
}

// authenticate a user (and authorise access to this method/service)
func Authenticate(cs *rpc.CallState) error {
	if *debug_auth {
		cs.Debug = true
	}
	err := initrpc()
	if err != nil {
		return err
	}
	if cs.Debug {
		fmt.Printf("[go-easyops] Calling RPC Interceptor...\n")
	}

	cs.Metadata = MetaFromContext(cs.Context)
	if cs.Debug {
		fmt.Printf("[go-easyops] Inbound metadata: %#v\n", cs.Metadata)
	}

	// call the interceptor
	irr := &rc.InterceptRPCRequest{
		InMetadata: cs.Metadata,
		Service:    cs.ServiceName,
		Method:     cs.MethodName,
	}

	// preserve some of the inbound metadata information (before we overwrite it withour outbound data)
	if cs.Metadata != nil {
		// is user signature valid??
		if !common.VerifySignature(cs.Metadata.User) {
			// it is not - clear/remove field immediately
			if *debug_auth {
				fmt.Printf("[go-easyops] WARNING - user signature invalid and ignored\n")
			}
			cs.Metadata.User = nil
		}
		if !common.VerifySignature(cs.Metadata.Service) {
			// it is not - clear/remove field immediately
			if *debug_auth {
				fmt.Printf("[go-easyops] WARNING - user signature invalid and ignored\n")
			}
			cs.Metadata.Service = nil
		}
		cs.CallingMethodID = cs.Metadata.CallerMethodID
		if cs.Metadata.UserToken == "" && cs.Metadata.ServiceToken == "" && cs.Metadata.User == nil && cs.Metadata.Service == nil {
			t, ok := peer.FromContext(cs.Context)
			if ok && t != nil && t.Addr != nil {
				irr.Source = t.Addr.String()
			}
			fmt.Printf("[go-easyops] no identification by caller whatsoever (from %s)\n", irr.Source)
		}
	} else {
	}

	res, err := rpcclient.InterceptRPC(cs.Context, irr)
	if err != nil {
		if *debug_auth {
			fmt.Printf("[go-easyops] RPCInterceptor.InterceptRPC() failed: %s\n", utils.ErrorString(err))
		}
		return err
	}
	cs.RPCIResponse = res

	// copy /some/ responses to inmeta
	cs.Metadata.RequestID = cs.RPCIResponse.RequestID
	cs.Metadata.CallerMethodID = cs.RPCIResponse.CallerMethodID
	if cs.RPCIResponse.CallerUser != nil {
		cs.Metadata.UserID = cs.RPCIResponse.CallerUser.ID
		cs.Metadata.User = cs.RPCIResponse.CallerUser
	}
	cs.Metadata.CallerServiceID = cs.MyServiceID
	// all subsequent rpcs propagate OUR servicetoken
	cs.Metadata.ServiceToken = tokens.GetServiceTokenParameter()
	cs.Metadata.FooBar = "authmoo"
	if *debug_auth {
		fmt.Printf("[go-easyops] metadata after rpc interceptor: %#v\n", cs.Metadata)
		fmt.Printf("[go-easyops] RPC Interceptor (reject=%t) said: %v\n", cs.RPCIResponse.Reject, cs.RPCIResponse)
	}
	return nil
}

func MetaFromContext(ctx context.Context) *rc.InMetadata {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("[go-easyops] Warning - cannot extract metadata from context (peer=%s)\n", peerFromContext(ctx))
		return nil
	}
	ims := headers[tokens.METANAME]
	if ims == nil || len(ims) == 0 {
		fmt.Printf("[go-easyops] Warning - metadata in context is nil or 0 (peer=%s)\n", peerFromContext(ctx))
		return nil
	}
	res := &rc.InMetadata{}
	err := utils.Unmarshal(ims[0], res)
	if err != nil {
		fmt.Printf("[go-easyops] Warning - unable to unmarshal metadata (%s)\n", err)
		return nil
	}
	return res
}

func peerFromContext(ctx context.Context) string {
	s := ""
	t, ok := peer.FromContext(ctx)
	if ok && t != nil && t.Addr != nil {
		s = t.Addr.String()
	}
	return s
}
