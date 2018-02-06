package ip2country

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkGetCountry(b *testing.B) {

	err := Load("./dbip-country.csv")
	if err != nil {
		b.Fatal(err)
	}

	for i := 1; i <= b.N; i++ {
		GetCountry(createRandomeIP())
	}

}

func createRandomeIP() string {
	p1 := rand.Int31n(256)
	p2 := rand.Int31n(256)
	p3 := rand.Int31n(256)
	p4 := rand.Int31n(256)

	return fmt.Sprintf("%d.%d.%d.%d", p1, p2, p3, p4)
}

func TestIPtoCountryLookup(t *testing.T) {

	err := Load("./dbip-country.csv")
	if err != nil {
		t.Fatal(err)
	}

	country := GetCountry("0.0.0.0")
	if country != "ZZ" {
		t.Errorf("Expected ZZ but found %s", country)
	}
	country = GetCountry("0.0.0.1")
	if country != "ZZ" {
		t.Errorf("Expected ZZ but found %s", country)
	}
	country = GetCountry("0.255.255.255")
	if country != "ZZ" {
		t.Errorf("Expected ZZ but found %s", country)
	}

	country = GetCountry("50.97.196.208")
	if country != "US" {
		t.Errorf("Expected US but found %s", country)
	}

	country = GetCountry("50.97.198.135")
	if country != "CN" {
		t.Errorf("Expected US but found %s", country)
	}

	country = GetCountry("50.97.198.135")
	if country != "CN" {
		t.Errorf("Expected US but found %s", country)
	}

	country = GetCountry("200.95.185.34")
	if country != "CL" {
		t.Errorf("Expected CL but found %s", country)
	}

	country = GetCountry("223.255.255.255")
	if country != "AU" {
		t.Errorf("Expected AU but found %s", country)
	}
}

func TestIPAddressToInt(t *testing.T) {
	ipNumb, err := ipToInt("0.0.0.255")
	if err != nil {
		t.Error(err)
	}
	if ipNumb != 255 {
		t.Errorf("Expected 255 but found %d", ipNumb)
	}
	//-----------------------------------------------
	ipNumb, err = ipToInt("0.0.1.255")
	if err != nil {
		t.Error(err)
	}
	if ipNumb != 511 {
		t.Errorf("Expected 511 but found %d", ipNumb)
	}
	//-----------------------------------------------
	ipNumb, err = ipToInt("255.0.0.0")
	if err != nil {
		t.Error(err)
	}
	if ipNumb != 4278190080 {
		t.Errorf("Expected 4278190080 but found %d", ipNumb)
	}
	//-----------------------------------------------
	ipNumb, err = ipToInt("255.255.255.255")
	if err != nil {
		t.Error(err)
	}
	if ipNumb != 4294967295 {
		t.Errorf("Expected 4294967295 but found %d", ipNumb)
	}
}
