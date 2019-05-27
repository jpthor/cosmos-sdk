// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/querier
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/staking/types
package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/querier"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	DefaultParamspace                  = keeper.DefaultParamspace
	QueryValidators                    = querier.QueryValidators
	QueryValidator                     = querier.QueryValidator
	QueryDelegatorDelegations          = querier.QueryDelegatorDelegations
	QueryDelegatorUnbondingDelegations = querier.QueryDelegatorUnbondingDelegations
	QueryRedelegations                 = querier.QueryRedelegations
	QueryValidatorDelegations          = querier.QueryValidatorDelegations
	QueryValidatorRedelegations        = querier.QueryValidatorRedelegations
	QueryValidatorUnbondingDelegations = querier.QueryValidatorUnbondingDelegations
	QueryDelegation                    = querier.QueryDelegation
	QueryUnbondingDelegation           = querier.QueryUnbondingDelegation
	QueryDelegatorValidators           = querier.QueryDelegatorValidators
	QueryDelegatorValidator            = querier.QueryDelegatorValidator
	QueryPool                          = querier.QueryPool
	QueryParameters                    = querier.QueryParameters
	DefaultCodespace                   = types.DefaultCodespace
	CodeInvalidValidator               = types.CodeInvalidValidator
	CodeInvalidDelegation              = types.CodeInvalidDelegation
	CodeInvalidInput                   = types.CodeInvalidInput
	CodeValidatorJailed                = types.CodeValidatorJailed
	CodeInvalidAddress                 = types.CodeInvalidAddress
	CodeUnauthorized                   = types.CodeUnauthorized
	CodeInternal                       = types.CodeInternal
	CodeUnknownRequest                 = types.CodeUnknownRequest
	ModuleName                         = types.ModuleName
	StoreKey                           = types.StoreKey
	TStoreKey                          = types.TStoreKey
	QuerierRoute                       = types.QuerierRoute
	RouterKey                          = types.RouterKey
	DefaultUnbondingTime               = types.DefaultUnbondingTime
	DefaultMaxValidators               = types.DefaultMaxValidators
	DefaultMaxEntries                  = types.DefaultMaxEntries
	MaxMonikerLength                   = types.MaxMonikerLength
	MaxIdentityLength                  = types.MaxIdentityLength
	MaxWebsiteLength                   = types.MaxWebsiteLength
	MaxDetailsLength                   = types.MaxDetailsLength
	DoNotModifyDesc                    = types.DoNotModifyDesc
	Bonded                             = types.Bonded
	Unbonding                          = types.Unbonding
	Unbonded                           = types.Unbonded
)

