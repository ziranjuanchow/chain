package ca

import "chain/crypto/ed25519/ecmath"

// AssetCommitment is a point pair representing an ElGamal commitment to an AssetPoint.
type AssetCommitment PointPair

// CreateAssetCommitment creates an asset commitment. Nil aek means
// make it nonblinded (and the returned scalar blinding factor is
// nil).
func CreateAssetCommitment(assetID AssetID, aek AssetKey) (*AssetCommitment, *ecmath.Scalar) {
	if aek == nil {
		A := ecmath.Point(CreateAssetPoint(assetID))
		return &AssetCommitment{A, ecmath.ZeroPoint}, nil
	}
	c := scalarHash("ChainCA.AC.c", assetID[:], aek)
	return createRawAssetCommitment(assetID, &c), &c
}

func createRawAssetCommitment(assetID AssetID, c *ecmath.Scalar) *AssetCommitment {
	A := ecmath.Point(CreateAssetPoint(assetID))
	var H, C ecmath.Point
	H.ScMulAdd(&A, &ecmath.One, c)
	C.ScMul(&J, c)
	return &AssetCommitment{H, C}
}

func (ac *AssetCommitment) Bytes() []byte {
	return (*PointPair)(ac).Bytes()
}