// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bitcoin

import (
	"errors"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

var (
	bigOne = big.NewInt(1)

	mainPowLimit   = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
	mainPowLimitV1 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 24), bigOne)
	mainPowLimitV2 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)

	regressionPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)

	testNet3PowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)
)

type Checkpoint struct {
	Height int32
	Hash   *chainhash.Hash
}

type DNSSeed struct {
	Host         string
	HasFiltering bool
}

type ConsensusDeployment struct {
	ActivationHeight    uint32
	hashActivationBlock string
	// BitNumber defines the specific bit number within the block version
	// this particular soft-fork deployment refers to.
	BitNumber uint8

	// MinActivationHeight is an optional field that when set (default
	// value being zero), modifies the traditional BIP 9 state machine by
	// only transitioning from LockedIn to Active once the block height is
	// greater than (or equal to) thus specified height.
	MinActivationHeight uint32

	// CustomActivationThreshold if set (non-zero), will _override_ the
	// existing RuleChangeActivationThreshold value set at the
	// network/chain level. This value divided by the active
	// MinerConfirmationWindow denotes the threshold required for
	// activation. A value of 1815 block denotes a 90% threshold.
	CustomActivationThreshold uint32

	// DeploymentStarter is used to determine if the given
	// ConsensusDeployment has started or not.
	DeploymentStarter ConsensusDeploymentStarter

	// DeploymentEnder is used to determine if the given
	// ConsensusDeployment has ended or not.
	DeploymentEnder ConsensusDeploymentEnder
}

// Constants that define the deployment offset in the deployments field of the
// parameters for each deployment.  This is useful to be able to get the details
// of a specific deployment by name.
const (
	// DeploymentTestDummy defines the rule change deployment ID for testing
	// purposes.
	DeploymentTestDummy = iota

	// DeploymentTestDummyMinActivation defines the rule change deployment
	// ID for testing purposes. This differs from the DeploymentTestDummy
	// in that it specifies the newer params the taproot fork used for
	// activation: a custom threshold and a min activation height.
	DeploymentTestDummyMinActivation

	// DeploymentCSV defines the rule change deployment ID for the CSV
	// soft-fork package. The CSV package includes the deployment of BIPS
	// 68, 112, and 113.
	DeploymentCSV

	// DeploymentSegwit defines the rule change deployment ID for the
	// Segregated Witness (segwit) soft-fork package. The segwit package
	// includes the deployment of BIPS 141, 142, 144, 145, 147 and 173.
	DeploymentSegwit

	//pivx network upgrades
	UpgradeBaseNetwork
	UpgradeTestDummy
	UpgradePOS
	UpgradePOSV2
	UpgradeZC
	UpgradeZCV2
	UpgradeBIP65
	UpgradeZCPublic
	UpgradeV34
	UpgradeV40
	UpgradeV50Dummy

	// NOTE: DefinedDeployments must always come last since it is used to
	// determine how many defined deployments there currently are.

	// DefinedDeployments is the number of currently defined deployments.
	DefinedDeployments
)

