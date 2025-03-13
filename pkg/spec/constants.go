package spec

const (
	GnosisGenesis				 = 1638993340
	GnosisBeaconContractAddress  = "0x0B98057eA310F4d31F2a452B414647007d1645d9"
	DepositEventTopic            = "0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5"
	DepositEventDataLength       = 576
)

var BeaconContractAddresses = map[string]string{
	"gnosis": GnosisBeaconContractAddress,
}

/*
Phase0
*/

const (
	MaxEffectiveInc             = 32
	BaseRewardFactor            = 25
	BaseRewardPerEpoch          = 4
	EffectiveBalanceInc         = 1000000000
	SlotsPerEpoch               = 16
	ProposerRewardQuotient      = 8
	SlotsPerHistoricalRoot      = 8192
	SlotSeconds                 = 5
	EpochSlots                  = 16
	WhistleBlowerRewardQuotient = 512
	MinInclusionDelay           = 1

	AttSourceFlagIndex = 0
	AttTargetFlagIndex = 1
	AttHeadFlagIndex   = 2
)

/*
Altair
*/
const (
	// spec weight constants
	TimelySourceWeight = 14
	TimelyTargetWeight = 26
	TimelyHeadWeight   = 14

	SyncRewardWeight  = 2
	ProposerWeight    = 8
	WeightDenominator = 64
	SyncCommitteeSize = 512
)

var (
	ParticipatingFlagsWeight = [3]int{TimelySourceWeight, TimelyTargetWeight, TimelyHeadWeight}
)

type ModelType int8

const (
	BlockModel ModelType = iota
	BlockDropModel
	OrphanModel
	EpochModel
	EpochDropModel
	PoolSummaryModel
	ProposerDutyModel
	ProposerDutyDropModel
	ValidatorLastStatusModel
	ValidatorRewardsModel
	ValidatorRewardDropModel
	WithdrawalModel
	WithdrawalDropModel
	TransactionsModel
	TransactionDropModel
	ReorgModel
	FinalizedCheckpointModel
	HeadEventModel
	ValidatorRewardsAggregationModel
	SlashingModel
	BLSToExecutionChangeModel
	DepositModel
	ETH1DepositModel
)

type ValidatorStatus int8

const (
	QUEUE_STATUS ValidatorStatus = iota
	ACTIVE_STATUS
	EXIT_STATUS
	SLASHED_STATUS
	NUMBER_OF_STATUS // Add new status before this
)
