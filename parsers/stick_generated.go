// Code generated by go generate; DO NOT EDIT.
package parsers

import (
	"github.com/haproxytech/config-parser/common"
	"github.com/haproxytech/config-parser/errors"
	"github.com/haproxytech/config-parser/types"
)

func (p *Stick) Init() {
	p.data = []types.Stick{}
}

func (p *Stick) GetParserName() string {
	return "stick"
}

func (p *Stick) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.FetchError
	}
	return p.data, nil
}

func (p *Stick) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.FetchError
	}
	return p.data[index], nil
}

func (p *Stick) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.FetchError
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.Stick{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *Stick) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.InvalidData
	}
	switch newValue := data.(type) {
	case []types.Stick:
		p.data = newValue
	case *types.Stick:
		if index > -1 {
			if index > len(p.data) {
				return errors.IndexOutOfRange
			}
			p.data = append(p.data, types.Stick{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.Stick:
		if index > -1 {
			if index > len(p.data) {
				return errors.IndexOutOfRange
			}
			p.data = append(p.data, types.Stick{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.InvalidData
	}
	return nil
}

func (p *Stick) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.Stick:
		p.data = newValue
	case *types.Stick:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.IndexOutOfRange
		}
	case types.Stick:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else if index == -1 {
			p.data = append(p.data, newValue)
		} else {
			return errors.IndexOutOfRange
		}
	default:
		return errors.InvalidData
	}
	return nil
}