// Params defines a Bitcoin network by its parameters.  These parameters may be
// used by Bitcoin applications to differentiate networks as well as addresses
// and keys for one network from those intended for use on another network.
type Params struct {
	Name         string
	Net          wire.BitcoinNet
	DefaultPort  string
	DNSSeeds     []DNSSeed
	GenesisBlock *wire.MsgBlock
	GenesisHash  *chainhash.Hash

	PowAllowMinDifficultyBlocks bool
	PowLimit                    *big.Int
	PowLimitV1                  *big.Int
	PowLimitV2                  *big.Int
	BudgetCycleBlocks           int32
	BudgetFeeConfirmations      int32
	CoinbaseMaturity            uint16
	FutureTimeDriftPoW          int32
	FutureTimeDriftPoS          int32
	PoolMaxTransactions         int32
	ProposalEstablishmentTime   time.Duration
	StakeMinAge                 time.Duration
	StakeMinDepth               int
	TargetTimespan              time.Duration
	TargetTimespanV2            time.Duration
	TargetSpacing               time.Duration
	TimeSlotLength              time.Duration
	strSporkPubKey              string
	strSporkPubKeyOld           string
	Time_EnforceNewSporkKey     uint
	Time_RejectOldSporkKey      uint

	height_last_ZC_AccumCheckpoint   int32
	height_last_ZC_WrappedSerials    int32
	height_start_InvalidUTXOsCheck   int32
	height_start_ZC_InvalidSerials   int32
	height_start_ZC_SerialRangeCheck int32
	height_ZC_RecalcAccumulators     int32

	PivxBadBlockTime int
	PivxBadBlockBits uint

	//checkpoints
	Checkpoints []Checkpoint

	Upgrades [DefinedDeployments]ConsensusDeployment

	// Address encoding magics
	PubKeyHashAddrID byte
	ScriptHashAddrID byte
	PrivateKeyID     byte
	StakingKeyID     byte

	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte

	HDCoinType uint32

	Deployments [DefinedDeployments]ConsensusDeployment

	// These fields define the block heights at which the specified softfork
	// BIP became active.
	BIP0034Height int32
	BIP0065Height int32
	BIP0066Height int32

	// SubsidyReductionInterval is the interval of blocks before the subsidy
	// is reduced.
	SubsidyReductionInterval int32

	// TargetTimePerBlock is the desired amount of time to generate each
	// block.
	TargetTimePerBlock time.Duration

	// RetargetAdjustmentFactor is the adjustment factor used to limit
	// the minimum and maximum amount of adjustment that can occur between
	// difficulty retargets.
	RetargetAdjustmentFactor int64

	// ReduceMinDifficulty defines whether the network should reduce the
	// minimum required difficulty after a long enough period of time has
	// passed without finding a block.  This is really only useful for test
	// networks and should not be set on a main network.
	ReduceMinDifficulty bool

	// MinDiffReductionTime is the amount of time after which the minimum
	// required difficulty should be reduced when a block hasn't been found.
	//
	// NOTE: This only applies if ReduceMinDifficulty is true.
	MinDiffReductionTime time.Duration

	// GenerateSupported specifies whether or not CPU mining is allowed.
	GenerateSupported bool

	// Checkpoints ordered from oldest to newest.

	// These fields are related to voting on consensus rule changes as
	// defined by BIP0009.
	//
	// RuleChangeActivationThreshold is the number of blocks in a threshold
	// state retarget window for which a positive vote for a rule change
	// must be cast in order to lock in a rule change. It should typically
	// be 95% for the main network and 75% for test networks.
	//
	// MinerConfirmationWindow is the number of blocks in each threshold
	// state retarget window.
	//
	// Deployments define the specific consensus rule changes to be voted
	// on.
	RuleChangeActivationThreshold uint32
	MinerConfirmationWindow       uint32

	// Mempool parameters
	RelayNonStdTxs bool

	WitnessPubKeyHashAddrID byte // First byte of a P2WPKH address
	WitnessScriptHashAddrID byte // First byte of a P2WSH address

}

// Constants used to indicate the message dogecoin network.
const (
	// MainNet represents the main dogecoin network.
	MainNet wire.BitcoinNet = 0xe9fdc490

	// TestNet3 represents the test network (version 3).
	TestNet3 wire.BitcoinNet = 0xba657645
)

