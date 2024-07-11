package queryserver

import (
	"context"

	"github.com/allora-network/allora-chain/x/emissions/types"
)

func (qs queryServer) GetLatestInfererScore(
	ctx context.Context,
	req *types.QueryLatestInfererScoreRequest,
) (
	*types.QueryLatestInfererScoreResponse,
	error,
) {
	latestInfererScore, err := qs.k.GetLatestInfererScore(ctx, req.TopicId, req.Worker)
	if err != nil {
		return nil, err
	}

	return &types.QueryLatestInfererScoreResponse{Score: &latestInfererScore}, nil
}

func (qs queryServer) GetLatestForecasterScore(
	ctx context.Context,
	req *types.QueryLatestForecasterScoreRequest,
) (
	*types.QueryLatestForecasterScoreResponse,
	error,
) {
	latestForecasterScore, err := qs.k.GetLatestForecasterScore(ctx, req.TopicId, req.Forecaster)
	if err != nil {
		return nil, err
	}

	return &types.QueryLatestForecasterScoreResponse{Score: &latestForecasterScore}, nil
}

func (qs queryServer) GetLatestReputerScore(
	ctx context.Context,
	req *types.QueryLatestReputerScoreRequest,
) (
	*types.QueryLatestReputerScoreResponse,
	error,
) {
	latestReputerScore, err := qs.k.GetLatestReputerScore(ctx, req.TopicId, req.Reputer)
	if err != nil {
		return nil, err
	}

	return &types.QueryLatestReputerScoreResponse{Score: &latestReputerScore}, nil
}
