package msgserver_test

import (
	"math"
	"time"

	cosmosMath "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	state "github.com/allora-network/allora-chain/x/emissions"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (s *KeeperTestSuite) TestRequestInferenceSimple() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, senderAddr, state.AlloraRequestsModuleName, initialStakeCoins)
	response, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().NoError(err, "RequestInference should not return an error")
	s.Require().Equal(&state.MsgRequestInferenceResponse{}, response, "RequestInference should return an empty response on success")

	// Check updated stake for delegator
	r0 := state.CreateNewInferenceRequestFromListItem(r.Sender, r.Requests[0])
	requestId, err := r0.GetRequestId()
	s.Require().NoError(err)
	storedRequest, err := s.emissionsKeeper.GetMempoolInferenceRequestById(s.ctx, 0, requestId)
	s.Require().NoError(err)
	// the last checked time is not set in the request, so we can't compare it
	// we can compare the rest of the fields
	s.Require().Equal(r0.Sender, storedRequest.Sender, "Stored request sender should match the request")
	s.Require().Equal(r0.Nonce, storedRequest.Nonce, "Stored request nonce should match the request")
	s.Require().Equal(r0.TopicId, storedRequest.TopicId, "Stored request topic id should match the request")
	s.Require().Equal(r0.Cadence, storedRequest.Cadence, "Stored request cadence should match the request")
	s.Require().Equal(r0.MaxPricePerInference, storedRequest.MaxPricePerInference, "Stored request max price per inference should match the request")
	s.Require().Equal(r0.BidAmount, storedRequest.BidAmount, "Stored request bid amount should match the request")
	s.Require().GreaterOrEqual(storedRequest.LastChecked, timeNow, "LastChecked should be greater than timeNow")
	s.Require().Equal(r0.TimestampValidUntil, storedRequest.TimestampValidUntil, "Stored request timestamp valid until should match the request")
	s.Require().Equal(r0.ExtraData, storedRequest.ExtraData, "Stored request extra data should match the request")
}