// MainNetParams defines the network parameters for the main Bitcoin network.
var MainNetParams = Params{
	Name:        "main",
	Net:         MainNet,
	DefaultPort: "46462",
	DNSSeeds: []DNSSeed{
		{"seed.euno.network", true},
	},

	// Chain parameters
	GenesisBlock:                &genesisBlock,
	GenesisHash:                 &genesisHash,
	PowAllowMinDifficultyBlocks: false,
	PowLimit:                    mainPowLimit,
	PowLimitV1:                  mainPowLimitV1,
	PowLimitV2:                  mainPowLimitV2,
	BudgetCycleBlocks:           43200,
	BudgetFeeConfirmations:      6,
	CoinbaseMaturity:            100,
	FutureTimeDriftPoW:          7200,
	FutureTimeDriftPoS:          180,
	PoolMaxTransactions:         3,
	ProposalEstablishmentTime:   time.Hour * 60 * 24,
	StakeMinAge:                 time.Hour * 60,
	StakeMinDepth:               600,
	TargetTimespan:              time.Minute * 40 * 60,
	TargetTimespanV2:            time.Minute * 30 * 60,
	TargetSpacing:               time.Minute * 1 * 60,
	TimeSlotLength:              15,

	//spork keys
	strSporkPubKey:          "04653C23ED96425AE787AE88884CC53F30921D12B6BAA827B737CE845782C7A0FD1F226575EB99520F2ED03E55DC51FC817B640C1A822874B0CBF25924DE87C947",
	strSporkPubKeyOld:       "04653C23ED96425AE787AE88884CC53F30921D12B6BAA827B737CE845782C7A0FD1F226575EB99520F2ED03E55DC51FC817B640C1A822874B0CBF25924DE87C947",
	Time_EnforceNewSporkKey: 1566860400,
	Time_RejectOldSporkKey:  1569538800,

	// height-based activations
	height_last_ZC_AccumCheckpoint:   1441,
	height_last_ZC_WrappedSerials:    1441,
	height_start_InvalidUTXOsCheck:   1441,
	height_start_ZC_InvalidSerials:   1441,
	height_start_ZC_SerialRangeCheck: 1441,
	height_ZC_RecalcAccumulators:     1441,

	// validation by-pass
	PivxBadBlockTime: 1471401614,
	PivxBadBlockBits: 0x1c056dac,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{0, newHashFromStr("0000000069e244f73d78e8fd29ba2fd2ed618bd6fa2ee92559f542fdb26e7c1d")},
	},

	Upgrades: [DefinedDeployments]ConsensusDeployment{
		UpgradeBaseNetwork: {
			ActivationHeight: 0,
		},
		UpgradeTestDummy: {
			ActivationHeight: 0,
		},
		UpgradePOS: {
			ActivationHeight: 1441,
		},
		UpgradePOSV2: {
			ActivationHeight: 1441,
		},
		UpgradeZC: {
			ActivationHeight:    1441,
			hashActivationBlock: "00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f",
		},
		UpgradeZCV2: {
			ActivationHeight:    1441,
			hashActivationBlock: "00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f",
		},
		UpgradeBIP65: {
			ActivationHeight:    1441,
			hashActivationBlock: "00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f",
		},
		UpgradeZCPublic: {
			ActivationHeight: 1441,
		},
		UpgradeV34: {
			ActivationHeight:    1441,
			hashActivationBlock: "00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f",
		},
		UpgradeV40: {
			ActivationHeight:    1441,
			hashActivationBlock: "00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f",
		},
		UpgradeV50Dummy: {
			ActivationHeight: 0,
		},
	},

	// Address encoding magics
	PubKeyHashAddrID: 139, // starts with m or n
	ScriptHashAddrID: 19,  // starts with 2
	PrivateKeyID:     239, // starts with 9 (uncompressed) or c (compressed)
	StakingKeyID:     73,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x3a, 0x80, 0x58, 0x37}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x3a, 0x80, 0x61, 0xa0}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 303,
}

