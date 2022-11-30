package random

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/exp/constraints"
)

const (
	alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num   = "12345567890"
)

func TimeUTCNoNano() time.Time {
	t := Int64n(time.Now().Unix())
	return time.Unix(t, 0).UTC()
}

func UUID() uuid.UUID {
	return uuid.New()
}

func StatusString() string {
	return OneOf([]string{"PAID", "PENDING"})
}

func String() string {
	return StringAlphanum(Intn(100))
}

func StringOfType[T ~string]() T {
	return T(String())
}

func StringOfTypeBetween[T ~string](min, max int) T {
	diff := max - min
	return T(StringAlphanum(min + Intn(diff)))
}

func StringBetween(min, max int) string {
	diff := max - min
	return StringAlphanum(min + Intn(diff))
}

func StringAlpha(len int) string {
	return StringOf(alpha, len)
}

func StringNum(len int) string {
	return StringOf(num, len)
}

func StringAlphanum(len int) string {
	return StringOf(alpha+num, len)
}

func StringOf(s string, len int) string {
	var res string

	for i := 0; i < len; i++ {
		res += string(Rune(s))
	}

	return res
}

func StringWithPrefix(prefix string) string {
	return prefix + "_" + StringAlphanum(20)
}

func NumberString(len int) string {
	nums := make([]string, len)
	for i := range nums {
		if i == 0 {
			nums[i] += ItemFrom("1", "2", "3", "4", "5", "6", "7", "8", "9")
		} else {
			nums[i] += ItemFrom("0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
		}
	}
	return strings.Join(nums, "")
}

func Error() error {
	return fmt.Errorf("random error! %s", StringAlpha(100))
}

func ErrorWithPrefix(prefix string) error {
	return fmt.Errorf("random error! %s", StringWithPrefix(prefix))
}

func Int() int {
	return Intn(1000)
}

func Intn[T constraints.Integer](n T) int {
	return newRand().Intn(int(n))
}

func Int64() int64 {
	return Int64n(1000)
}

func Int64n[T constraints.Integer](n T) int64 {
	return newRand().Int63n(int64(n))
}

func Bool() bool {
	return Intn(2) == 0
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func ItemFrom[T any](items ...T) T {
	return items[rand.Intn(len(items))]
}

func SplitIntoTwoGroups[T any](items []T) ([]T, []T) {
	group1 := []T{items[0]}
	group2 := []T{items[1]}

	for _, pair := range items[2:] {
		if Bool() {
			group1 = append(group1, pair)
		} else {
			group2 = append(group2, pair)
		}
	}

	return group1, group2
}

func NumInRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func ArrayOf[T any](create func() T, min, max int) []T {
	count := NumInRange(min, max)
	arr := make([]T, count)
	for i := 0; i < count; i++ {
		arr[i] = create()
	}
	return arr
}

func Rune(s string) rune {
	r := []rune(s)
	return r[Intn(len(r))]
}

func OneOf[T any](ts []T) T {
	return ts[Intn(len(ts))]
}
