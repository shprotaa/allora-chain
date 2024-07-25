package types

import "cosmossdk.io/collections"

const (
	ModuleName                                 = "emissions"
	StoreKey                                   = "emissions"
	AlloraStakingAccountName                   = "allorastaking"
	AlloraRewardsAccountName                   = "allorarewards"
	AlloraPendingRewardForDelegatorAccountName = "allorapendingrewards"
)

var (
	ParamsKey                                         = collections.NewPrefix(0)
	TotalStakeKey                                     = collections.NewPrefix(1)
	TopicStakeKey                                     = collections.NewPrefix(2)
	RewardsKey                                        = collections.NewPrefix(3)
	NextTopicIdKey                                    = collections.NewPrefix(4)
	TopicsKey                                         = collections.NewPrefix(5)
	TopicWorkersKey                                   = collections.NewPrefix(6)
	TopicReputersKey                                  = collections.NewPrefix(7)
	DelegatorStakeKey                                 = collections.NewPrefix(8)
	DelegateStakePlacementKey                         = collections.NewPrefix(9)
	TargetStakeKey                                    = collections.NewPrefix(10)
	InferencesKey                                     = collections.NewPrefix(11)
	ForecastsKey                                      = collections.NewPrefix(12)
	WorkerNodesKey                                    = collections.NewPrefix(13)
	ReputerNodesKey                                   = collections.NewPrefix(14)
	LatestInferencesTsKey                             = collections.NewPrefix(15)
	ActiveTopicsKey                                   = collections.NewPrefix(16)
	AllInferencesKey                                  = collections.NewPrefix(17)
	AllForecastsKey                                   = collections.NewPrefix(18)
	AllLossBundlesKey                                 = collections.NewPrefix(19)
	StakeRemovalKey                                   = collections.NewPrefix(20)
	StakeByReputerAndTopicId                          = collections.NewPrefix(21)
	DelegateStakeRemovalKey                           = collections.NewPrefix(22)
	AllTopicStakeSumKey                               = collections.NewPrefix(23)
	AddressTopicsKey                                  = collections.NewPrefix(24)
	WhitelistAdminsKey                                = collections.NewPrefix(24)
	ChurnableTopicsKey                                = collections.NewPrefix(25)
	RewardableTopicsKey                               = collections.NewPrefix(26)
	NetworkLossBundlesKey                             = collections.NewPrefix(27)
	NetworkRegretsKey                                 = collections.NewPrefix(28)
	StakeByReputerAndTopicIdKey                       = collections.NewPrefix(29)
	ReputerScoresKey                                  = collections.NewPrefix(30)
	InferenceScoresKey                                = collections.NewPrefix(31)
	ForecastScoresKey                                 = collections.NewPrefix(32)
	ReputerListeningCoefficientKey                    = collections.NewPrefix(33)
	InfererNetworkRegretsKey                          = collections.NewPrefix(34)
	ForecasterNetworkRegretsKey                       = collections.NewPrefix(35)
	OneInForecasterNetworkRegretsKey                  = collections.NewPrefix(36)
	UnfulfilledWorkerNoncesKey                        = collections.NewPrefix(37)
	UnfulfilledReputerNoncesKey                       = collections.NewPrefix(38)
	FeeRevenueEpochKey                                = collections.NewPrefix(39)
	TopicFeeRevenueKey                                = collections.NewPrefix(40)
	PreviousTopicWeightKey                            = collections.NewPrefix(41)
	PreviousReputerRewardFractionKey                  = collections.NewPrefix(42)
	PreviousInferenceRewardFractionKey                = collections.NewPrefix(43)
	PreviousForecastRewardFractionKey                 = collections.NewPrefix(44)
	LatestInfererScoresByWorkerKey                    = collections.NewPrefix(45)
	LatestForecasterScoresByWorkerKey                 = collections.NewPrefix(46)
	LatestReputerScoresByReputerKey                   = collections.NewPrefix(47)
	TopicRewardNonceKey                               = collections.NewPrefix(48)
	DelegateRewardPerShare                            = collections.NewPrefix(49)
	PreviousPercentageRewardToStakedReputersKey       = collections.NewPrefix(50)
	StakeRemovalsByBlockKey                           = collections.NewPrefix(51)
	DelegateStakeRemovalsByBlockKey                   = collections.NewPrefix(52)
	StakeRemovalsByActorKey                           = collections.NewPrefix(53)
	DelegateStakeRemovalsByActorKey                   = collections.NewPrefix(54)
	TopicLastWorkerCommitKey                          = collections.NewPrefix(55)
	TopicLastReputerCommitKey                         = collections.NewPrefix(56)
	TopicLastWorkerPayloadKey                         = collections.NewPrefix(57)
	TopicLastReputerPayloadKey                        = collections.NewPrefix(58)
	LatestNaiveInfererNetworkRegretsKey               = collections.NewPrefix(59)
	LatestOneOutInfererInfererNetworkRegretsKey       = collections.NewPrefix(60)
	LatestOneOutInfererForecasterNetworkRegretsKey    = collections.NewPrefix(61)
	LatestOneOutForecasterInfererNetworkRegretsKey    = collections.NewPrefix(62)
	LatestOneOutForecasterForecasterNetworkRegretsKey = collections.NewPrefix(63)
)
