package main

import (
	"bytes"
	"fmt"
	mapsize "github.com/520MianXiangDuiXiang520/MapSize"
	"math/rand"
	"os"
)

func OrderInt64() {
	f, _ := os.OpenFile("./statistics/int64.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	for i := 0; i < 1000; i++ {
		n := i * 100
		m := make(map[int64]int64)
		for j := 0; j < n; j++ {
			m[int64(j)] = int64(i)
		}
		res := mapsize.Size(m)
		t := int64(16 * n)
		f.WriteString(fmt.Sprintf("%d,%d,%d,%d,%f\n", n, res, t, res-t, float64(t)/float64(res)))
	}
}

func RandInt64() {
	f, _ := os.OpenFile("./statistics/int64_rand.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	for i := 0; i < 1000; i++ {
		n := i * 100
		m := make(map[int64]int64)
		for j := 0; j < n; j++ {
			m[rand.Int63()] = int64(i)
		}
		res := mapsize.Size(m)
		t := int64(16 * n)
		f.WriteString(fmt.Sprintf("%d,%d,%d,%d,%f\n", n, res, t, res-t, float64(t)/float64(res)))
	}
}

func RandString() {
	randStr := func() string {
		res := bytes.NewBufferString("")
		for i := 0; i < 50; i++ {
			res.WriteByte(byte('a' + rand.Int31n(26)))
		}
		return res.String()
	}
	f, _ := os.OpenFile("./statistics/string_rand.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	for i := 0; i < 1000; i++ {
		n := i * 100
		m := make(map[string]string)
		for j := 0; j < n; j++ {
			m[randStr()] = "int64(i)"
		}
		res := mapsize.Size(m)
		t := int64(32 * n)
		f.WriteString(fmt.Sprintf("%d,%d,%d,%d,%f\n", n, res, t, res-t, float64(t)/float64(res)))
	}
}

func main() {
	//RandString()
	m := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		m[i] = struct{}{}
	}
	fmt.Println(mapsize.Size(m))
}
