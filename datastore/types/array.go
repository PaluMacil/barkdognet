package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type TextArray []string

func (ta *TextArray) Scan(value interface{}) error {
	if value == nil {
		*ta = []string{}
		return nil
	}
	var strVal string
	switch v := value.(type) {
	case []byte:
		strVal = processByteArray(v)
	case string:
		strVal = processStringVal(v)
	default:
		return fmt.Errorf("unsupported type for TextArray.Scan: %T", v)
	}

	if strVal == "" {
		*ta = []string{} // Ensures empty arrays are empty slices
	} else {
		*ta = strings.Split(strVal, ",")
	}
	return nil
}

func (ta TextArray) Value() (driver.Value, error) {
	if ta == nil || len(ta) == 0 {
		return "{}", nil
	}
	quoted := make([]string, len(ta))
	for i, s := range ta {
		quoted[i] = fmt.Sprintf("\"%s\"", s)
	}
	return "{" + strings.Join(quoted, ",") + "}", nil
}

type IntArray []int32

func (ia *IntArray) Scan(value interface{}) error {
	if value == nil {
		*ia = []int32{}
		return nil
	}
	var strVal string
	switch v := value.(type) {
	case string:
		strVal = processStringVal(v)
	case []uint8:
		strVal = processByteArray(v)
	default:
		return fmt.Errorf("unsupported type for IntArray.Scan: %T", v)
	}

	strVal = strings.NewReplacer("{", "", "}", "", `"`, "", " ", "").Replace(strVal)
	if strVal == "" {
		*ia = []int32{} // Ensures empty arrays are empty slices
	} else {
		strArray := strings.Split(strVal, ",")
		nums := make([]int32, 0, len(strArray))
		for _, s := range strArray {
			num, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return fmt.Errorf("invalid value scanned as IntArray (%s): %w", s, err)
			}
			nums = append(nums, int32(num))
		}
		*ia = nums
	}
	return nil
}

func (ia IntArray) Value() (driver.Value, error) {
	if ia == nil {
		return "{}", nil
	}

	sb := strings.Builder{}
	sb.WriteString("{")
	for i, val := range ia {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%d", val))
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func processStringVal(val string) string {
	if len(val) < 2 {
		return ""
	}
	return strings.TrimSpace(val[1 : len(val)-1])
}

func processByteArray(val []byte) string {
	return processStringVal(string(val))
}
