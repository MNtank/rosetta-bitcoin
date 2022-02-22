// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package bitcoin

import (
	"math/big"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func init() {
	if err := chaincfg.Register(&EunoMainNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&EunoTestNet3Params); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&EunoRegressionNetParams); err != nil {
		panic(err)
	}
}

var (
	bigOne = big.NewInt(1)

	PowLimit   = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
	PowLimitV1 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 24), bigOne)
	PowLimitV2 = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
)

var EunoMainNetParams = chaincfg.Params{
	Name:        "main",
	Net:         0xe9fdc490,
	DefaultPort: "46462",

	// Chain parameters
	GenesisBlock: &genesisBlock,
	GenesisHash:  &genesisHash,

	// Address encoding magics
	PubKeyHashAddrID:        139,
	ScriptHashAddrID:        19,
	PrivateKeyID:            239,
	WitnessPubKeyHashAddrID: 0x00,
	WitnessScriptHashAddrID: 0x00,

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
var EunoRegressionNetParams = chaincfg.Params{
	Name: "regtest",

	// DigiByte has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(digibyte)), so that we can
	// register the regtest network.
	// DigiByte Core Developers will change this soon.
	Net:         0xac7ecfa1,
	DefaultPort: "46466",

	// Chain parameters
	GenesisBlock: &regTestGenesisBlock,
	GenesisHash:  &regTestGenesisHash,

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 303,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var EunoTestNet3Params = chaincfg.Params{
	Name: "test",

	// DigiByte has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(digibyte)), so that we can
	// register the regtest network.
	// DigiByte Core Developers will change this soon.
	Net:         0xba657645,
	DefaultPort: "46464",

	// Chain parameters
	GenesisBlock: &testNet3GenesisBlock,
	GenesisHash:  &testNet3GenesisHash,

	// Address encoding magics
	PubKeyHashAddrID:        139,  // starts with 1
	ScriptHashAddrID:        19,   // starts with 3
	PrivateKeyID:            239,  // starts with 5 (uncompressed) or K (compressed)
	WitnessPubKeyHashAddrID: 0x00, // starts with p2
	WitnessScriptHashAddrID: 0x00, // starts with 7Xh

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x3a, 0x80, 0x58, 0x37}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x3a, 0x80, 0x61, 0xa0}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 303,
}

func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		panic(err)
	}
	return hash
}
