package authremote

import (
	"context"
	"fmt"
	apb "golang.conradwood.net/apis/auth"
	rc "golang.conradwood.net/apis/rpcinterceptor"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/cache"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/rpc"
	"golang.conradwood.net/go-easyops/tokens"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc/metadata"
	"os"
	"sync"
	"time"
)

var (
	userbyidcache    = cache.NewResolvingCache("userbyid", time.Duration(60)*time.Second, 9999)
	userbyemailcache = cache.NewResolvingCache("userbyemail", time.Duration(60)*time.Second, 9999)
	rpci             rc.RPCInterceptorServiceClient
	authServer       apb.AuthenticationServiceClient
	authServerLock   sync.Mutex
	authManager      apb.AuthManagerServiceClient
	authManagerLock  sync.Mutex
	contextRetrieved = false
	lastUser         *apb.User
	lastService      *apb.User
)

func Context() context.Context {
	client.GetSignatureFromAuth()
	return ContextWithTimeout(time.Duration(10) * time.Second)
}

/* this context gives a context with a full userobject
todo so it _has_ to call external servers to get a signed userobject.
if started_by_autodeployer will use tokens.ContextWithToken()
else if environment variable with context, will use auth.Context() (with variable)
else create context by asking auth service for a signed user object
*/
func ContextWithTimeout(t time.Duration) context.Context {
	if cmdline.Datacenter() {
		return tokens.ContextWithTokenAndTimeout(uint64(t.Seconds()))
	}
	sctx := os.Getenv("GE_CTX")
	if sctx != "" {
		return auth.Context(t)
	}

	if !contextRetrieved {
		lastUser = GetByToken(context.Background(), tokens.GetUserTokenParameter())
		lastService = GetByToken(context.Background(), tokens.GetServiceTokenParameter())
		contextRetrieved = true
	}
	luid := ""
	if lastUser != nil {
		luid = lastUser.ID
	}
	cs := &rpc.CallState{
		Metadata: &rc.InMetadata{
			UserID:  luid,
			Service: lastService,
			User:    lastUser,
		},
		RPCIResponse: &rc.InterceptRPCResponse{
			CallerUser:    lastUser,
			CallerService: lastService,
		},
	}
	err := cs.UpdateContextFromResponseWithTimeout(t)
	if err != nil {
		panic(fmt.Sprintf("bad context: %s", err))
	}
	return cs.Context
}

func GetAuthManagerClient() apb.AuthManagerServiceClient {
	managerClient()
	return authManager
}
func GetAuthClient() apb.AuthenticationServiceClient {
	authClient()
	return authServer
}

// create an outbound context for a given user. user must be valid and signed
// this is an expensive call ! (also calls rpcinterceptor)
// this is not privileged (user must be signed)
func ContextForUser(user *apb.User) (context.Context, error) {
	if user == nil {
		return nil, fmt.Errorf("Missing user")
	}
	if rpci == nil {
		rpci = rc.NewRPCInterceptorServiceClient(client.Connect("rpcinterceptor.RPCInterceptorService"))
	}
	token := tokens.GetServiceTokenParameter()
	mt := &rc.InMetadata{
		FooBar:       "local",
		ServiceToken: token,
		UserID:       user.ID,
		User:         user,
	}
	mts, err := utils.Marshal(mt)
	if err != nil {
		return nil, err
	}
	cs := &rpc.CallState{
		Started:  time.Now(),
		Debug:    true,
		Metadata: mt,
	}
	newmd := metadata.Pairs(tokens.METANAME, mts)
	ctx := tokens.ContextWithToken()
	ctx = context.WithValue(ctx, rpc.LOCALCONTEXTNAME, cs)
	res := metadata.NewOutgoingContext(ctx, newmd)
	cs.Context = ctx
	resp, err := rpci.InterceptRPC(ctx, &rc.InterceptRPCRequest{InMetadata: mt, Service: "local", Method: "local", Source: "go-easyops/auth/context.go"})
	if err != nil {
		return nil, err
	}
	cs.RPCIResponse = resp
	//	cs.PrintContext()
	if resp != nil && mt != nil && mt.User == nil {
		mt.User = resp.CallerUser
	}
	if resp != nil && mt != nil && mt.Service == nil {
		mt.Service = resp.CallerService
	}

	// now rebuild metadata again to add to outbound context
	// this must be the final step
	mts, err = utils.Marshal(mt)
	if err != nil {
		return nil, err
	}
	newmd = metadata.Pairs(tokens.METANAME, mts)
	ctx = tokens.ContextWithToken()
	ctx = context.WithValue(ctx, rpc.LOCALCONTEXTNAME, cs)
	res = metadata.NewOutgoingContext(ctx, newmd)
	cs.Context = ctx

	return res, nil

}

