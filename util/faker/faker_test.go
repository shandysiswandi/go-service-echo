package faker_test

import (
	"go-service-echo/util/faker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sentence(t *testing.T) {
	assert.GreaterOrEqual(t, len(faker.Sentence()), 25)
}

func Test_Paragraph(t *testing.T) {
	assert.GreaterOrEqual(t, len(faker.Paragraph()), 50)
}

func Test_Email(t *testing.T) {
	assert.GreaterOrEqual(t, len(faker.Email()), 10)
}

func Test_Name(t *testing.T) {
	assert.GreaterOrEqual(t, len(faker.Name()), 10)
}
