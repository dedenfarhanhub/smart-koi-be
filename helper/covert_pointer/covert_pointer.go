package covert_pointer

import "strconv"

func ConvertPointerString(x *string) *string {
	if *x == "" {
		x = nil
	}
	return x
}

func ConvertNilPointerString(x *string) string {
	if x == nil {
		return ""
	}
	return *x
}

func ConvertStringInt(x string) int  {
	y, err := strconv.Atoi(x)

	if err != nil {
		y=-1
	}
	return y
}