package compare

import (
	"os"
	"strings"
	"time"

	"github.com/kshvmdn/fsql/query"
)

// Alpha compares two strings a and b.
func Alpha(comp query.TokenType, a, b string) bool {
	switch comp {
	case query.BeginsWith:
		return strings.HasPrefix(a, b)
	case query.EndsWith:
		return strings.HasSuffix(a, b)
	case query.Equals:
		fallthrough
	case query.Is:
		return a == b
	case query.Contains:
		return strings.Contains(a, b)
	}
	return false
}

// Numeric compares two integers a and b.
func Numeric(comp query.TokenType, a, b int64) bool {
	switch comp {
	case query.Equals:
		return a == b
	case query.NotEquals:
		return a != b
	case query.GreaterThanEquals:
		return a >= b
	case query.GreaterThan:
		return a > b
	case query.LessThanEquals:
		return a <= b
	case query.LessThan:
		return a < b
	}
	return false
}

// Time compares two times a and b.
func Time(comp query.TokenType, a, b time.Time) bool {
	switch comp {
	case query.Equals:
		return a.Equal(b)
	case query.NotEquals:
		return !a.Equal(b)
	case query.GreaterThanEquals:
		return a.After(b) || a.Equal(b)
	case query.GreaterThan:
		return a.After(b)
	case query.LessThanEquals:
		return a.Before(b) || a.Equal(b)
	case query.LessThan:
		return a.Before(b)
	}
	return false
}

func File(comp query.TokenType, file os.FileInfo, filetype string) bool {
	switch comp {
	case query.Is:
		switch filetype {
		case "dir":
			return file.Mode().IsDir()
		case "reg":
			return file.Mode().IsRegular()
		}
	}
	return false
}
