package leap

// TestVersion The test version.
const TestVersion = 1

// IsLeapYear Determines whether or not a year is a leap year, ie, it is a multiple of 4 but not of
// 100, unless it is also a multiple of 400, in which case it reamins a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
