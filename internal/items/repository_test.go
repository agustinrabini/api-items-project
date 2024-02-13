package items

import (
	"testing"

	"github.com/agustinrabini/api-items-project/internal/platform/storage"

	"github.com/stretchr/testify/assert"
)

// TODO mock mongo DB and add repository tests

func TestRepository_NewRepository(t *testing.T) {
	db := storage.NewMock()
	defer db.Close()
	repository := NewRepository(db.NewCollection("test"))
	assert.NotNil(t, repository)
}
