// Copyright 2018 The Gringo Developers. All rights reserved.
// Use of this source code is governed by a GNU GENERAL PUBLIC LICENSE v3
// license that can be found in the LICENSE file.

package chain

import (
	"consensus"
	"sync"
	"container/list"
	"bytes"
	"time"
)

var Testnet1 = Chain{
	genesis: consensus.Block{
		Header: consensus.BlockHeader{
			Version:         1,
			Height:          0,
			Previous:        bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
			Timestamp:       time.Date(2017, 11, 16, 20, 0, 0, 0, time.UTC),
			Difficulty:      10,
			TotalDifficulty: 10,

			UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
			RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
			KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

			Nonce: 28205,
			POW: consensus.Proof{
				Nonces: []uint32{
					0x21e, 0x7a2, 0xeae, 0x144e, 0x1b1c, 0x1fbd,
					0x203a, 0x214b, 0x293b, 0x2b74, 0x2bfa, 0x2c26,
					0x32bb, 0x346a, 0x34c7, 0x37c5, 0x4164, 0x42cc,
					0x4cc3, 0x55af, 0x5a70, 0x5b14, 0x5e1c, 0x5f76,
					0x6061, 0x60f9, 0x61d7, 0x6318, 0x63a1, 0x63fb,
					0x649b, 0x64e5, 0x65a1, 0x6b69, 0x70f8, 0x71c7,
					0x71cd, 0x7492, 0x7b11, 0x7db8, 0x7f29, 0x7ff8,
				},
			},
		},
	},
}

var Testnet2 = Chain{
	genesis: consensus.Block{
		Header: consensus.BlockHeader{
			Version:         1,
			Height:          0,
			Previous:        bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
			Timestamp:       time.Date(2017, 11, 16, 20, 0, 0, 0, time.UTC),
			//Difficulty:      10,
			//TotalDifficulty: 10,

			UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
			RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
			KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

			Nonce: 70081,
			POW: consensus.Proof{
				Nonces: []uint32{
					0x43ee48, 0x18d5a49, 0x2b76803, 0x3181a29, 0x39d6a8a, 0x39ef8d8,
					0x478a0fb, 0x69c1f9e, 0x6da4bca, 0x6f8782c, 0x9d842d7, 0xa051397,
					0xb56934c, 0xbf1f2c7, 0xc992c89, 0xce53a5a, 0xfa87225, 0x1070f99e,
					0x107b39af, 0x1160a11b, 0x11b379a8, 0x12420e02, 0x12991602, 0x12cc4a71,
					0x13d91075, 0x15c950d0, 0x1659b7be, 0x1682c2b4, 0x1796c62f, 0x191cf4c9,
					0x19d71ac0, 0x1b812e44, 0x1d150efe, 0x1d15bd77, 0x1d172841, 0x1d51e967,
					0x1ee1de39, 0x1f35c9b3, 0x1f557204, 0x1fbf884f, 0x1fcf80bf, 0x1fd59d40,
				},
			},
		},
	},
}

var Mainnet = Chain{
	genesis: consensus.Block{
		Header: consensus.BlockHeader{
			Version:         1,
			Height:          0,
			Previous:        bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
			Timestamp:       time.Date(2018, 8, 14, 0, 0, 0, 0, time.UTC),
			Difficulty:      1000,
			TotalDifficulty: 1000,

			UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
			RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
			KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

			Nonce: 28205,
			POW: consensus.Proof{
				Nonces: []uint32{
					0x21e, 0x7a2, 0xeae, 0x144e, 0x1b1c, 0x1fbd,
					0x203a, 0x214b, 0x293b, 0x2b74, 0x2bfa, 0x2c26,
					0x32bb, 0x346a, 0x34c7, 0x37c5, 0x4164, 0x42cc,
					0x4cc3, 0x55af, 0x5a70, 0x5b14, 0x5e1c, 0x5f76,
					0x6061, 0x60f9, 0x61d7, 0x6318, 0x63a1, 0x63fb,
					0x649b, 0x64e5, 0x65a1, 0x6b69, 0x70f8, 0x71c7,
					0x71cd, 0x7492, 0x7b11, 0x7db8, 0x7f29, 0x7ff8,
				},
			},
		},
	},
}

type Chain struct {
	sync.RWMutex

	// Genesis block
	genesis consensus.Block

	blockHeaders list.List
}

// Genesis returns genesis block
func (c *Chain) Genesis() consensus.Block {
	return c.genesis
}

func (c *Chain) TotalDifficulty() consensus.Difficulty {
	return consensus.ZeroDifficulty
}

func (c *Chain) Height() uint64 {
	return 0
}

func (c *Chain) GetBlockHeaders(loc consensus.Locator) []consensus.BlockHeader {
	return nil
}

func (c *Chain) GetBlock(hash consensus.BlockHash) *consensus.Block {
	return nil
}

func (c *Chain) ProcessHeaders(headers []consensus.BlockHeader) error {
	return nil
}

func (c *Chain) ProcessBlock(block *consensus.Block) error {
	return nil
}
