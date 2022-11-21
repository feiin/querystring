package querystring

import (
	"net/url"
	"testing"
)

func TestEncode(t *testing.T) {

	type Item struct {
		ItemName  string `json:"item_name" url:"item_name"`
		ItemValue string `json:"item_value" url:"item_value"`
	}
	type Person struct {
		Name     string `json:"name"`
		NickName string `json:"nick_ame" url:"nick_name"`
		Items    []Item `json:"items" url:"items"`
	}

	p := Person{
		Name:     "Json",
		NickName: "J",
	}

	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})
	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})

	myMap := make(map[string]interface{})
	myMap["test1"] = "asdfasfd"
	myMap["test2"] = 2
	myMap["test3"] = "my_test"

	vs := make(url.Values)
	encode(p, "", vs)
	str := vs.Encode()
	t.Logf("result===:%v", str)

	vs = make(url.Values)

	encode(myMap, "", vs)
	t.Logf("result2:%v", vs.Encode())

}

func TestEncode2(t *testing.T) {

	type Item struct {
		ItemName  string `json:"item_name" url:"item_name"`
		ItemValue string `json:"item_value" url:"item_value"`
	}
	type Person struct {
		Name     string                 `json:"name"`
		NickName string                 `json:"nick_ame" url:"nick_name"`
		Items    []Item                 `json:"items" url:"items"`
		MItems   map[string]interface{} `json:"m_items" url:"m_items"`
	}

	p := Person{
		Name:     "Jason",
		NickName: "J",
	}

	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})
	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})

	myMap := make(map[string]interface{})
	myMap["test1"] = "asdfasfd"
	myMap["test2"] = 2
	myMap["test3"] = "my_test"

	p.MItems = myMap

	vs := make(url.Values)

	encode(p, "", vs)

	t.Logf("result:%+v", vs.Encode())

}

func TestMarshal1(t *testing.T) {
	type Item struct {
		ItemName  string `json:"item_name" url:"item_name"`
		ItemValue string `json:"item_value" url:"item_value"`
	}
	type Person struct {
		Name     string                 `json:"name"`
		NickName string                 `json:"nick_ame" url:"nick_name"`
		Items    []Item                 `json:"items" url:"items"`
		MItems   map[string]interface{} `json:"m_items" url:"m_items"`
	}

	p := Person{
		Name:     "Jason",
		NickName: "J",
	}

	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})
	p.Items = append(p.Items, Item{ItemName: "hello", ItemValue: "world"})

	myMap := make(map[string]interface{})
	myMap["test1"] = "asdfasfd"
	myMap["test2"] = 2
	myMap["test3"] = "my_test"

	p.MItems = myMap

	result, err := Marshal(p)
	t.Logf("result:%+v ,error:%+v", result, err)

}

func TestMarshal2(t *testing.T) {

	type Person struct {
		Name     string `json:"name"`
		NickName string `json:"nick_name" url:"nick_name"`
	}

	p := Person{
		Name:     "Json",
		NickName: "J",
	}

	result, err := Marshal(p)
	if result != "name=Json&nick_name=J" {
		t.Errorf("marshal error")
	}
	t.Logf("result:%+v ,error:%+v", result, err)

	result, err = Marshal(123)
	if result != "" {
		t.Errorf("marshal error")
	}
	t.Logf("result:%+v ,error:%+v", result, err)

	type Test struct {
		A string `url:"a"`
	}

	te := Test{
		A: "",
	}
	result, err = Marshal(te)
	if result != "a=%EE%80%80" {
		t.Errorf("marshal t error")
	}
	t.Logf("marshal t result :%+v ,error:%+v", result, err)

}
