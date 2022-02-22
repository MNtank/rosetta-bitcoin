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

	PowLimit   = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
	PowLimitV1 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 24), bigOne)
	PowLimitV2 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
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
}

// Constants that define the deployment offset in the deployments field of the
// parameters for each deployment.  This is useful to be able to get the details
// of a specific deployment by name.
const (
	// DeploymentTestDummy defines the rule change deployment ID for testing
	// purposes.
	DeploymentTestDummy = iota

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
type Param struct {
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
}

// Constants used to indicate the message dogecoin network.
const (
	// MainNet represents the main dogecoin network.
	MainNet wire.BitcoinNet = 0xe9fdc490

	// TestNet3 represents the test network (version 3).
	TestNet wire.BitcoinNet = 0xba657645
)

// MainNetParams defines the network parameters for the main Bitcoin network.
var MainNetParams = Param{
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
	PowLimit:                    PowLimit,
	PowLimitV1:                  PowLimitV1,
	PowLimitV2:                  PowLimitV2,
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
			ActivationHeight: 0, //no activation height
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
var RegressionNetParams = Param{
	Name:        "regtest",
	Net:         wire.TestNet,
	DefaultPort: "46466",
	DNSSeeds:    []DNSSeed{},

	// Chain parameters
	GenesisBlock:                &regTestGenesisBlock,
	GenesisHash:                 &regTestGenesisHash,
	PowAllowMinDifficultyBlocks: true,
	PowLimit:                    PowLimit,
	PowLimitV1:                  PowLimitV1,
	PowLimitV2:                  PowLimitV2,
	BudgetCycleBlocks:           144,
	BudgetFeeConfirmations:      3,
	CoinbaseMaturity:            15,
	FutureTimeDriftPoW:          7200,
	FutureTimeDriftPoS:          180,
	PoolMaxTransactions:         2,
	ProposalEstablishmentTime:   time.Hour * 5,
	StakeMinAge:                 time.Hour * 60,
	StakeMinDepth:               100,
	TargetTimespan:              time.Minute * 40 * 60,
	TargetTimespanV2:            time.Minute * 30 * 60,
	TargetSpacing:               time.Minute * 1 * 60,
	TimeSlotLength:              15,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3Params = Param{
	Name:        "test",
	Net:         TestNet,
	DefaultPort: "46464",
	DNSSeeds: []DNSSeed{
		{"seed.euno.network", true},
	},

	// Chain parameters
	GenesisBlock:                &testNet3GenesisBlock,
	GenesisHash:                 &testNet3GenesisHash,
	PowAllowMinDifficultyBlocks: true,
	PowLimit:                    PowLimit,
	PowLimitV1:                  PowLimitV1,
	PowLimitV2:                  PowLimitV2,
	BudgetCycleBlocks:           144,
	BudgetFeeConfirmations:      3,
	CoinbaseMaturity:            15,
	FutureTimeDriftPoW:          7200,
	FutureTimeDriftPoS:          180,
	PoolMaxTransactions:         2,
	ProposalEstablishmentTime:   time.Hour * 5,
	StakeMinAge:                 time.Hour * 60,
	StakeMinDepth:               100,
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
	height_last_ZC_AccumCheckpoint:   251,
	height_last_ZC_WrappedSerials:    0,
	height_start_InvalidUTXOsCheck:   251,
	height_start_ZC_InvalidSerials:   251,
	height_start_ZC_SerialRangeCheck: 0,
	height_ZC_RecalcAccumulators:     251,

	// validation by-pass
	PivxBadBlockTime: 1489001494,
	PivxBadBlockBits: 0x1e0a20bd,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{0, newHashFromStr("00000055bbced95bf1fe0b4eccaee23ece9d82ceae67e3a4ec4e757619159a9f")},
	},

	Upgrades: [DefinedDeployments]ConsensusDeployment{
		UpgradeBaseNetwork: {
			ActivationHeight: 0,
		},
		UpgradeTestDummy: {
			ActivationHeight: 0,
		},
		UpgradePOS: {
			ActivationHeight: 251,
		},
		UpgradePOSV2: {
			ActivationHeight: 251,
		},
		UpgradeZC: {
			ActivationHeight:    251,
			hashActivationBlock: "258c489f42f03cb97db2255e47938da4083eee4e242853c2d48bae2b1d0110a6",
		},
		UpgradeZCV2: {
			ActivationHeight:    251,
			hashActivationBlock: "258c489f42f03cb97db2255e47938da4083eee4e242853c2d48bae2b1d0110a6",
		},
		UpgradeBIP65: {
			ActivationHeight:    251,
			hashActivationBlock: "258c489f42f03cb97db2255e47938da4083eee4e242853c2d48bae2b1d0110a6",
		},
		UpgradeZCPublic: {
			ActivationHeight: 251,
		},
		UpgradeV34: {
			ActivationHeight:    251,
			hashActivationBlock: "258c489f42f03cb97db2255e47938da4083eee4e242853c2d48bae2b1d0110a6",
		},
		UpgradeV40: {
			ActivationHeight:    251,
			hashActivationBlock: "258c489f42f03cb97db2255e47938da4083eee4e242853c2d48bae2b1d0110a6",
		},
		UpgradeV50Dummy: {
			ActivationHeight: 0, //no activation height
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
func Register(params *Param) error {
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

func mustRegister(params *Param) {
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
