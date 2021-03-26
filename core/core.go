package core

import (
	"context"
	"net/http"

	// grpcweb "github.com/improbable-eng/grpc-web/go/grpcweb"
	connmgr "github.com/libp2p/go-libp2p-core/connmgr"
	ma "github.com/multiformats/go-multiaddr"
	logging "github.com/ipfs/go-log/v2"
	threadsClient "github.com/textileio/go-threads/api/client"
	threadsNetclient "github.com/textileio/go-threads/net/api/client"
	tc "github.com/textileio/go-threads/common"
	"google.golang.org/grpc"
)

var (
	log = logging.Logger("core")
)

type Doru struct {
	threadsNetBootStrapper tc.NetBoostrapper

	threadsClient *threadsClient.Client
	threadsNet    *threadsNetclient.Client

	server *grpc.Server
	proxy  *http.Server

	config Config
}

type Config struct {
	Debug bool

	AddressApi         ma.Multiaddr
	AddressThreadsHost ma.Multiaddr
	AddressIpfsApiHost ma.Multiaddr

	ThreadsConnectionManager connmgr.ConnManager
}

func NewDoru(
	ctx context.Context,
	config Config,
	opts ...Option,
) (*Doru, error) {
	var args Options
	for _, opt := range opts {
		opt(&args)
	}

	return nil, nil
}
