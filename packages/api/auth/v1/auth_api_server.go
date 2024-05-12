package authv1

import "context"

var (
	_ AuthAPIServer = &authAPIServer{}
)

type authAPIServer struct{}

func NewAuthAPIServer() *authAPIServer {
	return &authAPIServer{}
}

// AuthKakao implements AuthAPIServer.
func (*authAPIServer) AuthKakao(context.Context, *AuthKakaoRequest) (*AuthKakaoResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedAuthAPIServer implements AuthAPIServer.
func (*authAPIServer) mustEmbedUnimplementedAuthAPIServer() {
	panic("unimplemented")
}
