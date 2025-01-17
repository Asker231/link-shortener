package link

import (
	"math/rand"
	"gorm.io/gorm"
)


type Link struct{
	*gorm.Model
	URL string `json:"url"`
	Hash string `json:"hash"`
}

func NewLink(url string)*Link{
	return &Link{
		URL: url,
		Hash: GenerateHash(6),
	}
}

func GenerateHash(n int)string{
	var letter []rune = []rune("abcdefgrtqwkjhzm123456789@")
	word := make([]rune,n)

	for i := range word{
		word[i] = letter[rand.Intn(len(letter))]
	}
	return string(word)
}