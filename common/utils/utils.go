package utils

import (
	"encoding/json"
	"os"
)

func RemoveElementFromSlice[T comparable](slice *[]T, elementToRemove T) {
	for i, v := range *slice {
		if v == elementToRemove {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
		}
	}
}

func GetEnv(name string, fallback string) string {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}
	return val
}

func SerializeJson(data interface{}) []byte {
	res, _ := json.Marshal(data)
	return res
}

func ParseJson(data []byte, dst interface{}) error {
	err := json.Unmarshal(data, dst)
	return err
}

func Min(a float64, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}

func Max(a float64, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func Pow2(x float64) float64 {
	return x * x
}
