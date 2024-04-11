// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: archivematica/ccp/admin/v1/admin.proto

package adminv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/artefactual/archivematica/hack/ccp/internal/api/gen/archivematica/ccp/admin/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// AdminServiceName is the fully-qualified name of the AdminService service.
	AdminServiceName = "archivematica.ccp.admin.v1.AdminService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AdminServiceListActivePackagesProcedure is the fully-qualified name of the AdminService's
	// ListActivePackages RPC.
	AdminServiceListActivePackagesProcedure = "/archivematica.ccp.admin.v1.AdminService/ListActivePackages"
	// AdminServiceListAwaitingDecisionsProcedure is the fully-qualified name of the AdminService's
	// ListAwaitingDecisions RPC.
	AdminServiceListAwaitingDecisionsProcedure = "/archivematica.ccp.admin.v1.AdminService/ListAwaitingDecisions"
	// AdminServiceResolveAwaitingDecisionProcedure is the fully-qualified name of the AdminService's
	// ResolveAwaitingDecision RPC.
	AdminServiceResolveAwaitingDecisionProcedure = "/archivematica.ccp.admin.v1.AdminService/ResolveAwaitingDecision"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	adminServiceServiceDescriptor                       = v1.File_archivematica_ccp_admin_v1_admin_proto.Services().ByName("AdminService")
	adminServiceListActivePackagesMethodDescriptor      = adminServiceServiceDescriptor.Methods().ByName("ListActivePackages")
	adminServiceListAwaitingDecisionsMethodDescriptor   = adminServiceServiceDescriptor.Methods().ByName("ListAwaitingDecisions")
	adminServiceResolveAwaitingDecisionMethodDescriptor = adminServiceServiceDescriptor.Methods().ByName("ResolveAwaitingDecision")
)

// AdminServiceClient is a client for the archivematica.ccp.admin.v1.AdminService service.
type AdminServiceClient interface {
	ListActivePackages(context.Context, *connect.Request[v1.ListActivePackagesRequest]) (*connect.Response[v1.ListActivePackagesResponse], error)
	ListAwaitingDecisions(context.Context, *connect.Request[v1.ListAwaitingDecisionsRequest]) (*connect.Response[v1.ListAwaitingDecisionsResponse], error)
	ResolveAwaitingDecision(context.Context, *connect.Request[v1.ResolveAwaitingDecisionRequest]) (*connect.Response[v1.ResolveAwaitingDecisionResponse], error)
}

