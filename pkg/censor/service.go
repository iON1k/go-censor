package censor

import "strings"

// Сервис цензурирования
type Service interface {
	Validate(content string) bool
}

type ServiceImpl struct {
	bad_words []string
}

var common_bad_words = []string{
	"qwerty",
	"йцукен",
	"zxvbnm",
}

func New() *ServiceImpl {
	var bad_words []string
	for _, w := range common_bad_words {
		bad_words = append(bad_words, strings.ToLower(w))
	}
	return &ServiceImpl{bad_words}
}

func (s *ServiceImpl) Validate(content string) bool {
	l_content := strings.ToLower(content)
	for _, w := range s.bad_words {
		if strings.Contains(l_content, w) {
			return false
		}
	}

	return true
}
