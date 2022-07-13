package banchmarks

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func BenchmarkGetPostsHandler(b *testing.B) {
	b.Run("Endpoint: GET ", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			get, err := http.Get("https://...")
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			if get.StatusCode != http.StatusOK {
				fmt.Printf("not ok status code %v", get.StatusCode)
				os.Exit(1)
			}
		}
	})
}
