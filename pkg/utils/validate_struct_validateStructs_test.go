package utils

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	workingCheck = func(t *testing.T) func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
		return func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
			_, ok := tag.Lookup("required")
			assert.True(t, ok)
			return nil
		}
	}
	failingCheck = func(t *testing.T) func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
		return func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
			assert.Fail(t, "this function must to be ignored")
			return nil
		}
	}
	checkReturningError = func(t *testing.T, errorMessage string) func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
		return func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string {
			_, ok := tag.Lookup("required")
			assert.True(t, ok)
			return &errorMessage
		}
	}
)

func Test_ValidateStruct_ok_withSubobject(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		workingCheck(t),
	}

	result := ValidateStruct(mocks.FakeWithSubObject{Toto: mocks.FakeRequiredValue{Toto: "toto"}})

	assert.Nil(t, result)
}

func Test_ValidateStruct_ok(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		workingCheck(t),
	}

	result := ValidateStruct(mocks.FakeRequiredValue{Toto: "toto"})

	assert.Nil(t, result)
}

func Test_ValidateStruct_nok_missingField(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		failingCheck(t),
	}

	result := ValidateStruct(mocks.FakeRequiredValue{})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"field Toto is mandatory"})
}

func Test_ValidateStruct_nok_missingField_subobject(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		failingCheck(t),
	}

	result := ValidateStruct(mocks.FakeWithSubObject{Toto: mocks.FakeRequiredValue{}})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"field Toto is mandatory"})
}

func Test_ValidateStruct_nok_firstCheckFails(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		checkReturningError(t, "failing test"),
		failingCheck(t),
	}

	result := ValidateStruct(mocks.FakeRequiredValue{Toto: "toto"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"failing test"})
}

func Test_ValidateStruct_nok_secondCheckFails(t *testing.T) {
	checkFunctions = []func(tag reflect.StructTag, v reflect.Value, i int, fieldName string) *string{
		workingCheck(t),
		checkReturningError(t, "failing test"),
		failingCheck(t),
	}

	result := ValidateStruct(mocks.FakeRequiredValue{Toto: "toto"})

	assert.NotNil(t, result)
	assert.Equal(t, result, []string{"failing test"})
}

// ----- isEmptyValue ----- //
func Test_isEmptyValue_ok(t *testing.T) {
	var interfaceValue interface{} = "t"
	testCases := []struct {
		Name  string
		Value interface{}
	}{
		{
			Name:  "String",
			Value: "r",
		},
		{
			Name:  "int",
			Value: 1,
		},
		{
			Name:  "bool",
			Value: true,
		},
		{
			Name:  "pointer",
			Value: &mocks.FakeMax{},
		},
		{
			Name:  "interface",
			Value: interfaceValue,
		},
		{
			Name:  "slice",
			Value: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			v := reflect.ValueOf(testCase.Value)

			result := isEmptyValue(v)

			assert.False(t, result)
		})
	}
}

func Test_isEmptyValue_nok(t *testing.T) {
	var pointerValue *mocks.FakeMax = nil
	var sliceValue []string
	testCases := []struct {
		Name  string
		Value interface{}
	}{
		{
			Name:  "String",
			Value: "",
		},
		{
			Name:  "bool",
			Value: false,
		},
		{
			Name:  "pointer",
			Value: pointerValue,
		},
		{
			Name:  "slice",
			Value: sliceValue,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			v := reflect.ValueOf(testCase.Value)

			result := isEmptyValue(v)

			assert.True(t, result)
		})
	}
}
