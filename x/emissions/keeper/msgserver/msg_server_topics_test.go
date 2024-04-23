package msgserver_test

import (
	cosmosMath "cosmossdk.io/math"
	alloraMath "github.com/allora-network/allora-chain/math"
	"github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

/// Topics tests

func (s *KeeperTestSuite) TestMsgCreateNewTopic() {
	ctx, msgServer := s.ctx, s.msgServer
	require := s.Require()

	senderAddr := sdk.AccAddress(PKS[0].Address())
	sender := senderAddr.String()

	// Create a MsgCreateNewTopic message
	newTopicMsg := &types.MsgCreateNewTopic{
		Creator:          sender,
		Metadata:         "Some metadata for the new topic",
		LossLogic:        "logic",
		EpochLength:      10800,
		InferenceLogic:   "Ilogic",
		InferenceMethod:  "Imethod",
		DefaultArg:       "ETH",
		AlphaRegret:      alloraMath.NewDecFromInt64(10),
		PrewardReputer:   alloraMath.NewDecFromInt64(11),
		PrewardInference: alloraMath.NewDecFromInt64(12),
		PrewardForecast:  alloraMath.NewDecFromInt64(13),
		FTolerance:       alloraMath.NewDecFromInt64(14),
	}

	s.MintTokensToAddress(senderAddr, types.DefaultParamsCreateTopicFee())

	// s.PrepareForCreateTopic(newTopicMsg.Creator)
	result, err := msgServer.CreateNewTopic(ctx, newTopicMsg)
	require.NoError(err, "CreateTopic fails on first creation")
	s.Require().NotNil(result)

	pagination := &types.SimpleCursorPaginationRequest{
		Limit: 100,
	}
	activeTopics, _, err := s.emissionsKeeper.GetActiveTopics(s.ctx, pagination)
	require.NoError(err, "CreateTopic fails on first creation")
	found := false
	for _, topicId := range activeTopics {
		if topicId == result.TopicId {
			found = true
			break
		}
	}
	require.True(found, "Added topic not found in active topics")
}

func (s *KeeperTestSuite) TestMsgCreateNewTopicInvalidUnauthorized() {
	ctx, msgServer := s.ctx, s.msgServer
	require := s.Require()

	// Create a MsgCreateNewTopic message
	newTopicMsg := &types.MsgCreateNewTopic{
		Creator:          nonAdminAccounts[0].String(),
		Metadata:         "Some metadata for the new topic",
		LossLogic:        "logic",
		EpochLength:      10800,
		InferenceLogic:   "Ilogic",
		InferenceMethod:  "Imethod",
		DefaultArg:       "ETH",
		AlphaRegret:      alloraMath.NewDecFromInt64(10),
		PrewardReputer:   alloraMath.NewDecFromInt64(11),
		PrewardInference: alloraMath.NewDecFromInt64(12),
		PrewardForecast:  alloraMath.NewDecFromInt64(13),
		FTolerance:       alloraMath.NewDecFromInt64(14),
	}

	_, err := msgServer.CreateNewTopic(ctx, newTopicMsg)
	require.ErrorIs(err, types.ErrNotInTopicCreationWhitelist, "CreateTopic should return an error")
}

func (s *KeeperTestSuite) TestMsgActivateTopic() {
	ctx, msgServer := s.ctx, s.msgServer
	require := s.Require()

	topicCreator := sdk.AccAddress(PKS[0].Address()).String()
	topicId := s.CreateOneTopic()

	// Deactivate topic
	s.emissionsKeeper.InactivateTopic(ctx, topicId)

	// Set unmet demand for topic
	s.emissionsKeeper.SetTopicUnmetDemand(ctx, topicId, cosmosMath.NewUint(100))

	// Create a MsgCreateNewTopic message
	activateTopicMsg := &types.MsgActivateTopic{
		Sender:  topicCreator,
		TopicId: topicId,
	}

	_, err := msgServer.ActivateTopic(ctx, activateTopicMsg)
	require.NoError(err, "ActivateTopic should not return an error")

	// Check if topic is active
	isActive, err := s.emissionsKeeper.IsTopicActive(ctx, topicId)
	require.NoError(err)
	require.True(isActive, "Topic should be active")
}

func (s *KeeperTestSuite) TestMsgActivateTopicInvalidNotEnoughDemand() {
	ctx, msgServer := s.ctx, s.msgServer
	require := s.Require()

	topicCreator := sdk.AccAddress(PKS[0].Address()).String()
	s.CreateOneTopic()

	// Deactivate topic
	s.emissionsKeeper.InactivateTopic(ctx, 0)

	// Create a MsgCreateNewTopic message
	activateTopicMsg := &types.MsgActivateTopic{
		Sender:  topicCreator,
		TopicId: 0,
	}

	_, err := msgServer.ActivateTopic(ctx, activateTopicMsg)
	require.ErrorIs(err, types.ErrTopicNotEnoughDemand, "ctivateTopic should return an error")
}

func (s *KeeperTestSuite) TestUpdateTopicLossUpdateLastRan() {
	ctx := s.ctx
	require := s.Require()
	topicId := s.CreateOneTopic()

	// Mock setup for topic
	inferenceTs := int64(0x0)

	err := s.emissionsKeeper.UpdateTopicEpochLastEnded(ctx, topicId, inferenceTs)
	require.NoError(err, "UpdateTopicEpochLastEnded should not return an error")

	result, err := s.emissionsKeeper.GetTopicEpochLastEnded(s.ctx, topicId)
	s.Require().NoError(err)
	s.Require().NotNil(result)
	s.Require().Equal(result, inferenceTs)
}