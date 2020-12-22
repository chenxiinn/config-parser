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

package actions

import (
	"fmt"

	"github.com/haproxytech/config-parser/v3/common"
)

type SetPriorityClass struct {
	Expr common.Expression
}

func (f *SetPriorityClass) Parse(parts []string) error {
	expr := common.Expression{}

	err := expr.Parse([]string{parts[1]})
	if err != nil {
		return fmt.Errorf("not enough params")
	}

	f.Expr = expr

	return nil
}

func (f *SetPriorityClass) String() string {
	return fmt.Sprintf("set-priority-class %s", f.Expr.String())
}