// RegressionNetParams defines the network parameters for the regression test
// Bitcoin network.  Not to be confused with the test Bitcoin network (version
// 3), this network is sometimes simply called "testnet".
var RegressionNetParams = Params{
	Name:        "regtest",
	Net:         wire.TestNet,
	DefaultPort: "18444",
	DNSSeeds:    []DNSSeed{},

	// Chain parameters
	GenesisBlock:             &regTestGenesisBlock,
	GenesisHash:              &regTestGenesisHash,
	PowLimit:                 regressionPowLimit,
	CoinbaseMaturity:         100,
	BIP0034Height:            100000000, // Not active - Permit ver 1 blocks
	BIP0065Height:            1351,      // Used by regression tests
	BIP0066Height:            1251,      // Used by regression tests
	SubsidyReductionInterval: 150,
	TargetTimespan:           time.Hour * 24 * 14, // 14 days
	TargetTimePerBlock:       time.Minute * 10,    // 10 minutes
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 20, // TargetTimePerBlock * 2
	GenerateSupported:        true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 108, // 75%  of MinerConfirmationWindow
	MinerConfirmationWindow:       144,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		UpgradeBaseNetwork: {
			MinActivationHeight: 0,
		},
		DeploymentTestDummyMinActivation: {
			BitNumber:                 22,
			CustomActivationThreshold: 72,  // Only needs 50% hash rate.
			MinActivationHeight:       600, // Can only activate after height 600.
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Time{}, // Always available for vote
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Time{}, // Never expires
			),
		},
		DeploymentCSV: {
			BitNumber: 0,
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Time{}, // Always available for vote
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Time{}, // Never expires
			),
		},
		DeploymentSegwit: {
			BitNumber: 1,
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Time{}, // Always available for vote
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Time{}, // Never expires.
			),
		},
	},

	// Address encoding magics
	PubKeyHashAddrID: 139, // starts with m or n
	ScriptHashAddrID: 19,  // starts with 2
	PrivateKeyID:     239, // starts with 9 (uncompressed) or c (compressed)
	StakingKeyID:     73,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3Params = Params{
	Name:        "testnet3",
	Net:         wire.TestNet3,
	DefaultPort: "18333",
	DNSSeeds: []DNSSeed{
		{"testnet-seed.bitcoin.jonasschnelli.ch", true},
		{"testnet-seed.bitcoin.schildbach.de", false},
		{"seed.tbtc.petertodd.org", true},
		{"testnet-seed.bluematt.me", false},
	},

	// Chain parameters
	GenesisBlock:             &testNet3GenesisBlock,
	GenesisHash:              &testNet3GenesisHash,
	PowLimit:                 testNet3PowLimit,
	BIP0034Height:            21111,  // 0000000023b3a96d3484e5abb3755c413e7d41500f8e2a5c3f0dd01299cd8ef8
	BIP0065Height:            581885, // 00000000007f6655f22f98e72ed80d8b06dc761d5da09df0fa1dc4be4f861eb6
	BIP0066Height:            330776, // 000000002104c8c45e99a8853285a3b592602a3ccde2b832481da85e9e4ba182
	CoinbaseMaturity:         100,
	SubsidyReductionInterval: 210000,
	TargetTimespan:           time.Hour * 24 * 14, // 14 days
	TargetTimePerBlock:       time.Minute * 10,    // 10 minutes
	RetargetAdjustmentFactor: 4,                   // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 20, // TargetTimePerBlock * 2
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{546, newHashFromStr("000000002a936ca763904c3c35fce2f3556c559c0214345d31b1bcebf76acb70")},
		{100000, newHashFromStr("00000000009e2958c15ff9290d571bf9459e93b19765c6801ddeccadbb160a1e")},
		{200000, newHashFromStr("0000000000287bffd321963ef05feab753ebe274e1d78b2fd4e2bfe9ad3aa6f2")},
		{300001, newHashFromStr("0000000000004829474748f3d1bc8fcf893c88be255e6d7f571c548aff57abf4")},
		{400002, newHashFromStr("0000000005e2c73b8ecb82ae2dbc2e8274614ebad7172b53528aba7501f5a089")},
		{500011, newHashFromStr("00000000000929f63977fbac92ff570a9bd9e7715401ee96f2848f7b07750b02")},
		{600002, newHashFromStr("000000000001f471389afd6ee94dcace5ccc44adc18e8bff402443f034b07240")},
		{700000, newHashFromStr("000000000000406178b12a4dea3b27e13b3c4fe4510994fd667d7c1e6a3f4dc1")},
		{800010, newHashFromStr("000000000017ed35296433190b6829db01e657d80631d43f5983fa403bfdb4c1")},
		{900000, newHashFromStr("0000000000356f8d8924556e765b7a94aaebc6b5c8685dcfa2b1ee8b41acd89b")},
		{1000007, newHashFromStr("00000000001ccb893d8a1f25b70ad173ce955e5f50124261bbbc50379a612ddf")},
		{1100007, newHashFromStr("00000000000abc7b2cd18768ab3dee20857326a818d1946ed6796f42d66dd1e8")},
		{1200007, newHashFromStr("00000000000004f2dc41845771909db57e04191714ed8c963f7e56713a7b6cea")},
		{1300007, newHashFromStr("0000000072eab69d54df75107c052b26b0395b44f77578184293bf1bb1dbd9fa")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1512, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       2016,
	Deployments: [DefinedDeployments]ConsensusDeployment{
		DeploymentTestDummy: {
			BitNumber: 28,
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Unix(1199145601, 0), // January 1, 2008 UTC
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Unix(1230767999, 0), // December 31, 2008 UTC
			),
		},
		DeploymentTestDummyMinActivation: {
			BitNumber:                 22,
			CustomActivationThreshold: 1815,    // Only needs 90% hash rate.
			MinActivationHeight:       10_0000, // Can only activate after height 10k.
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Time{}, // Always available for vote
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Time{}, // Never expires
			),
		},
		DeploymentCSV: {
			BitNumber: 0,
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Unix(1456790400, 0), // March 1st, 2016
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Unix(1493596800, 0), // May 1st, 2017
			),
		},
		DeploymentSegwit: {
			BitNumber: 1,
			DeploymentStarter: NewMedianTimeDeploymentStarter(
				time.Unix(1462060800, 0), // May 1, 2016 UTC
			),
			DeploymentEnder: NewMedianTimeDeploymentEnder(
				time.Unix(1493596800, 0), // May 1, 2017 UTC.
			),
		},
	},

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID:        0x6f, // starts with m or n
	ScriptHashAddrID:        0xc4, // starts with 2
	WitnessPubKeyHashAddrID: 0x03, // starts with QW
	WitnessScriptHashAddrID: 0x28, // starts with T7n
	PrivateKeyID:            0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

