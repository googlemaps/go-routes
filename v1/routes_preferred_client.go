// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package routes

import (
	"context"
	"math"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	routespb "google.golang.org/genproto/googleapis/maps/routes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newRoutesPreferredClientHook clientHook

// RoutesPreferredCallOptions contains the retry settings for each method of RoutesPreferredClient.
type RoutesPreferredCallOptions struct {
	ComputeRoutes      []gax.CallOption
	ComputeRouteMatrix []gax.CallOption
}

func defaultRoutesPreferredGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("routespreferred.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("routespreferred.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://routespreferred.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultRoutesPreferredCallOptions() *RoutesPreferredCallOptions {
	return &RoutesPreferredCallOptions{
		ComputeRoutes:      []gax.CallOption{},
		ComputeRouteMatrix: []gax.CallOption{},
	}
}

// internalRoutesPreferredClient is an interface that defines the methods availaible from Routes Preferred API.
type internalRoutesPreferredClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ComputeRoutes(context.Context, *routespb.ComputeRoutesRequest, ...gax.CallOption) (*routespb.ComputeRoutesResponse, error)
	ComputeRouteMatrix(context.Context, *routespb.ComputeRouteMatrixRequest, ...gax.CallOption) (routespb.RoutesPreferred_ComputeRouteMatrixClient, error)
}

// RoutesPreferredClient is a client for interacting with Routes Preferred API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Routes Preferred API.
type RoutesPreferredClient struct {
	// The internal transport-dependent client.
	internalClient internalRoutesPreferredClient

	// The call options for this service.
	CallOptions *RoutesPreferredCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RoutesPreferredClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RoutesPreferredClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *RoutesPreferredClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ComputeRoutes returns the primary route along with optional alternate routes, given a set
// of terminal and intermediate waypoints.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using URL parameter
// $fields or fields, or by using an HTTP/gRPC header X-Goog-FieldMask
// (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters). The value
// is a comma separated list of field paths. See detailed documentation about
// how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//   Field mask of all available fields (for manual inspection):
//   X-Goog-FieldMask: *
//
//   Field mask of Route-level duration, distance, and polyline (an example
//   production setup):
//   X-Goog-FieldMask: routes.duration,routes.distanceMeters,routes.polyline.encodedPolyline
//
// Google discourage the use of the wildcard (*) response field mask, or
// specifying the field mask at the top level (routes), because:
//
//   Selecting only the fields that you need helps our server save computation
//   cycles, allowing us to return the result to you with a lower latency.
//
//   Selecting only the fields that you need
//   in your production job ensures stable latency performance. We might add
//   more response fields in the future, and those new fields might require
//   extra computation time. If you select all fields, or if you select all
//   fields at the top level, then you might experience performance degradation
//   because any new field we add will be automatically included in the
//   response.
//
//   Selecting only the fields that you need results in a smaller response
//   size, and thus higher network throughput.
func (c *RoutesPreferredClient) ComputeRoutes(ctx context.Context, req *routespb.ComputeRoutesRequest, opts ...gax.CallOption) (*routespb.ComputeRoutesResponse, error) {
	return c.internalClient.ComputeRoutes(ctx, req, opts...)
}

// ComputeRouteMatrix takes in a list of origins and destinations and returns a stream containing
// route information for each combination of origin and destination.
//
// NOTE: This method requires that you specify a response field mask in
// the input. You can provide the response field mask by using the URL
// parameter $fields or fields, or by using the HTTP/gRPC header
// X-Goog-FieldMask (see the available URL parameters and
// headers (at https://cloud.google.com/apis/docs/system-parameters). The value
// is a comma separated list of field paths. See this detailed documentation
// about how to construct the field
// paths (at https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto).
//
// For example, in this method:
//
//   Field mask of all available fields (for manual inspection):
//   X-Goog-FieldMask: *
//
//   Field mask of route durations, distances, element status, condition, and
//   element indices (an example production setup):
//   X-Goog-FieldMask: originIndex,destinationIndex,status,condition,distanceMeters,duration
//
// It is critical that you include status in your field mask as otherwise
// all messages will appear to be OK. Google discourages the use of the
// wildcard (*) response field mask, because:
//
//   Selecting only the fields that you need helps our server save computation
//   cycles, allowing us to return the result to you with a lower latency.
//
//   Selecting only the fields that you need in your production job ensures
//   stable latency performance. We might add more response fields in the
//   future, and those new fields might require extra computation time. If you
//   select all fields, or if you select all fields at the top level, then you
//   might experience performance degradation because any new field we add will
//   be automatically included in the response.
//
//   Selecting only the fields that you need results in a smaller response
//   size, and thus higher network throughput.
func (c *RoutesPreferredClient) ComputeRouteMatrix(ctx context.Context, req *routespb.ComputeRouteMatrixRequest, opts ...gax.CallOption) (routespb.RoutesPreferred_ComputeRouteMatrixClient, error) {
	return c.internalClient.ComputeRouteMatrix(ctx, req, opts...)
}

// routesPreferredGRPCClient is a client for interacting with Routes Preferred API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type routesPreferredGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing RoutesPreferredClient
	CallOptions **RoutesPreferredCallOptions

	// The gRPC API client.
	routesPreferredClient routespb.RoutesPreferredClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewRoutesPreferredClient creates a new routes preferred client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Routes Preferred API.
func NewRoutesPreferredClient(ctx context.Context, opts ...option.ClientOption) (*RoutesPreferredClient, error) {
	clientOpts := defaultRoutesPreferredGRPCClientOptions()
	if newRoutesPreferredClientHook != nil {
		hookOpts, err := newRoutesPreferredClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := RoutesPreferredClient{CallOptions: defaultRoutesPreferredCallOptions()}

	c := &routesPreferredGRPCClient{
		connPool:              connPool,
		disableDeadlines:      disableDeadlines,
		routesPreferredClient: routespb.NewRoutesPreferredClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *routesPreferredGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *routesPreferredGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *routesPreferredGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *routesPreferredGRPCClient) ComputeRoutes(ctx context.Context, req *routespb.ComputeRoutesRequest, opts ...gax.CallOption) (*routespb.ComputeRoutesResponse, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ComputeRoutes[0:len((*c.CallOptions).ComputeRoutes):len((*c.CallOptions).ComputeRoutes)], opts...)
	var resp *routespb.ComputeRoutesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routesPreferredClient.ComputeRoutes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *routesPreferredGRPCClient) ComputeRouteMatrix(ctx context.Context, req *routespb.ComputeRouteMatrixRequest, opts ...gax.CallOption) (routespb.RoutesPreferred_ComputeRouteMatrixClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp routespb.RoutesPreferred_ComputeRouteMatrixClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.routesPreferredClient.ComputeRouteMatrix(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
