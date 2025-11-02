package admin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Admin_ToMap(t *testing.T) {
	result := admin.ToMap()

	assert.Equal(t, result, adminMap)
}

func Test_Admin_ToAdminModel(t *testing.T) {
	result := ToAdminModel(admin)

	assert.Equal(t, result, adminModel)
}
