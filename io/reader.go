/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * io/reader.go                                           *
 *                                                        *
 * hprose reader for Go.                                  *
 *                                                        *
 * LastModified: Aug 16, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package io

// Reader is a fine-grained operation struct for Hprose unserialization
// when JSONCompatible is true, the Map data will unserialize to map[string]interface as the default type
type Reader struct {
	classref       []interface{}
	fieldsref      [][]string
	JSONCompatible bool
}

// Unmarshaler is a interface for unserializing user custum type
type Unmarshaler interface {
	UnmarshalHprose(reader *Reader, tag byte) error
}