// create an outbound context for a given user by id (with current service token)
// this is an expensive call ! (also calls rpcinterceptor)
// it is also privileged
func ContextForUserID(userid string) (context.Context, error) {
	if userid == "" || userid == "0" {
		return nil, fmt.Errorf("Missing userid")
	}
	if rpci == nil {
		rpci = rc.NewRPCInterceptorServiceClient(client.Connect("rpcinterceptor.RPCInterceptorService"))
	}
	token := tokens.GetServiceTokenParameter()
	if token == "" {
		return nil, fmt.Errorf("no service token parameter to generate contextforuser by id")
	}
	mt := &rc.InMetadata{
		FooBar:       "local",
		ServiceToken: token,
		UserID:       userid,
	}
	mts, err := utils.Marshal(mt)
	if err != nil {
		return nil, err
	}
	cs := &rpc.CallState{
		Started:  time.Now(),
		Debug:    true,
		Metadata: mt,
	}
	newmd := metadata.Pairs(tokens.METANAME, mts)
	ctx := tokens.ContextWithToken()
	ctx = context.WithValue(ctx, rpc.LOCALCONTEXTNAME, cs)
	res := metadata.NewOutgoingContext(ctx, newmd)
	cs.Context = ctx
	resp, err := rpci.InterceptRPC(ctx, &rc.InterceptRPCRequest{InMetadata: mt, Service: "local", Method: "local", Source: "go-easyops/auth/context.go"})
	if err != nil {
		return nil, err
	}
	cs.RPCIResponse = resp
	//	cs.PrintContext()
	if resp != nil && mt != nil && mt.User == nil {
		mt.User = resp.CallerUser
	}
	if resp != nil && mt != nil && mt.Service == nil {
		mt.Service = resp.CallerService
	}

	// now rebuild metadata again to add to outbound context
	// this must be the final step
	mts, err = utils.Marshal(mt)
	if err != nil {
		return nil, err
	}
	newmd = metadata.Pairs(tokens.METANAME, mts)
	ctx = tokens.ContextWithToken()
	ctx = context.WithValue(ctx, rpc.LOCALCONTEXTNAME, cs)
	res = metadata.NewOutgoingContext(ctx, newmd)
	cs.Context = ctx

	return res, nil

}
func GetUserByID(ctx context.Context, userid string) (*apb.User, error) {
	if userid == "" {
		return nil, fmt.Errorf("[go-easyops] No userid provided")
	}
	o, err := userbyidcache.Retrieve(userid, func(k string) (interface{}, error) {
		managerClient()
		res, err := authManager.GetUserByID(ctx, &apb.ByIDRequest{UserID: k})
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return o.(*apb.User), nil
}
func GetUserByEmail(ctx context.Context, email string) (*apb.User, error) {
	if email == "" {
		return nil, fmt.Errorf("[go-easyops] No email provided")
	}
	o, err := userbyemailcache.Retrieve(email, func(k string) (interface{}, error) {
		managerClient()
		res, err := authManager.GetUserByEmail(ctx, &apb.ByEmailRequest{Email: k})
		return res, err
	})
	if err != nil {
		return nil, err
	}
	return o.(*apb.User), nil
}
func WhoAmI() *apb.User {
	tok := tokens.GetUserTokenParameter()
	return GetByToken(context.Background(), tok)
}
func GetByToken(ctx context.Context, token string) *apb.User {
	if token == "" {
		return nil
	}
	authClient()
	ar, err := authServer.GetByToken(ctx, &apb.AuthenticateTokenRequest{Token: token})
	if err != nil {
		return nil
	}
	if !ar.Valid {
		return nil
	}
	if !ar.User.Active {
		return nil
	}
	return ar.User
}

func authClient() {
	if authServer == nil {
		authServerLock.Lock()
		defer authServerLock.Unlock()
		if authServer != nil {
			return
		}
		authServer = apb.NewAuthenticationServiceClient(client.Connect("auth.AuthenticationService"))
	}
}
func managerClient() {
	if authManager == nil {
		authManagerLock.Lock()
		defer authManagerLock.Unlock()
		if authManager != nil {
			return
		}
		authManager = apb.NewAuthManagerServiceClient(client.Connect("auth.AuthManagerService"))
	}
}
