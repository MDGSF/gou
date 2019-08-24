package ini

import (
	"fmt"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	conf, err := NewConfigFromData([]byte{})
	if err != nil {
		t.Fatal(err)
	}
	conf.Set("runmode", "dev")
	v := conf.String("runmode")
	if v != "dev" {
		t.Fatal(v)
	}

	key := "serveraddrs"
	value := "http://127.0.0.1:8080; ws://10.0.0.1:12306"
	conf.Set(key, value)
	addrs := conf.Strings(key)
	if len(addrs) != 2 {
		t.Fatal(addrs)
	}
}

func TestIni(t *testing.T) {

	var (
		inicontext = `
;comment one
#comment two
appname = beeapi
httpport = 8080
mysqlport = 3600
PI = 3.1415976
runmode = dev
autorender = false
copyrequestbody = true
session= on
cookieon= off
newreg = OFF
needlogin = ON
enableSession = Y
enableCookie = N
flag = 1
[demo]
key1="asta"
key2 = "xie"
CaseInsensitive = true
peers = one;two;three
`

		keyValue = map[string]interface{}{
			"appname":               "beeapi",
			"httpport":              8080,
			"mysqlport":             int64(3600),
			"PI":                    3.1415976,
			"runmode":               "dev",
			"autorender":            false,
			"copyrequestbody":       true,
			"session":               true,
			"cookieon":              false,
			"newreg":                false,
			"needlogin":             true,
			"enableSession":         true,
			"enableCookie":          false,
			"flag":                  true,
			"demo::key1":            "\"asta\"",
			"demo::key2":            "\"xie\"",
			"demo::CaseInsensitive": true,
			"demo::peers":           []string{"one", "two", "three"},
			"null":                  "",
			"demo2::key1":           "",
			"error":                 "",
			"emptystrings":          []string{},
		}
	)

	f, err := os.Create("testini.conf")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(inicontext)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove("testini.conf")

	iniconf, err := NewConfig("testini.conf")
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range keyValue {
		var err error
		var value interface{}
		switch v.(type) {
		case int:
			value, err = iniconf.Int(k)
		case int64:
			value, err = iniconf.Int64(k)
		case float64:
			value, err = iniconf.Float(k)
		case bool:
			value, err = iniconf.Bool(k)
		case []string:
			value = iniconf.Strings(k)
		case string:
			value = iniconf.String(k)
		default:
			value, err = iniconf.DIY(k)
		}
		if err != nil {
			t.Fatalf("get key %q value fail,err %s", k, err)
		} else if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", value) {
			t.Fatalf("get key %q value, want %v got %v .", k, v, value)
		}

	}
	if err = iniconf.Set("name", "astaxie"); err != nil {
		t.Fatal(err)
	}
	if iniconf.String("name") != "astaxie" {
		t.Fatal("get name error")
	}

}
