package ip2country

import (
	"testing"
)

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
