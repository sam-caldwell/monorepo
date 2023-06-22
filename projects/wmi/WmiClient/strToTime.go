package wmiclient

/*
 * strToTime.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This Go function, strToTime, is designed to parse a string representing a time
 * and assign it to a provided reflect.Value of type time.Time.
 */

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// strToTime -  parse a string representing a time and assign it to a provided reflect.Value of type time.Time.
func strToTime(field reflect.Value, val string) (err error) {
	if field.Type().Name() == "Time" {
		if len(val) == 25 {
			var minutes int
			if minutes, err = strconv.Atoi(val[22:]); err != nil {
				return err //malformed time string or other error
			}
			val = val[:22] + fmt.Sprintf("%02d%02d", minutes/60, minutes%60)
		}
		var thisTime time.Time
		if thisTime, err = time.Parse("20060102150405.000000-0700", val); err != nil {
			return err //error parsing time.
		}
		field.Set(reflect.ValueOf(thisTime))
		return nil
	}
	return fmt.Errorf("unsupported string to struct")
}
