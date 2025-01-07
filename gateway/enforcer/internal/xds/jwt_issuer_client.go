package xds

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	jwt_issuer_ads "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/service/subscription"
	"github.com/wso2/apk/gateway/enforcer/internal/config"
	"github.com/wso2/apk/gateway/enforcer/internal/logging"
	"github.com/wso2/apk/gateway/enforcer/internal/util"
	"google.golang.org/grpc"
	v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	status "google.golang.org/genproto/googleapis/rpc/status"
	subscription "github.com/wso2/apk/adapter/pkg/discovery/api/wso2/discovery/subscription"
	"google.golang.org/protobuf/proto"
	"github.com/wso2/apk/gateway/enforcer/internal/datastore"
)

const (
	jwtIssuerTypedURL = "type.googleapis.com/wso2.discovery.subscription.JWTIssuerList"
)

// JWTIssuerXDSClient is a client for managing gRPC connections to the API Discovery Service (XDS) 
// for JWT issuer-related configurations. It includes retry logic, TLS configuration, and logging.
type JWTIssuerXDSClient struct {
	Host          string
	Port          string
	maxRetries    int
	retryInterval time.Duration
	tlsConfig     *tls.Config
	grpcConn      *grpc.ClientConn
	ctx           context.Context
	cancel        context.CancelFunc
	client        jwt_issuer_ads.JWTIssuerDiscoveryServiceClient
	log           logging.Logger
	cfg 		   *config.Server
	latestReceived *v3.DiscoveryResponse
	latestACKed    *v3.DiscoveryResponse
	stream 	     jwt_issuer_ads.JWTIssuerDiscoveryService_StreamJWTIssuersClient 
	jwtIssuersDatastore *datastore.JWTIssuerStore
}

// NewJWTIssuerXDSClient creates a new instance of JWTIssuerXDSClient.
// It initializes the client with the given host, port, retry parameters, TLS configuration, and logger.
func NewJWTIssuerXDSClient(host string, port string, maxRetries int, retryInterval time.Duration, tlsConfig *tls.Config, cfg *config.Server, jwtIssuersDatastore *datastore.JWTIssuerStore) *JWTIssuerXDSClient {
	// Create a new APIClient object
	return &JWTIssuerXDSClient{
		Host:          host,
		Port:          port,
		maxRetries:    maxRetries,
		retryInterval: retryInterval,
		tlsConfig:     tlsConfig,
		grpcConn:      nil,
		log:           cfg.Logger,
		cfg:		   cfg,
		jwtIssuersDatastore: jwtIssuersDatastore,
	}
}

// InitiateSubscriptionXDSConnection establishes and maintains a gRPC connection to the API Discovery Service.
// It also handles reconnection logic on errors and listens for incoming JWT issuer configuration streams.
func (c *JWTIssuerXDSClient) InitiateSubscriptionXDSConnection() {
	grpcConn := util.CreateGRPCConnectionWithRetryAndPanic(nil, c.Host, c.Port, c.tlsConfig, c.maxRetries, c.retryInterval)
	c.grpcConn = grpcConn
	client := jwt_issuer_ads.NewJWTIssuerDiscoveryServiceClient(grpcConn)

	ctx, cancel := context.WithCancel(context.Background())
	c.ctx = ctx
	c.cancel = cancel

	stream, err := client.StreamJWTIssuers(ctx)
	if err != nil {
		cancel()
		c.grpcConn.Close()
		panic(fmt.Errorf("Failed to initiate XDS connection with API Discovery Service: %v", err))
	}

	c.stream = stream
	// Send initial request
	dreq := DiscoveryRequestForNode(CreateNode(commonEnforcerLabel, c.cfg.InstanceIdentifier), "", "", nil, jwtIssuerTypedURL)
	if err := stream.Send(dreq); err != nil {
		cancel()
		c.grpcConn.Close()
		panic(fmt.Errorf("failed to send initial discovery request: %v", err))
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				c.log.Error(err, "Failed to receive jwt issuer")
				c.nack(err)
				cancel()
				c.grpcConn.Close()
				go c.InitiateSubscriptionXDSConnection()
				break
			}
			c.log.Info(fmt.Sprintf("Received jwtossier resp: %v", resp))
			c.latestReceived = resp
			handleRespErr := c.handleResponse(resp)
			if handleRespErr != nil {
				c.nack(handleRespErr)
				continue
			}
			c.ack()
		}
	}()
}



func (c *JWTIssuerXDSClient) ack() {
	dreq := DiscoveryRequestForNode(CreateNode(commonEnforcerLabel, c.cfg.InstanceIdentifier), c.latestReceived.GetVersionInfo(), c.latestReceived.GetNonce(), nil, jwtIssuerTypedURL)
	c.stream.Send(dreq)
	c.latestACKed = c.latestReceived
}

func (c *JWTIssuerXDSClient) nack(e error) {
	errDetail := &status.Status{
		Message: e.Error(),
	}
	dreq := DiscoveryRequestForNode(CreateNode(commonEnforcerLabel, c.cfg.InstanceIdentifier), c.latestACKed.GetVersionInfo(), c.latestReceived.GetNonce(), errDetail, jwtIssuerTypedURL)
	c.stream.Send(dreq)
	c.latestACKed = c.latestReceived
}


func (c *JWTIssuerXDSClient) handleResponse(response *v3.DiscoveryResponse) error {

	var jwtIssuerLists []*subscription.JWTIssuerList
	for _, res := range response.GetResources() {
		var jwtIssuerResource subscription.JWTIssuerList
		if err := proto.Unmarshal(res.GetValue(), &jwtIssuerResource); err != nil {
			c.log.Info(fmt.Sprintf("tFailed to unmarshal jwt issuers resource: %v", err))
			return err
		}
		jwtIssuerLists = append(jwtIssuerLists, &jwtIssuerResource)
	}
	var jwtIssuers []*subscription.JWTIssuer
	for _, jwtIssuerList := range jwtIssuerLists {
		jwtIssuers = append(jwtIssuers, jwtIssuerList.GetList()...)
	}
	c.jwtIssuersDatastore.AddJWTIssuers(jwtIssuers)
	c.log.Info(fmt.Sprintf("Number of jwt issuers received: %d", len(jwtIssuerLists)))
	return nil
}
