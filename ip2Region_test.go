package ip2region

import (
	"testing"
)

func TestIp2Region_BinarySearch(t *testing.T) {
	region, err := New("ip2region.db")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 1000; i++ {
		ip := RandomIp()
		ipInfo, err := region.BtreeSearch(ip)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ip, ipInfo.City)
		}
	}
}
func BenchmarkBtreeSearch(B *testing.B) {
	region, err := New("ip2region.db")
	if err != nil {
		B.Error(err)
	}
	for i := 0; i < B.N; i++ {
		region.BtreeSearch("127.0.0.1")
	}
}

func BenchmarkMemorySearch(B *testing.B) {
	region, err := New("ip2region.db")
	if err != nil {
		B.Error(err)
	}
	for i := 0; i < B.N; i++ {
		region.MemorySearch("127.0.0.1")
	}

}

func BenchmarkBinarySearch(B *testing.B) {
	region, err := New("ip2region.db")
	if err != nil {
		B.Error(err)
	}
	for i := 0; i < B.N; i++ {
		region.BinarySearch("127.0.0.1")
	}

}

func TestIp2long(t *testing.T) {
	ip, err := ip2long("127.0.0.1")
	if err != nil {
		t.Error(err)
	}
	if ip != 2130706433 {
		t.Error("result error")
	}
	t.Log(ip)
}
