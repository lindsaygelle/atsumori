package main

type villagerFear uint8

type villagersFear struct {
	Fear villagerFear
}

const (
	villagerFearBugMask villagerFear = iota
	villagerFearMonsterMask
	villagerFearMummyMask
	villagerFearSkeletonHood
	villagerFearWerewolfHood
)
