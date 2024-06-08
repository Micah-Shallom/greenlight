package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

//declare custom runtime type which has underlying type int32
type Runtime int32

// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. This should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONvalue := strconv.Quote(jsonValue)

	return []byte(quotedJSONvalue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	//we expect that the incoming json value will be a string in the format "<runtime> mins"
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	//split the string to isolate the part containig the number
	parts := strings.Split(unquotedJSONValue, " ")

	//sanity check
	if len(parts) != 2 || parts[1] != "mins"{
		return ErrInvalidRuntimeFormat
	}

	//otherwise parse the string containing the number into an int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	//Convert the int32 to a runtime type and assign this to the receiver
	*r = Runtime(i)

	return nil
}