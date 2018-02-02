package ip2country

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIPtoCountryLookup(t *testing.T) {
	file, err := ioutil.TempFile("", "ip2country")
	if err != nil {
		t.Fatal(err)
	}
	file.Write([]byte(`"0.0.0.0","0.0.0.255","AA"`))
	file.Write([]byte("\n"))
	file.Write([]byte(`"0.0.32.0","0.0.64.255","BB"`))
	file.Write([]byte("\n"))
	file.Write([]byte(`"0.128.0.0","0.128.255.255","CC"`))
	file.Write([]byte("\n"))
	file.Write([]byte(`"80.0.0.0","90.0.0.0","DD"`))
	file.Write([]byte("\n"))
	file.Write([]byte(`"91.0.0.0","255.255.255.255","EE"`))
	err = file.Close()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	err = Load(file.Name())
	if err != nil {
		t.Fatal(err)
	}

	country := GetCountry("0.0.0.1")
	if country != "AA" {
		t.Errorf("Expected AA but found %s", country)
	}

	country = GetCountry("0.0.0.255")
	if country != "AA" {
		t.Errorf("Expected AA but found %s", country)
	}

	country = GetCountry("0.0.32.0")
	if country != "BB" {
		t.Errorf("Expected BB but found %s", country)
	}

	country = GetCountry("0.0.54.100")
	if country != "BB" {
		t.Errorf("Expected BB but found %s", country)
	}

	country = GetCountry("0.128.1.1")
	if country != "CC" {
		t.Errorf("Expected CC but found %s", country)
	}

	country = GetCountry("1.0.0.0") //not in a database
	if country != "" {
		t.Errorf("Expected NOTHING but found %s", country)
	}

	country = GetCountry("85.128.129.130")
	if country != "DD" {
		t.Errorf("Expected DD but found %s", country)
	}

	country = GetCountry("91.0.0.0")
	if country != "EE" {
		t.Errorf("Expected EE but found %s", country)
	}

	country = GetCountry("92.93.94.95")
	if country != "EE" {
		t.Errorf("Expected EE but found %s", country)
	}

	country = GetCountry("255.255.255.255")
	if country != "EE" {
		t.Errorf("Expected EE but found %s", country)
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
