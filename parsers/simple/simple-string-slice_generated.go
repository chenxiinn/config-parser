// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package simple

import (
	"github.com/haproxytech/config-parser/v4/common"
	"github.com/haproxytech/config-parser/v4/errors"
	"github.com/haproxytech/config-parser/v4/types"
)

func (p *StringSlice) Init() {
	p.data = nil
	p.preComments = []string{}
}

func (p *StringSlice) GetParserName() string {
	return p.Name
}

func (p *StringSlice) Get(createIfNotExist bool) (common.ParserData, error) {
	if p.data == nil {
		if createIfNotExist {
			p.data = &types.StringSliceC{}
			return p.data, nil
		}
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *StringSlice) GetPreComments() ([]string, error) {
	return p.preComments, nil
}

func (p *StringSlice) SetPreComments(preComments []string) {
	p.preComments = preComments
}

func (p *StringSlice) GetOne(index int) (common.ParserData, error) {
	if index > 0 {
		return nil, errors.ErrFetch
	}
	if p.data == nil {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *StringSlice) Delete(index int) error {
	p.Init()
	return nil
}

func (p *StringSlice) Insert(data common.ParserData, index int) error {
	return p.Set(data, index)
}

func (p *StringSlice) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case *types.StringSliceC:
		p.data = newValue
	case types.StringSliceC:
		p.data = &newValue
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *StringSlice) PreParse(line string, parts []string, preComments []string, comment string) (changeState string, err error) {
	changeState, err = p.Parse(line, parts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

func (p *StringSlice) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
