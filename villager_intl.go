package main

import "golang.org/x/text/language"

type villagerIntl struct {
	Language language.Tag
	Value    string
}

type villagersIntlNames struct {
	Names []villagerIntl
}

type villagersIntlPhrases struct {
	Phrases []villagerIntl
}
