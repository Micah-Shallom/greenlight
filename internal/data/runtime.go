package data

import (
	"fmt"
	"strconv"
)

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