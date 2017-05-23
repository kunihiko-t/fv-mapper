/*
The MIT License (MIT)
Copyright (c) 2017 kunihiko-t.
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package fvm

import (
	"fmt"
	"github.com/serenize/snaker"
	"net/http"
	"regexp"
	"unicode"
)

//GetMap is a function to fetch all params.
func GetMap(r *http.Request) map[string]string {
	return getMap(r)
}

//GetMapSequential is a function to fetch values from sequential map value with given base variable name.
func GetMapSequential(name string, r *http.Request) map[string]string {
	m := getMap(r)
	s := fmt.Sprintf(`(^%v_[\d]+$)|(^[\d]+_%v$)`, name, name)
	var re = regexp.MustCompile(s)
	for key, _ := range r.Form {
		if !re.MatchString(key) {
			delete(m, key)
		}
	}
	return m
}

func getMap(r *http.Request) map[string]string {
	m := map[string]string{}
	for key, values := range r.Form {
		for _, value := range values {
			if m[key] != "" {
				m[key] += "\t" + value
			} else {
				m[key] = value
			}
		}
	}
	return m
}

//GetCamelMap is a function to fetch all params with camel case.
func GetCamelMap(capitalStart bool, r *http.Request) map[string]string {
	m := getMap(r)
	for key, _ := range r.Form {
		v := m[key]
		delete(m, key)
		k := snaker.SnakeToCamel(key)
		rk := []rune(k)
		if capitalStart {
			rk[0] = unicode.ToUpper(rk[0])
		} else {
			rk[0] = unicode.ToLower(rk[0])
		}
		m[string(rk)] = v
	}
	return m
}

func GetSnakeMap(r *http.Request) map[string]string {
	m := getMap(r)
	for key, _ := range r.Form {
		v := m[key]
		delete(m, key)
		m[snaker.CamelToSnake(key)] = v
	}
	return m
}
