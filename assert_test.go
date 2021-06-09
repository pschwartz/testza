package testutil

import (
	"fmt"
	"go/types"
	"math/rand"
	"testing"
)

type assertionTestStruct struct {
	Name string
	Age  int
	Meta assertionTestStructNested
}

type assertionTestStructNested struct {
	ID    int
	Admin bool
}

// var testStruct = assertionTestStruct{Name: "Marvin Wendt", Age: 20, Meta: assertionTestStructNested{ID: 1337, Admin: true}}
// var testStructDifferent = assertionTestStruct{Name: "Marvin Wendt", Age: 20, Meta: assertionTestStructNested{ID: 7331, Admin: true}}

func randomString() string {
	return Input.Strings.GenerateRandom(rand.Intn(10), 1)[0]
}

func generateStruct() (ret assertionTestStruct) {
	ret.Name = randomString()
	ret.Age = rand.Intn(40)
	ret.Meta.ID = rand.Intn(10_000)
	ret.Meta.Admin = rand.Intn(1) == 1

	return
}

func testEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	t.Run("Equal", func(t *testing.T) {
		t.Helper()

		AssertHelper.Equal(t, expected, actual)
	})

	t.Run("EqualValues", func(t *testing.T) {
		t.Helper()

		AssertHelper.EqualValues(t, expected, actual)
	})
}

func TestAssert_Equal(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Input.Strings.RunTests(t, Input.Strings.All(), func(t *testing.T, index int, str string) {
			t.Helper()

			testEqual(t, str, str)
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			testEqual(t, s, s)
		}
	})
}

func TestAssert_NotEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Input.Strings.RunTests(t, Input.Strings.All(), func(t *testing.T, index int, str string) {
			AssertHelper.NotEqual(t, str, str+" addon")
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				AssertHelper.NotEqual(t, s, s2)
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				AssertHelper.NotEqualValues(t, s, s2)
			})
		}
	})
}

func TestAssert_EqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Input.Strings.RunTests(t, Input.Strings.All(), func(t *testing.T, index int, str string) {
			AssertHelper.EqualValues(t, str, str)
		})
	})
}

func TestAssert_NotEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		Input.Strings.RunTests(t, Input.Strings.All(), func(t *testing.T, index int, str string) {
			AssertHelper.NotEqualValues(t, str, str+" addon")
		})
	})
}

func TestAssert_Implements(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		AssertHelper.Implements(t, (*fmt.Stringer)(nil), new(types.Const))
	})
}

func TestAssert_NotImplements(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		AssertHelper.NotImplements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
	})
}

func TestFail(t *testing.T) {

}