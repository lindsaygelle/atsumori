package main

type villagerSkill uint16

type villagersSkill struct {
	Skill *villagerSkill
}

const (
	villagersSkillBellyDancing villagerSkill = iota
	villagersSkillCalligraphy
	villagersSkillCrammingForTests
	villagersSkillEatingQuickly
	villagersSkillFishing
	villagersSkillForgettingStuff
	villagersSkillKaraoke
	villagersSkillLimboing
	villagersSkillMakingCrafts
	villagersSkillOversleeping
	villagersSkillSkimmingBooks
	villagersSkillTiptoeing
	villagersSkillTrivia
	villagersSkillWakingUpEarly
	villagersSkillWritingBackwards
)
