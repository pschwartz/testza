package testza

import (
	"math/rand"
	"testing"

	"github.com/MarvinJWendt/testza/internal"
)

// MockInputsStringsHelper contains strings test sets.
type MockInputsStringsHelper struct{}

// Empty returns a test set with a single empty string.
func (s MockInputsStringsHelper) Empty() []string {
	return []string{""}
}

// Long returns a test set with long random strings.
// Returns:
// - Random string (length: 25)
// - Random string (length: 50)
// - Random string (length: 100)
// - Random string (length: 1,000)
// - Random string (length: 100,000)
func (s MockInputsStringsHelper) Long() (testSet []string) {
	testSet = append(testSet, s.GenerateRandom(1, 25)...)
	testSet = append(testSet, s.GenerateRandom(1, 50)...)
	testSet = append(testSet, s.GenerateRandom(1, 100)...)
	testSet = append(testSet, s.GenerateRandom(1, 1_000)...)
	testSet = append(testSet, s.GenerateRandom(1, 100_000)...)

	return
}

// Numeric returns a test set with strings that are numeric.
// The highest number in here is "9223372036854775807", which is equal to the maxmim int64.
func (s MockInputsStringsHelper) Numeric() []string {
	positiveNumbers := []string{"0", "1", "2", "3", "100", "1.1", "1337", "13.37", "0.000000000001", "9223372036854775807"}
	negativeNumbers := s.Modify(positiveNumbers, func(index int, value string) string { return "-" + value })
	return append(positiveNumbers, negativeNumbers...)
}

// Usernames returns a test set with usernames.
func (s MockInputsStringsHelper) Usernames() []string {
	return []string{"MarvinJWendt", "Zipper1337", "n00b", "l33t", "j0rgan", "test", "test123", "TEST", "test_", "TEST_"}
}

// EmailAddresses returns a test set with valid email addresses.
func (s MockInputsStringsHelper) EmailAddresses() []string {
	return []string{
		"hello@world.com",
		"hello+world@example.com",
		"hello.world@example.com",
		"a@a.xyz",
		"test@127.0.0.1",
		"test@[127.0.0.1]",
		"1@example.com",
		"_____@example.com",
		"test@subdomain.domain.xyz",
		`valid.”email\ address@example.com`,
		`first.last@iana.org`,
		`1234567890123456789012345678901234567890123456789012345678901234@iana.org`,
		`"first\"last"@iana.org`,
		`"first@last"@iana.org`,
		`"first\\last"@iana.org`,
		`x@x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x23456789.x2`,
		`1234567890123456789012345678901234567890123456789012345678@12345678901234567890123456789012345678901234567890123456789.12345678901234567890123456789012345678901234567890123456789.123456789012345678901234567890123456789012345678901234567890123.iana.org`,
		`first.last@[12.34.56.78]`,
		`first.last@[IPv6:::12.34.56.78]`,
		`first.last@[IPv6:::b3:b4]`,
		`first.last@[IPv6:::]`,
		`first.last@[IPv6:1111:2222:3333::4444:12.34.56.78]`,
		`"first\last"@iana.org`,
		`user+mailbox@iana.org`,
		`customer/department@iana.org `,
		`customer/department=shipping@iana.org`,
		`"Doug \"Ace\" L."@iana.org`,
		`+1~1+@iana.org`,
		`{_test_}@iana.org`,
		`"[[ test ]]"@iana.org`,
		`"test&#13;&#10; blah"@iana.org`,
		`(foo)cal(bar)@(baz)iamcal.com(quux)`,
		`cal(woo(yay)hoopla)@iamcal.com`,
		`cal(foo\@bar)@iamcal.com`,
		`cal(foo\)bar)@iamcal.com`,
		`first(Welcome to&#13;&#10; the ("wonderful" (!)) world&#13;&#10; of email)@iana.org`,
		`pete(his account)@silly.test(his host)`,
		`c@(Chris's host.)public.example`,
	}
}

// HtmlTags returns a test set with html tags.
func (s MockInputsStringsHelper) HtmlTags() []string {
	return []string{
		"<script>alert('XSS')</script>",
		"<script>",
		`<a href="https://github.com/MarvinJWendt/testza">link</a>`,
		`</body>`,
		`</html>`,
	}
}

// Full contains all string test sets plus ten generated random strings.
func (s MockInputsStringsHelper) Full() (ret []string) {
	ret = append(ret, s.Usernames()...)
	ret = append(ret, s.HtmlTags()...)
	ret = append(ret, s.EmailAddresses()...)
	ret = append(ret, s.Empty()...)
	ret = append(ret, s.Numeric()...)
	ret = append(ret, s.Long()...)

	for i := 0; i < 10; i++ {
		ret = append(ret, s.GenerateRandom(1, i)...)
	}

	return
}

// Limit limits a test set in size.
func (s MockInputsStringsHelper) Limit(testSet []string, max int) []string {
	if len(testSet) <= max {
		return testSet
	}

	if max <= 0 {
		return []string{}
	}

	return testSet[:max]
}

// GenerateRandom returns random StringsHelper in a test set.
func (s MockInputsStringsHelper) GenerateRandom(count, length int) (result []string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := 0; i < count; i++ {
		str := make([]rune, length)

		for i := range str {
			str[i] = letters[rand.Intn(len(letters))]
		}
		result = append(result, string(str))
	}
	return
}

// RunTests runs a test for every value in a testset.
// You can use the value as input parameter for your functions, to sanity test against many different cases.
// This ensures that your functions have a correct error handling and enables you to test against hunderts of cases easily.
func (s MockInputsStringsHelper) RunTests(t testRunner, testSet []string, testFunc func(t *testing.T, index int, str string)) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	test := internal.GetTest(t)
	if test == nil {
		t.Error(internal.ErrCanNotRunIfNotBuiltinTesting)
		return
	}

	for i, str := range testSet {
		test.Run(str, func(t *testing.T) {
			t.Helper()

			testFunc(t, i, str)
		})
	}
}

// Modify returns a modified version of a test set.
func (s MockInputsStringsHelper) Modify(inputSlice []string, f func(index int, value string) string) (ret []string) {
	for i, str := range inputSlice {
		ret = append(ret, f(i, str))
	}

	return
}
