package main

import (
	"crypto/cipher"
)

func DecryptAES128ECB(block cipher.Block, dst, src []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
}

func EncryptAES128ECB(block cipher.Block, dst, src []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
}

func DecryptAES128CBC(block cipher.Block, dst, src, iv []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	previous := iv
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])

		EncodeFixedXor(dst, dst, previous[:bs])

		previous = src[:bs]
		src = src[bs:]
		dst = dst[bs:]
	}
}

func EncryptAES128CBC(block cipher.Block, dst, src, iv []byte) {
	bs := block.BlockSize()

	if len(src)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	previous := iv
	for len(src) > 0 {
		EncodeFixedXor(src, src, previous[:bs])

		block.Encrypt(dst, src[:bs])

		previous = dst[:bs]
		src = src[bs:]
		dst = dst[bs:]
	}
}