var (
	// functions aliases
	RegisterInvariants                 = keeper.RegisterInvariants
	AllInvariants                      = keeper.AllInvariants
	SupplyInvariants                   = keeper.SupplyInvariants
	NonNegativePowerInvariant          = keeper.NonNegativePowerInvariant
	PositiveDelegationInvariant        = keeper.PositiveDelegationInvariant
	DelegatorSharesInvariant           = keeper.DelegatorSharesInvariant
	NewKeeper                          = keeper.NewKeeper
	GetValidatorKey                    = keeper.GetValidatorKey
	GetValidatorByConsAddrKey          = keeper.GetValidatorByConsAddrKey
	AddressFromLastValidatorPowerKey   = keeper.AddressFromLastValidatorPowerKey
	GetValidatorsByPowerIndexKey       = keeper.GetValidatorsByPowerIndexKey
	GetLastValidatorPowerKey           = keeper.GetLastValidatorPowerKey
	GetValidatorQueueTimeKey           = keeper.GetValidatorQueueTimeKey
	GetDelegationKey                   = keeper.GetDelegationKey
	GetDelegationsKey                  = keeper.GetDelegationsKey
	GetUBDKey                          = keeper.GetUBDKey
	GetUBDByValIndexKey                = keeper.GetUBDByValIndexKey
	GetUBDKeyFromValIndexKey           = keeper.GetUBDKeyFromValIndexKey
	GetUBDsKey                         = keeper.GetUBDsKey
	GetUBDsByValIndexKey               = keeper.GetUBDsByValIndexKey
	GetUnbondingDelegationTimeKey      = keeper.GetUnbondingDelegationTimeKey
	GetREDKey                          = keeper.GetREDKey
	GetREDByValSrcIndexKey             = keeper.GetREDByValSrcIndexKey
	GetREDByValDstIndexKey             = keeper.GetREDByValDstIndexKey
	GetREDKeyFromValSrcIndexKey        = keeper.GetREDKeyFromValSrcIndexKey
	GetREDKeyFromValDstIndexKey        = keeper.GetREDKeyFromValDstIndexKey
	GetRedelegationTimeKey             = keeper.GetRedelegationTimeKey
	GetREDsKey                         = keeper.GetREDsKey
	GetREDsFromValSrcIndexKey          = keeper.GetREDsFromValSrcIndexKey
	GetREDsToValDstIndexKey            = keeper.GetREDsToValDstIndexKey
	GetREDsByDelToValDstIndexKey       = keeper.GetREDsByDelToValDstIndexKey
	ParamKeyTable                      = keeper.ParamKeyTable
	ValEq                              = keeper.ValEq
	MakeTestCodec                      = keeper.MakeTestCodec
	CreateTestInput                    = keeper.CreateTestInput
	NewPubKey                          = keeper.NewPubKey
	TestAddr                           = keeper.TestAddr
	ValidatorByPowerIndexExists        = keeper.ValidatorByPowerIndexExists
	TestingUpdateValidator             = keeper.TestingUpdateValidator
	RandomValidator                    = keeper.RandomValidator
	RandomBondedValidator              = keeper.RandomBondedValidator
	NewQuerier                         = querier.NewQuerier
	NewQueryDelegatorParams            = querier.NewQueryDelegatorParams
	NewQueryValidatorParams            = querier.NewQueryValidatorParams
	NewQueryBondsParams                = querier.NewQueryBondsParams
	NewQueryRedelegationParams         = querier.NewQueryRedelegationParams
	NewQueryValidatorsParams           = querier.NewQueryValidatorsParams
	RegisterCodec                      = types.RegisterCodec
	NewCommissionMsg                   = types.NewCommissionMsg
	NewCommission                      = types.NewCommission
	NewCommissionWithTime              = types.NewCommissionWithTime
	NewDelegation                      = types.NewDelegation
	MustMarshalDelegation              = types.MustMarshalDelegation
	MustUnmarshalDelegation            = types.MustUnmarshalDelegation
	UnmarshalDelegation                = types.UnmarshalDelegation
	NewUnbondingDelegation             = types.NewUnbondingDelegation
	NewUnbondingDelegationEntry        = types.NewUnbondingDelegationEntry
	MustMarshalUBD                     = types.MustMarshalUBD
	MustUnmarshalUBD                   = types.MustUnmarshalUBD
	UnmarshalUBD                       = types.UnmarshalUBD
	NewRedelegation                    = types.NewRedelegation
	NewRedelegationEntry               = types.NewRedelegationEntry
	MustMarshalRED                     = types.MustMarshalRED
	MustUnmarshalRED                   = types.MustUnmarshalRED
	UnmarshalRED                       = types.UnmarshalRED
	NewDelegationResp                  = types.NewDelegationResp
	NewRedelegationResponse            = types.NewRedelegationResponse
	NewRedelegationEntryResponse       = types.NewRedelegationEntryResponse
	ErrNilValidatorAddr                = types.ErrNilValidatorAddr
	ErrBadValidatorAddr                = types.ErrBadValidatorAddr
	ErrNoValidatorFound                = types.ErrNoValidatorFound
	ErrValidatorOwnerExists            = types.ErrValidatorOwnerExists
	ErrValidatorPubKeyExists           = types.ErrValidatorPubKeyExists
	ErrValidatorPubKeyTypeNotSupported = types.ErrValidatorPubKeyTypeNotSupported
	ErrValidatorJailed                 = types.ErrValidatorJailed
	ErrBadRemoveValidator              = types.ErrBadRemoveValidator
	ErrDescriptionLength               = types.ErrDescriptionLength
	ErrCommissionNegative              = types.ErrCommissionNegative
	ErrCommissionHuge                  = types.ErrCommissionHuge
	ErrCommissionGTMaxRate             = types.ErrCommissionGTMaxRate
	ErrCommissionUpdateTime            = types.ErrCommissionUpdateTime
	ErrCommissionChangeRateNegative    = types.ErrCommissionChangeRateNegative
	ErrCommissionChangeRateGTMaxRate   = types.ErrCommissionChangeRateGTMaxRate
	ErrCommissionGTMaxChangeRate       = types.ErrCommissionGTMaxChangeRate
	ErrSelfDelegationBelowMinimum      = types.ErrSelfDelegationBelowMinimum
	ErrMinSelfDelegationInvalid        = types.ErrMinSelfDelegationInvalid
	ErrMinSelfDelegationDecreased      = types.ErrMinSelfDelegationDecreased
	ErrNilDelegatorAddr                = types.ErrNilDelegatorAddr
	ErrBadDenom                        = types.ErrBadDenom
	ErrBadDelegationAddr               = types.ErrBadDelegationAddr
	ErrBadDelegationAmount             = types.ErrBadDelegationAmount
	ErrNoDelegation                    = types.ErrNoDelegation
	ErrBadDelegatorAddr                = types.ErrBadDelegatorAddr
	ErrNoDelegatorForAddress           = types.ErrNoDelegatorForAddress
	ErrInsufficientShares              = types.ErrInsufficientShares
	ErrDelegationValidatorEmpty        = types.ErrDelegationValidatorEmpty
	ErrNotEnoughDelegationShares       = types.ErrNotEnoughDelegationShares
	ErrBadSharesAmount                 = types.ErrBadSharesAmount
	ErrBadSharesPercent                = types.ErrBadSharesPercent
	ErrNotMature                       = types.ErrNotMature
	ErrNoUnbondingDelegation           = types.ErrNoUnbondingDelegation
	ErrMaxUnbondingDelegationEntries   = types.ErrMaxUnbondingDelegationEntries
	ErrBadRedelegationAddr             = types.ErrBadRedelegationAddr
	ErrNoRedelegation                  = types.ErrNoRedelegation
	ErrSelfRedelegation                = types.ErrSelfRedelegation
	ErrVerySmallRedelegation           = types.ErrVerySmallRedelegation
	ErrBadRedelegationDst              = types.ErrBadRedelegationDst
	ErrTransitiveRedelegation          = types.ErrTransitiveRedelegation
	ErrMaxRedelegationEntries          = types.ErrMaxRedelegationEntries
	ErrDelegatorShareExRateInvalid     = types.ErrDelegatorShareExRateInvalid
	ErrBothShareMsgsGiven              = types.ErrBothShareMsgsGiven
	ErrNeitherShareMsgsGiven           = types.ErrNeitherShareMsgsGiven
	ErrMissingSignature                = types.ErrMissingSignature
	NewGenesisState                    = types.NewGenesisState
	DefaultGenesisState                = types.DefaultGenesisState
	NewMultiStakingHooks               = types.NewMultiStakingHooks
	NewMsgCreateValidator              = types.NewMsgCreateValidator
	NewMsgEditValidator                = types.NewMsgEditValidator
	NewMsgDelegate                     = types.NewMsgDelegate
	NewMsgBeginRedelegate              = types.NewMsgBeginRedelegate
	NewMsgUndelegate                   = types.NewMsgUndelegate
	NewParams                          = types.NewParams
	DefaultParams                      = types.DefaultParams
	MustUnmarshalParams                = types.MustUnmarshalParams
	UnmarshalParams                    = types.UnmarshalParams
	InitialPool                        = types.InitialPool
	MustUnmarshalPool                  = types.MustUnmarshalPool
	UnmarshalPool                      = types.UnmarshalPool
	NewValidator                       = types.NewValidator
	MustMarshalValidator               = types.MustMarshalValidator
	MustUnmarshalValidator             = types.MustUnmarshalValidator
	UnmarshalValidator                 = types.UnmarshalValidator
	NewDescription                     = types.NewDescription

	// variable aliases
	PoolKey                          = keeper.PoolKey
	LastValidatorPowerKey            = keeper.LastValidatorPowerKey
	LastTotalPowerKey                = keeper.LastTotalPowerKey
	ValidatorsKey                    = keeper.ValidatorsKey
	ValidatorsByConsAddrKey          = keeper.ValidatorsByConsAddrKey
	ValidatorsByPowerIndexKey        = keeper.ValidatorsByPowerIndexKey
	DelegationKey                    = keeper.DelegationKey
	UnbondingDelegationKey           = keeper.UnbondingDelegationKey
	UnbondingDelegationByValIndexKey = keeper.UnbondingDelegationByValIndexKey
	RedelegationKey                  = keeper.RedelegationKey
	RedelegationByValSrcIndexKey     = keeper.RedelegationByValSrcIndexKey
	RedelegationByValDstIndexKey     = keeper.RedelegationByValDstIndexKey
	UnbondingQueueKey                = keeper.UnbondingQueueKey
	RedelegationQueueKey             = keeper.RedelegationQueueKey
	ValidatorQueueKey                = keeper.ValidatorQueueKey
	Addrs                            = keeper.Addrs
	PKs                              = keeper.PKs
	ModuleCdc                        = types.ModuleCdc
	KeyUnbondingTime                 = types.KeyUnbondingTime
	KeyMaxValidators                 = types.KeyMaxValidators
	KeyMaxEntries                    = types.KeyMaxEntries
	KeyBondDenom                     = types.KeyBondDenom
)