// test more than one inference in the message
func (s *KeeperTestSuite) TestRequestInferenceBatchSimple() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	var requestStake int64 = 500
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	requestStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(requestStake)))
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(requestStake)),
				BidAmount:            cosmosMath.NewUint(uint64(requestStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
			{
				Nonce:                1,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(requestStake)),
				BidAmount:            cosmosMath.NewUint(uint64(requestStake)),
				TimestampValidUntil:  timeNow + 400,
				ExtraData:            nil,
			},
		},
	}
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, senderAddr, state.AlloraRequestsModuleName, requestStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, senderAddr, state.AlloraRequestsModuleName, requestStakeCoins)
	response, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().NoError(err, "RequestInference should not return an error")
	s.Require().Equal(&state.MsgRequestInferenceResponse{}, response, "RequestInference should return an empty response on success")

	// Check updated stake for delegator
	r0 := state.CreateNewInferenceRequestFromListItem(r.Sender, r.Requests[0])
	requestId, err := r0.GetRequestId()
	s.Require().NoError(err)
	storedRequest, err := s.emissionsKeeper.GetMempoolInferenceRequestById(s.ctx, 0, requestId)
	s.Require().NoError(err)
	s.Require().Equal(r0.Sender, storedRequest.Sender, "Stored request sender should match the request")
	s.Require().Equal(r0.Nonce, storedRequest.Nonce, "Stored request nonce should match the request")
	s.Require().Equal(r0.TopicId, storedRequest.TopicId, "Stored request topic id should match the request")
	s.Require().Equal(r0.Cadence, storedRequest.Cadence, "Stored request cadence should match the request")
	s.Require().Equal(r0.MaxPricePerInference, storedRequest.MaxPricePerInference, "Stored request max price per inference should match the request")
	s.Require().Equal(r0.BidAmount, storedRequest.BidAmount, "Stored request bid amount should match the request")
	s.Require().GreaterOrEqual(storedRequest.LastChecked, timeNow, "LastChecked should be greater than timeNow")
	s.Require().Equal(r0.TimestampValidUntil, storedRequest.TimestampValidUntil, "Stored request timestamp valid until should match the request")
	s.Require().Equal(r0.ExtraData, storedRequest.ExtraData, "Stored request extra data should match the request")
	r1 := state.CreateNewInferenceRequestFromListItem(r.Sender, r.Requests[1])
	requestId, err = r1.GetRequestId()
	s.Require().NoError(err)
	storedRequest, err = s.emissionsKeeper.GetMempoolInferenceRequestById(s.ctx, 0, requestId)
	s.Require().NoError(err)
	s.Require().Equal(r1.Sender, storedRequest.Sender, "Stored request sender should match the request")
	s.Require().Equal(r1.Nonce, storedRequest.Nonce, "Stored request nonce should match the request")
	s.Require().Equal(r1.TopicId, storedRequest.TopicId, "Stored request topic id should match the request")
	s.Require().Equal(r1.Cadence, storedRequest.Cadence, "Stored request cadence should match the request")
	s.Require().Equal(r1.MaxPricePerInference, storedRequest.MaxPricePerInference, "Stored request max price per inference should match the request")
	s.Require().Equal(r1.BidAmount, storedRequest.BidAmount, "Stored request bid amount should match the request")
	s.Require().GreaterOrEqual(storedRequest.LastChecked, timeNow, "LastChecked should be greater than timeNow")
	s.Require().Equal(r1.TimestampValidUntil, storedRequest.TimestampValidUntil, "Stored request timestamp valid until should match the request")
	s.Require().Equal(r1.ExtraData, storedRequest.ExtraData, "Stored request extra data should match the request")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidTopicDoesNotExist() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address()).String()
	r := state.MsgRequestInference{
		Sender: senderAddr,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(100),
				BidAmount:            cosmosMath.NewUint(100),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInvalidTopicId, "RequestInference should return an error when the topic does not exist")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidBidAmountNotEnoughForPriceSet() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake + 20)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestBidAmountLessThanPrice, "RequestInference should return an error when the bid amount is less than the price")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidSendSameRequestTwice() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(s.ctx, senderAddr, state.AlloraRequestsModuleName, initialStakeCoins)
	s.msgServer.RequestInference(s.ctx, &r)

	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestAlreadyInMempool, "RequestInference should return an error when the request already exists")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidRequestInThePast() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow - 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestTimestampValidUntilInPast, "RequestInference should return an error when the request timestamp is in the past")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidRequestTooFarInFuture() {
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  math.MaxUint64,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestTimestampValidUntilTooFarInFuture, "RequestInference should return an error when the request timestamp is too far in the future")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidRequestCadenceHappensAfterNoLongerValid() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              1000,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestWillNeverBeScheduled, "RequestInference should return an error when the request cadence happens after the request is no longer valid")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidRequestCadenceTooFast() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              1,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestCadenceTooFast, "RequestInference should return an error when the request cadence is too fast")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidRequestCadenceTooSlow() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              math.MaxUint64,
				MaxPricePerInference: cosmosMath.NewUint(uint64(initialStake)),
				BidAmount:            cosmosMath.NewUint(uint64(initialStake)),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestCadenceTooSlow, "RequestInference should return an error when the request cadence is too slow")
}

func (s *KeeperTestSuite) TestRequestInferenceInvalidBidAmountLessThanGlobalMinimum() {
	timeNow := uint64(time.Now().UTC().Unix())
	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()
	s.CreateOneTopic()
	var initialStake int64 = 1000
	initialStakeCoins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, cosmosMath.NewInt(initialStake)))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.MintCoins(s.ctx, state.AlloraStakingModuleName, initialStakeCoins)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, state.AlloraStakingModuleName, senderAddr, initialStakeCoins)
	r := state.MsgRequestInference{
		Sender: sender,
		Requests: []*state.RequestInferenceListItem{
			{
				Nonce:                0,
				TopicId:              0,
				Cadence:              0,
				MaxPricePerInference: cosmosMath.ZeroUint(),
				BidAmount:            cosmosMath.ZeroUint(),
				TimestampValidUntil:  timeNow + 100,
				ExtraData:            []byte("Test"),
			},
		},
	}
	_, err := s.msgServer.RequestInference(s.ctx, &r)
	s.Require().ErrorIs(err, state.ErrInferenceRequestBidAmountTooLow, "RequestInference should return an error when the bid amount is below global minimum threshold")
}