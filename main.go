package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Bearer Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzk1ODUyNzMsImlhdCI6MTY3OTU4NTI3MywiaXNzIjoiY29vbGNhciIsInN1YiI6IjEifQ.GGSeGRFzYFx8hlnvDtgLqnbgMSJDplIUlFUvRGAMBCpgAaBUxudjpZpOUcBfTNufkVLGnfuMNG1ONPa1BXBf4JyqLfvIed6xa1uHeLxDZQUh0hm7uyjqhOndtlskU0GxCqwkq-RGHLovvUY8yRW1HKug4AvMK5uz_03-Or24Jh1aW_ODRVK3SVY_2-FVZtvWZlz0jOqVE9dICArx8Wib7rM0wcqGQHf3juEZW8hGGENLX01DPcNQof9PZ6v7x1JvJDagJz7QmuR54iit_YG1BfYmuTudLet-GFPjQ_YZ_4agEEqAU7a9qcq1HXeFs7lnrYVpkvYXcNbdAdknIUvF9A"

	v := "Bearer "

	fmt.Printf("here:%s", strings.Trim(s, v))
	// strs := strings.Split(s, " ")
	// fmt.Println(strs[len(strs)-1])
}
