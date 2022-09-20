package random

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/cnlesscode/gotool/slice"
	"github.com/google/uuid"
)

// Create an integer random number
func RangeIntRand(min, max int64) int64 {
	if min > max {
		return min
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

// Create an integer random float
func RangeFloat(IntegerRange, FloatRange []int) float64 {
	integerVal := RangeIntRand(int64(IntegerRange[0]), int64(IntegerRange[1]))
	floatVal := RangeIntRand(int64(FloatRange[0]), int64(FloatRange[1]))
	stringVal := strconv.Itoa(int(integerVal)) + "." + strconv.Itoa(int(floatVal))
	result, err := strconv.ParseFloat(stringVal, 64)
	if err != nil {
		return 0
	} else {
		return result
	}
}

// Creates a random character of a specified length
func RandomCharacters(length int, intLength int) string {
	if intLength >= length {
		intLength = 0
	}
	var letters string = "abcdefghjkmnpqrstwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]string, 0)
	for i := 0; i < length-intLength; i++ {
		randNum := RangeIntRand(0, 46)
		result = append(result, letters[randNum:randNum+1])
	}
	for i := 0; i < intLength; i++ {
		result = append(result, strconv.Itoa(int(RangeIntRand(0, 9))))
	}
	result = slice.SortRandomlyString(result)
	return strings.Join(result, "")
}

// UUID
func UUID() string {
	return uuid.NewString()
}

// md5
func Md5UUID() string {
	uuid := uuid.NewString()
	md5in := md5.New()
	md5in.Write([]byte(uuid))
	return hex.EncodeToString(md5in.Sum(nil))
}
