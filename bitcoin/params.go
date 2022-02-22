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
	bigOne   = big.NewInt(1)
	PowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 20), bigOne)
)

var EunoMainNetParams = chaincfg.Params{
	Name:        "main",
	Net:         0xe9fdc490,
	DefaultPort: "46462",

	// Chain parameters
	GenesisBlock: &genesisBlock,
	GenesisHash:  &genesisHash,

	// Address encoding magics
	PubKeyHashAddrID:        0x21,
	ScriptHashAddrID:        0x11,
	PrivateKeyID:            0x9,
	WitnessPubKeyHashAddrID: 0x00,
	WitnessScriptHashAddrID: 0x00,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x3a, 0x80, 0x58, 0x37}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x3a, 0x80, 0x61, 0xa0}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 0x12F,
}

var EunoRegressionNetParams = chaincfg.Params{
	Name: "regtest",

	Net:         0xac7ecfa1,
	DefaultPort: "46466",

	// Chain parameters
	GenesisBlock: &regTestGenesisBlock,
	GenesisHash:  &regTestGenesisHash,

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 0x12F,
}

var EunoTestNet3Params = chaincfg.Params{
	Name: "test",

	Net:         0xba657645,
	DefaultPort: "46464",

	// Chain parameters
	GenesisBlock: &testNet3GenesisBlock,
	GenesisHash:  &testNet3GenesisHash,

	// Address encoding magics
	PubKeyHashAddrID:        0x21,
	ScriptHashAddrID:        0x11,
	PrivateKeyID:            0x9,
	WitnessPubKeyHashAddrID: 0x00,
	WitnessScriptHashAddrID: 0x00,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x3a, 0x80, 0x58, 0x37},
	HDPublicKeyID:  [4]byte{0x3a, 0x80, 0x61, 0xa0},

	HDCoinType: 0x12F,
}

func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		panic(err)
	}
	return hash
}