// NewAdminServiceClient constructs a client for the archivematica.ccp.admin.v1.AdminService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAdminServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AdminServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &adminServiceClient{
		listActivePackages: connect.NewClient[v1.ListActivePackagesRequest, v1.ListActivePackagesResponse](
			httpClient,
			baseURL+AdminServiceListActivePackagesProcedure,
			connect.WithSchema(adminServiceListActivePackagesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listAwaitingDecisions: connect.NewClient[v1.ListAwaitingDecisionsRequest, v1.ListAwaitingDecisionsResponse](
			httpClient,
			baseURL+AdminServiceListAwaitingDecisionsProcedure,
			connect.WithSchema(adminServiceListAwaitingDecisionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resolveAwaitingDecision: connect.NewClient[v1.ResolveAwaitingDecisionRequest, v1.ResolveAwaitingDecisionResponse](
			httpClient,
			baseURL+AdminServiceResolveAwaitingDecisionProcedure,
			connect.WithSchema(adminServiceResolveAwaitingDecisionMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// adminServiceClient implements AdminServiceClient.
type adminServiceClient struct {
	listActivePackages      *connect.Client[v1.ListActivePackagesRequest, v1.ListActivePackagesResponse]
	listAwaitingDecisions   *connect.Client[v1.ListAwaitingDecisionsRequest, v1.ListAwaitingDecisionsResponse]
	resolveAwaitingDecision *connect.Client[v1.ResolveAwaitingDecisionRequest, v1.ResolveAwaitingDecisionResponse]
}

// ListActivePackages calls archivematica.ccp.admin.v1.AdminService.ListActivePackages.
func (c *adminServiceClient) ListActivePackages(ctx context.Context, req *connect.Request[v1.ListActivePackagesRequest]) (*connect.Response[v1.ListActivePackagesResponse], error) {
	return c.listActivePackages.CallUnary(ctx, req)
}

// ListAwaitingDecisions calls archivematica.ccp.admin.v1.AdminService.ListAwaitingDecisions.
func (c *adminServiceClient) ListAwaitingDecisions(ctx context.Context, req *connect.Request[v1.ListAwaitingDecisionsRequest]) (*connect.Response[v1.ListAwaitingDecisionsResponse], error) {
	return c.listAwaitingDecisions.CallUnary(ctx, req)
}

// ResolveAwaitingDecision calls archivematica.ccp.admin.v1.AdminService.ResolveAwaitingDecision.
func (c *adminServiceClient) ResolveAwaitingDecision(ctx context.Context, req *connect.Request[v1.ResolveAwaitingDecisionRequest]) (*connect.Response[v1.ResolveAwaitingDecisionResponse], error) {
	return c.resolveAwaitingDecision.CallUnary(ctx, req)
}

// AdminServiceHandler is an implementation of the archivematica.ccp.admin.v1.AdminService service.
type AdminServiceHandler interface {
	ListActivePackages(context.Context, *connect.Request[v1.ListActivePackagesRequest]) (*connect.Response[v1.ListActivePackagesResponse], error)
	ListAwaitingDecisions(context.Context, *connect.Request[v1.ListAwaitingDecisionsRequest]) (*connect.Response[v1.ListAwaitingDecisionsResponse], error)
	ResolveAwaitingDecision(context.Context, *connect.Request[v1.ResolveAwaitingDecisionRequest]) (*connect.Response[v1.ResolveAwaitingDecisionResponse], error)
}

// NewAdminServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAdminServiceHandler(svc AdminServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	adminServiceListActivePackagesHandler := connect.NewUnaryHandler(
		AdminServiceListActivePackagesProcedure,
		svc.ListActivePackages,
		connect.WithSchema(adminServiceListActivePackagesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	adminServiceListAwaitingDecisionsHandler := connect.NewUnaryHandler(
		AdminServiceListAwaitingDecisionsProcedure,
		svc.ListAwaitingDecisions,
		connect.WithSchema(adminServiceListAwaitingDecisionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	adminServiceResolveAwaitingDecisionHandler := connect.NewUnaryHandler(
		AdminServiceResolveAwaitingDecisionProcedure,
		svc.ResolveAwaitingDecision,
		connect.WithSchema(adminServiceResolveAwaitingDecisionMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/archivematica.ccp.admin.v1.AdminService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AdminServiceListActivePackagesProcedure:
			adminServiceListActivePackagesHandler.ServeHTTP(w, r)
		case AdminServiceListAwaitingDecisionsProcedure:
			adminServiceListAwaitingDecisionsHandler.ServeHTTP(w, r)
		case AdminServiceResolveAwaitingDecisionProcedure:
			adminServiceResolveAwaitingDecisionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAdminServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAdminServiceHandler struct{}

func (UnimplementedAdminServiceHandler) ListActivePackages(context.Context, *connect.Request[v1.ListActivePackagesRequest]) (*connect.Response[v1.ListActivePackagesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("archivematica.ccp.admin.v1.AdminService.ListActivePackages is not implemented"))
}

func (UnimplementedAdminServiceHandler) ListAwaitingDecisions(context.Context, *connect.Request[v1.ListAwaitingDecisionsRequest]) (*connect.Response[v1.ListAwaitingDecisionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("archivematica.ccp.admin.v1.AdminService.ListAwaitingDecisions is not implemented"))
}

func (UnimplementedAdminServiceHandler) ResolveAwaitingDecision(context.Context, *connect.Request[v1.ResolveAwaitingDecisionRequest]) (*connect.Response[v1.ResolveAwaitingDecisionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("archivematica.ccp.admin.v1.AdminService.ResolveAwaitingDecision is not implemented"))
}