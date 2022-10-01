package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	regex := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`

	re := regexp.MustCompile(regex)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex := `<[*~=-]*>`

	re := regexp.MustCompile(regex)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := `(?i)".*(password).*"`
	re := regexp.MustCompile(regex)

	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := `end-of-line[0-9]*`
	re := regexp.MustCompile(regex)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	regex := `User\s+(\w+)`
	re := regexp.MustCompile(regex)

	result := []string{}
	for _, line := range lines {
		if m := re.FindString(line); len(m) > 0 {
			name := strings.TrimSpace(m[4:])
			line = fmt.Sprintf("[USR] %s %s", name, line)
		}
		result = append(result, line)
	}
	return result
}
