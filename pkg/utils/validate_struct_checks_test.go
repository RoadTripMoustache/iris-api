package utils

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// ----- checkMin ----- //
func Test_checkMin_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMin,
	}

	t.Run("Equal to min", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMin{Toto: 0})

		assert.Nil(t, result)
	})

	t.Run("Higher than min", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMin{Toto: 1})

		assert.Nil(t, result)
	})

}

func Test_checkMin_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMin,
	}

	result := ValidateStruct(mocks.FakeMin{Toto: -1})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"the value of the field Toto must be equal or over than 0"})
}

// ----- checkMax ----- //
func Test_checkMax_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMax,
	}

	t.Run("Equal to max", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMax{Toto: 10})

		assert.Nil(t, result)
	})

	t.Run("Lower than max", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMax{Toto: 9})

		assert.Nil(t, result)
	})

}

func Test_checkMax_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMax,
	}

	result := ValidateStruct(mocks.FakeMax{Toto: 1112})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"the value of the field Toto must be equal or lower than 10"})
}

// ----- checkMinLength ----- //
func Test_checkMinLength_string_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMinLength,
	}

	t.Run("Equal to min length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMinLengthString{Toto: "10"})

		assert.Nil(t, result)
	})

	t.Run("Higher than min length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMinLengthString{Toto: "dddd"})

		assert.Nil(t, result)
	})

}

func Test_checkMinLength_string_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMinLength,
	}

	result := ValidateStruct(mocks.FakeMinLengthString{Toto: "d"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"Toto is too short. Min 2 characters"})
}

func Test_checkMinLength_slice_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMinLength,
	}

	t.Run("Equal to min length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMinLengthSlice{Toto: []string{"d", "n"}})

		assert.Nil(t, result)
	})

	t.Run("Higher than minlength", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMinLengthSlice{Toto: []string{"d", "n", "d"}})

		assert.Nil(t, result)
	})

}

func Test_checkMinLength_slice_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMinLength,
	}

	result := ValidateStruct(mocks.FakeMinLengthSlice{Toto: []string{"d"}})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"Toto is too short. Min 2 entries"})
}

// ----- checkMaxLength ----- //
func Test_checkMaxLength_string_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMaxLength,
	}

	t.Run("Equal to max length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMaxLengthString{Toto: "10"})

		assert.Nil(t, result)
	})

	t.Run("Lower than max length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMaxLengthString{Toto: "s"})

		assert.Nil(t, result)
	})

}

func Test_checkMaxLength_string_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMaxLength,
	}

	result := ValidateStruct(mocks.FakeMaxLengthString{Toto: "dddd"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"Toto is too long. Max 2 characters"})
}

func Test_checkMaxLength_slice_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMaxLength,
	}

	t.Run("Equal to max length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMaxLengthSlice{Toto: []string{"d", "n"}})

		assert.Nil(t, result)
	})

	t.Run("Lower than max length", func(t *testing.T) {
		result := ValidateStruct(mocks.FakeMaxLengthSlice{Toto: []string{"d"}})

		assert.Nil(t, result)
	})

}

func Test_checkMaxLength_slice_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkMaxLength,
	}

	result := ValidateStruct(mocks.FakeMaxLengthSlice{Toto: []string{"d", "d", "d"}})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"Toto is too long. Max 2 entries"})
}

// ----- checkPattern ----- //
func Test_checkPattern_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkPattern,
	}

	result := ValidateStruct(mocks.FakePattern{Toto: "QQQ"})

	assert.Nil(t, result)

}

func Test_checkPattern_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkPattern,
	}

	result := ValidateStruct(mocks.FakePattern{Toto: "dddd"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"the value of the field Toto doesn't match the expected pattern"})
}

// ----- checkDatePattern ----- //
func Test_checkDatePattern_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkDatePattern,
	}

	result := ValidateStruct(mocks.FakeDatePattern{Toto: "2024-01-02T00:00:00Z"})

	assert.Nil(t, result)

}

func Test_checkDatePattern_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkDatePattern,
	}

	result := ValidateStruct(mocks.FakeDatePattern{Toto: "dddd"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"the value of the field Toto doesn't match the expected pattern"})
}

// ----- checkEnum ----- //
func Test_checkEnum_string_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkEnum,
	}

	result := ValidateStruct(mocks.FakeEnumString{Toto: "toto"})

	assert.Nil(t, result)

}

func Test_checkEnum_string_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkEnum,
	}

	result := ValidateStruct(mocks.FakeEnumString{Toto: "dddd"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"field Toto contains value which are not allowed - dddd"})
}

func Test_checkEnum_slice_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkEnum,
	}

	result := ValidateStruct(mocks.FakeEnumSlice{Toto: []string{"toto", "tata"}})

	assert.Nil(t, result)

}

func Test_checkEnum_slice_nok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkEnum,
	}

	result := ValidateStruct(mocks.FakeEnumSlice{Toto: []string{"toto", "dddd"}})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"field Toto contains value which are not allowed - dddd"})
}

// ----- Check multiples ----- //
func Test_checkMultiples(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkEnum,
		checkMinLength,
	}

	result := ValidateStruct(mocks.FakeMultiple{
		Tata: nil,
		Toto: []string{"toto", "dddd"},
	})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{
		"Tata is too short. Min 2 entries",
		"field Toto contains value which are not allowed - dddd",
	})
}
