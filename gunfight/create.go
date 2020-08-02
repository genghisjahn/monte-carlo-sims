package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func createFighters(n int) (string, error) {
	gs := []*baseFigterStats{}
	for f := 0; f < n; f++ {
		gf := baseFigterStats{}
		gf.Name = generateStupidName()
		gf.Accuracy = r1.Float32() * 100
		gf.Confidence = r1.Float32() * 100
		gf.Speed = r1.Float32() * 100
		gs = append(gs, &gf)
	}
	data, err := json.Marshal(gs)
	if err != nil {
		return "", err
	}
	fname := getFileName(8) + ".json"
	errJ := ioutil.WriteFile(fmt.Sprintf("fighters/%v", fname), data, 0644)
	if errJ != nil {
		return "", errJ
	}
	return "", nil
}

func getFileName(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
