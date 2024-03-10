//go:build !grpc
// +build !grpc

package lib

type Service struct{}

func (s *Service) Action() int {
	return 1
}
