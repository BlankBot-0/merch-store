package merch_store

import "context"

func (s *Service) BuyItem(ctx context.Context, request *app.BuyItemRequest) (*app.InfoResponse, error) {
	res, err := s.Deps.Repository.BuyItem(ctx, request.GetUserId(), request.GetItemId())
	return &res, err
}
