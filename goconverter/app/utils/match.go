package utils

func StartsWithCurlyBraceAndEndsWithClosingCurlyBrace(text string) bool {
	return len(text) >= 2 && text[0] == '{' && text[len(text)-1] == '}'
}
