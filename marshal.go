package querystring

import (
	"fmt"
	"reflect"
	// "strconv"
	"net/url"
	"strings"
	"time"
)

//Marshal  nested struct or map to url query strings
func Marshal(obj interface{}) (string, error) {
	values := make(url.Values)
	err := encode(obj, "", values)
	return values.Encode(), err
}

//Stringify
func Stringify(obj interface{}) (string, error) {
	return Marshal(obj)
}

func encode(val interface{}, prefix string, values url.Values) error {
	v := reflect.ValueOf(val)
	// v := reflect.TypeOf(val)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	switch v.Kind() {

	case reflect.Struct:

		err := encodeStruct(v, prefix, values)
		if err != nil {
			return err
		}

	case reflect.Map:
		keys := v.MapKeys()
		for _, k := range keys {
			itemValue := v.MapIndex(k)

			key := k.Interface().(string)
			if len(prefix) > 0 {
				key = fmt.Sprintf("[%s]", k.Interface().(string))
			}

			err := encode(itemValue.Interface(), prefix+key, values)
			// fmt.Printf("map str %+v %+v", values, err)

			if err != nil {
				return err
			}

		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {

			sliceValue := v.Index(i)

			if sliceValue.Kind() == reflect.Ptr {
				sliceValue = sliceValue.Elem()
			}
			key := fmt.Sprintf("[%d]", i)

			err := encode(sliceValue.Interface(), prefix+key, values)
			// fmt.Printf("slice str %s %+v", str, err)

			if err != nil {
				return err
			}

		}
	default:
		if len(prefix) > 0 {
			values.Add(prefix, valueToString(v))
		}

	}
	return nil
}

func parseFieldTag(tag string) (string, []string) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}
func encodeStruct(v reflect.Value, prefix string, values url.Values) error {

	// fmt.Printf("encodeStruct...prefix:%s\n", prefix)
	// components := ""
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tf := typ.Field(i)
		sv := v.Field(i)

		tag, _ := tf.Tag.Lookup("url")

		if tag == "-" {
			continue
		}

		name := strings.ToLower(tf.Name)

		if len(tag) >= 0 {
			tname, _ := parseFieldTag(tag)
			if len(tname) > 0 {
				name = tname
			}
		}

		key := name
		if len(prefix) > 0 {
			key = fmt.Sprintf("[%s]", name)
		}

		// fmt.Printf("tf.Name:%s\n", name)
		switch sv.Interface().(type) {

		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			bool, string, time.Time:

			values.Add(prefix+key, valueToString(sv))

		case *int, *int8, *int16, *int32, *int64,
			*uint, *uint8, *uint16, *uint32, *uint64,
			*float32, *float64,
			*bool, *string, *time.Time:

			values.Add(prefix+key, valueToString(sv))

		default:
			err := encode(sv.Interface(), prefix+key, values)
			if err != nil {
				return err
			}
		}

	}
	return nil

}

func valueToString(v reflect.Value) string {

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()

	}

	if v.Type() == reflect.TypeOf(time.Time{}) {
		t := v.Interface().(time.Time)

		return t.Format(time.RFC3339)

	}

	return fmt.Sprint(v.Interface())
}
