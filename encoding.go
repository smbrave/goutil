package goutil

import (
	"encoding/json"
	"encoding/xml"
)

// EncodeJSON encode anything to json string
func EncodeJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// EncodeJSON encode anything to json string
func EncodeJSONIndent(v interface{}) string {
	b, _ := json.MarshalIndent(v, "    ", "    ")
	return string(b)
}

// EncodeJSON encode anything to json string
func EncodeXML(v interface{}) string {
	b, _ := xml.Marshal(v)
	return string(b)
}

// EncodeJSON encode anything to json string
func EncodeXMLIndent(v interface{}) string {
	b, _ := xml.MarshalIndent(v, "    ", "    ")
	return string(b)
}
