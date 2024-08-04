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
	TopicToNextPossibleChurningBlockKey               = collections.NewPrefix(16)
	BlockToLowestActiveTopicWeightKey                 = collections.NewPrefix(17)
	BlockToActiveTopicsKey                            = collections.NewPrefix(18)
	AllInferencesKey                                  = collections.NewPrefix(19)
	AllForecastsKey                                   = collections.NewPrefix(20)
	AllLossBundlesKey                                 = collections.NewPrefix(21)
	StakeRemovalKey                                   = collections.NewPrefix(22)
	StakeByReputerAndTopicId                          = collections.NewPrefix(23)
	DelegateStakeRemovalKey                           = collections.NewPrefix(24)
	AllTopicStakeSumKey                               = collections.NewPrefix(25)
	AddressTopicsKey                                  = collections.NewPrefix(26)
	WhitelistAdminsKey                                = collections.NewPrefix(27)
	ChurnableTopicsKey                                = collections.NewPrefix(28)
	RewardableTopicsKey                               = collections.NewPrefix(29)
	NetworkLossBundlesKey                             = collections.NewPrefix(30)
	NetworkRegretsKey                                 = collections.NewPrefix(31)
	StakeByReputerAndTopicIdKey                       = collections.NewPrefix(32)
	ReputerScoresKey                                  = collections.NewPrefix(33)
	InferenceScoresKey                                = collections.NewPrefix(34)
	ForecastScoresKey                                 = collections.NewPrefix(35)
	ReputerListeningCoefficientKey                    = collections.NewPrefix(36)
	InfererNetworkRegretsKey                          = collections.NewPrefix(37)
	ForecasterNetworkRegretsKey                       = collections.NewPrefix(38)
	OneInForecasterNetworkRegretsKey                  = collections.NewPrefix(39)
	UnfulfilledWorkerNoncesKey                        = collections.NewPrefix(40)
	UnfulfilledReputerNoncesKey                       = collections.NewPrefix(41)
	FeeRevenueEpochKey                                = collections.NewPrefix(42)
	TopicFeeRevenueKey                                = collections.NewPrefix(43)
	PreviousTopicWeightKey                            = collections.NewPrefix(44)
	PreviousReputerRewardFractionKey                  = collections.NewPrefix(45)
	PreviousInferenceRewardFractionKey                = collections.NewPrefix(46)
	PreviousForecastRewardFractionKey                 = collections.NewPrefix(47)
	LatestInfererScoresByWorkerKey                    = collections.NewPrefix(48)
	LatestForecasterScoresByWorkerKey                 = collections.NewPrefix(49)
	LatestReputerScoresByReputerKey                   = collections.NewPrefix(50)
	TopicRewardNonceKey                               = collections.NewPrefix(51)
	DelegateRewardPerShare                            = collections.NewPrefix(52)
	PreviousPercentageRewardToStakedReputersKey       = collections.NewPrefix(53)
	StakeRemovalsByBlockKey                           = collections.NewPrefix(54)
	DelegateStakeRemovalsByBlockKey                   = collections.NewPrefix(55)
	StakeRemovalsByActorKey                           = collections.NewPrefix(56)
	DelegateStakeRemovalsByActorKey                   = collections.NewPrefix(57)
	TopicLastWorkerCommitKey                          = collections.NewPrefix(58)
	TopicLastReputerCommitKey                         = collections.NewPrefix(59)
	TopicLastWorkerPayloadKey                         = collections.NewPrefix(60)
	TopicLastReputerPayloadKey                        = collections.NewPrefix(61)
	OpenWorkerWindowsKey                              = collections.NewPrefix(62)
	LatestNaiveInfererNetworkRegretsKey               = collections.NewPrefix(63)
	LatestOneOutInfererInfererNetworkRegretsKey       = collections.NewPrefix(64)
	LatestOneOutInfererForecasterNetworkRegretsKey    = collections.NewPrefix(65)
	LatestOneOutForecasterInfererNetworkRegretsKey    = collections.NewPrefix(66)
	LatestOneOutForecasterForecasterNetworkRegretsKey = collections.NewPrefix(67)
	PreviousForecasterScoreRatioKey                   = collections.NewPrefix(68)
)