type (
	Keeper                    = keeper.Keeper
	QueryDelegatorParams      = querier.QueryDelegatorParams
	QueryValidatorParams      = querier.QueryValidatorParams
	QueryBondsParams          = querier.QueryBondsParams
	QueryRedelegationParams   = querier.QueryRedelegationParams
	QueryValidatorsParams     = querier.QueryValidatorsParams
	Commission                = types.Commission
	CommissionMsg             = types.CommissionMsg
	DVPair                    = types.DVPair
	DVVTriplet                = types.DVVTriplet
	Delegation                = types.Delegation
	DelegationInterface       = types.DelegationInterface
	Delegations               = types.Delegations
	UnbondingDelegation       = types.UnbondingDelegation
	UnbondingDelegationEntry  = types.UnbondingDelegationEntry
	UnbondingDelegations      = types.UnbondingDelegations
	Redelegation              = types.Redelegation
	RedelegationEntry         = types.RedelegationEntry
	Redelegations             = types.Redelegations
	DelegationResponse        = types.DelegationResponse
	DelegationResponses       = types.DelegationResponses
	RedelegationResponse      = types.RedelegationResponse
	RedelegationEntryResponse = types.RedelegationEntryResponse
	RedelegationResponses     = types.RedelegationResponses
	CodeType                  = types.CodeType
	DistributionKeeper        = types.DistributionKeeper
	FeeCollectionKeeper       = types.FeeCollectionKeeper
	BankKeeper                = types.BankKeeper
	AccountKeeper             = types.AccountKeeper
	GenesisState              = types.GenesisState
	LastValidatorPower        = types.LastValidatorPower
	MultiStakingHooks         = types.MultiStakingHooks
	MsgCreateValidator        = types.MsgCreateValidator
	MsgEditValidator          = types.MsgEditValidator
	MsgDelegate               = types.MsgDelegate
	MsgBeginRedelegate        = types.MsgBeginRedelegate
	MsgUndelegate             = types.MsgUndelegate
	Params                    = types.Params
	Pool                      = types.Pool
	Validator                 = types.Validator
	ValidatorInterface        = types.ValidatorInterface
	Validators                = types.Validators
	Description               = types.Description
)
