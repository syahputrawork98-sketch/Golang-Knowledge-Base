package oncefunc

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnceFunc(t *testing.T) {
	// Fitur Go 1.21+: sync.OnceFunc
	// Membungkus logika log agar hanya muncul sekali meskipun dipanggil berkali-kali
	
	count := 0
	logOnce := sync.OnceFunc(func() {
		count++
		fmt.Println("CRITICAL: System initialized (Only shows once)")
	})

	// Panggil 10 kali
	for i := 0; i < 10; i++ {
		logOnce()
	}

	if count != 1 {
		t.Errorf("Expected count 1, got %d", count)
	}
}

func TestOnceValue(t *testing.T) {
	// sync.OnceValue untuk inisialisasi yang mengembalikan nilai
	loadConfig := sync.OnceValue(func() string {
		fmt.Println("Reading config file...")
		return "API_KEY=SECRET_123"
	})

	val1 := loadConfig()
	val2 := loadConfig()

	if val1 != val2 {
		t.Error("Values should be identical")
	}
	t.Logf("Config value: %s", val1)
}
