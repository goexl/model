package model

import (
	"encoding/json"
	"encoding/xml"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

// Slice 切片，既可以兼容单个值也可以兼容数组
type Slice[T any] []*T

func (s *Slice[T]) UnmarshalJSON(bytes []byte) (err error) {
	t := new(T)
	if ue := json.Unmarshal(bytes, t); nil != ue {
		err = json.Unmarshal(bytes, s)
	} else {
		*s = []*T{t}
	}

	return
}

func (s *Slice[T]) UnmarshalYAML(node *yaml.Node) (err error) {
	t := new(T)
	if ue := node.Decode(t); nil != ue {
		err = node.Decode(s)
	} else {
		*s = []*T{t}
	}

	return
}

func (s *Slice[T]) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	t := new(T)
	if ue := decoder.DecodeElement(t, &start); nil != ue {
		err = decoder.DecodeElement(s, &start)
	} else {
		*s = []*T{t}
	}

	return
}

func (s *Slice[T]) UnmarshalText(bytes []byte) (err error) {
	t := new(T)
	if ue := toml.Unmarshal(bytes, t); nil != ue {
		err = toml.Unmarshal(bytes, s)
	} else {
		*s = []*T{t}
	}

	return
}
