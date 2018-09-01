package convt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Ps pointer str to string
func Ps(ps *string) string {
	return *ps
}

// PS string to poiter string
func PS(ps string) *string {
	return &ps
}

// Pi pointer int toint
func Pi(ps *int) int {
	return *ps
}

// PI int to poiter int
func PI(ps int) *int {
	return &ps
}

// PIf2Str any to string
func PIf2Str(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// PIf2Int any to number
func PIf2Int(i interface{}) (int, error) {
	return strconv.Atoi(PIf2Str(i))
}

// StrDate2TimeStamp string date to timestamp
func StrDate2TimeStamp(str string) (int, error) {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return 0, err
	}
	return int(t.UnixNano() / 1000 / 1000), nil
}

// CreateError string date to timestamp
func CreateError(any string) error {
	return errors.New(any)
}

// CreateID for key
func CreateID(prefix string) string {
	hash, _ := uuid.NewUUID()
	return prefix + strings.Replace(hash.String(), "-", ".", -1)
}
