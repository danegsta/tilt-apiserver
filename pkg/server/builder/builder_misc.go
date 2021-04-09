package builder

import (
	"io"

	"github.com/tilt-dev/tilt-apiserver/pkg/server/apiserver"
	"github.com/tilt-dev/tilt-apiserver/pkg/server/start"
	openapicommon "k8s.io/kube-openapi/pkg/common"
)

// WithBindPort registers a default port to serve on.
// This port can still be overidden with flags.
func (a *Server) WithBindPort(port int) *Server {
	a.serving.BindPort = port
	return a
}

// WithConnProvider registers a provider of connections.
//
// If this is set, the server will use this listener and dialer instead of a network host/port.
//
// Mostly useful for testing.
func (a *Server) WithConnProvider(connProvider apiserver.ConnProvider) *Server {
	a.connProvider = connProvider
	return a
}

// WithOpenAPIDefinitions registers resource OpenAPI definitions generated by openapi-gen.
//
//    export K8sAPIS=k8s.io/apimachinery/pkg/api/resource,\
//      k8s.io/apimachinery/pkg/apis/meta/v1,\
//      k8s.io/apimachinery/pkg/runtime,\
//      k8s.io/apimachinery/pkg/version
//    export MY_APIS=my-go-pkg/pkg/apis/my-group/my-version
//    export OPENAPI=my-go-pkg/pkg/generated/openapi
//    openapi-gen --input-dirs $K8SAPIS,$MY_APIS --output-package $OPENAPI \
//      -O zz_generated.openapi --output-base ../../.. --go-header-file ./hack/boilerplate.go.txt
func (a *Server) WithOpenAPIDefinitions(
	name, version string, openAPI openapicommon.GetOpenAPIDefinitions) *Server {
	a.recommendedConfigFns = append(a.recommendedConfigFns, start.SetOpenAPIDefinitionFn(a.scheme, name, version, openAPI))
	return a
}

// WithOutputWriter redirects output from both stdout and stderr to a custom writer.
func (a *Server) WithOutputWriter(out io.Writer) *Server {
	a.stdout = out
	a.stderr = out
	return a
}

// WithBearerToken sets up an auth token that's needed to send server requests.
func (a *Server) WithBearerToken(token string) *Server {
	a.serving.BearerToken = token
	return a
}