var (
	// ErrDuplicateNet describes an error where the parameters for a Bitcoin
	// network could not be set due to the network already being a standard
	// network or previously-registered into this package.
	ErrDuplicateNet = errors.New("duplicate Bitcoin network")

	// ErrUnknownHDKeyID describes an error where the provided id which
	// is intended to identify the network for a hierarchical deterministic
	// private extended key is not registered.
	ErrUnknownHDKeyID = errors.New("unknown hd private extended key bytes")

	// ErrInvalidHDKeyID describes an error where the provided hierarchical
	// deterministic version bytes, or hd key id, is malformed.
	ErrInvalidHDKeyID = errors.New("invalid hd extended key version bytes")
)

var (
	registeredNets     = make(map[wire.BitcoinNet]struct{})
	pubKeyHashAddrIDs  = make(map[byte]struct{})
	scriptHashAddrIDs  = make(map[byte]struct{})
	stakingHashAddrIDs = make(map[byte]struct{})
	hdPrivToPubKeyIDs  = make(map[[4]byte][]byte)
)

// String returns the hostname of the DNS seed in human-readable form.
func (d DNSSeed) String() string {
	return d.Host
}

// Register registers the network parameters for a Bitcoin network.  This may
// error with ErrDuplicateNet if the network is already registered (either
// due to a previous Register call, or the network being one of the default
// networks).
//
// Network parameters should be registered into this package by a main package
// as early as possible.  Then, library packages may lookup networks or network
// parameters based on inputs and work regardless of the network being standard
// or not.
func Register(params *Params) error {
	if _, ok := registeredNets[params.Net]; ok {
		return ErrDuplicateNet
	}
	registeredNets[params.Net] = struct{}{}
	pubKeyHashAddrIDs[params.PubKeyHashAddrID] = struct{}{}
	scriptHashAddrIDs[params.ScriptHashAddrID] = struct{}{}
	stakingHashAddrIDs[params.StakingKeyID] = struct{}{}

	err := RegisterHDKeyID(params.HDPublicKeyID[:], params.HDPrivateKeyID[:])
	if err != nil {
		return err
	}

	return nil
}

func mustRegister(params *Params) {
	if err := Register(params); err != nil {
		panic("failed to register network: " + err.Error())
	}
}

func IsPubKeyHashAddrID(id byte) bool {
	_, ok := pubKeyHashAddrIDs[id]
	return ok
}

func IStakingAddrID(id byte) bool {
	_, ok := stakingHashAddrIDs[id]
	return ok
}

func IsScriptHashAddrID(id byte) bool {
	_, ok := scriptHashAddrIDs[id]
	return ok
}

func RegisterHDKeyID(hdPublicKeyID []byte, hdPrivateKeyID []byte) error {
	if len(hdPublicKeyID) != 4 || len(hdPrivateKeyID) != 4 {
		return ErrInvalidHDKeyID
	}

	var keyID [4]byte
	copy(keyID[:], hdPrivateKeyID)
	hdPrivToPubKeyIDs[keyID] = hdPublicKeyID

	return nil
}

func HDPrivateKeyToPublicKeyID(id []byte) ([]byte, error) {
	if len(id) != 4 {
		return nil, ErrUnknownHDKeyID
	}

	var key [4]byte
	copy(key[:], id)
	pubBytes, ok := hdPrivToPubKeyIDs[key]
	if !ok {
		return nil, ErrUnknownHDKeyID
	}

	return pubBytes, nil
}

func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		// Ordinarily I don't like panics in library code since it
		// can take applications down without them having a chance to
		// recover which is extremely annoying, however an exception is
		// being made in this case because the only way this can panic
		// is if there is an error in the hard-coded hashes.  Thus it
		// will only ever potentially panic on init and therefore is
		// 100% predictable.
		panic(err)
	}
	return hash
}

func init() {
	// Register all default networks when the package is initialized.
	mustRegister(&MainNetParams)
	mustRegister(&TestNet3Params)
	mustRegister(&RegressionNetParams)
}
