package utils


import (
	"time"
	"math/rand"
	"strings"
	"encoding/json"
	"github.com/google/uuid"
)

func Contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
 }

func RemoveElementFromSlice(series []string, value string) []string{
	i := indexOf(value, series)
	if i != -1 {
		series[i] = series[len(series)-1] // Copy last element to index i.
		series[len(series)-1] = ""   // Erase last element (write zero value).
		series = series[:len(series)-1]   // Truncate slice.
	}
	return series
}

func RecordToJSONString(record RecordEvent) (string, error) {
	e, err := json.Marshal(record)
    if err != nil {
        return "", err
    }
    return string(e), nil
}

func GetUUID() string {
	uuidWithHyphen := uuid.New()
    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
    return uuid
}

func GetFormattedDateNow() string {
	t := time.Now()
	result := t.Format("2012-01-03T15:02:51.738Z")
	return result
}


func GetRandomIDInt() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(10000)
}