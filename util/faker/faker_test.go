package faker_test

import (
	"go-service-echo/util/faker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentence(t *testing.T) {
	ac := len(faker.Sentence())

	assert.GreaterOrEqual(t, ac, 25)
}

func TestParagraph(t *testing.T) {
	ac := len(faker.Paragraph())

	assert.GreaterOrEqual(t, ac, 50)
}
