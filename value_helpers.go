// Copyright 2013-2016 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

// MapIter allows to define general maps of your own type to be used in the Go client
// without the use of reflection.
type MapIter interface {
	Range(func(interface{}, interface{}) error) error
	Len() int
}

// ListIter allows to define general maps of your own type to be used in the Go client
// without the use of reflection.
type ListIter interface {
	Range(func(interface{}) error) error
	Len() int
}

// supports old value arrays
type ifcValueList []interface{}

func (m ifcValueList) Range(f func(interface{}) error) error {
	for idx := range []interface{}(m) {
		if err := f(NewValue(m[idx])); err != nil {
			return err
		}
	}

	return nil
}

func (m ifcValueList) Len() int {
	return len(m)
}

// supports old value arrays
type valueList []Value

func (m valueList) Range(f func(interface{}) error) error {
	for idx := range []Value(m) {
		if err := f(m[idx]); err != nil {
			return err
		}
	}

	return nil
}

func (m valueList) Len() int {
	return len(m)
}
