package merch_store

import "context"

func (s *Service) Info(ctx context.Context, request *app.InfoRequest) (*app.InfoResponse, error) {
	res, err := s.Deps.Repository.Info(ctx, request.GetUserId())
	return &res, err
}
