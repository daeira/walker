package merger

import "fmt"

// Merge merges src map recursively into dst map.
// TODO(all): throw error if we try to merge slice into map?
func Merge(dst map[interface{}]interface{}, src map[interface{}]interface{}) {
	for k, v := range src {
		if srcMap, ok := v.(map[interface{}]interface{}); ok {
			if dstMap, ok := dst[k].(map[interface{}]interface{}); ok {
				Merge(dstMap, srcMap)
				continue
			}
		}
		dst[k] = v
	}
}

func Origin(in interface{}, origin string) {
	if m, ok := in.(map[interface{}]interface{}); ok {
		for k, v := range m {
			if emptyValue, ok := mapOrSlice(v, origin); ok {
				if len(emptyValue) > 0 {
					m[k] = emptyValue
					continue
				}
				Origin(v, origin)
				continue
			}
			m[k] = fmt.Sprintf("%v [%s]", v, origin)
		}
	}
	if s, ok := in.([]interface{}); ok {
		for i, v := range s {
			if emptyValue, ok := mapOrSlice(v, origin); ok {
				if len(emptyValue) > 0 {
					s[i] = emptyValue
					continue
				}
				Origin(v, origin)
				continue
			}
			s[i] = fmt.Sprintf("%v [%s]", v, origin)
		}
	}
}

func mapOrSlice(item interface{}, origin string) (string, bool) {
	isMapOrSlice := false
	if childMap, ok := item.(map[interface{}]interface{}); ok {
		isMapOrSlice = true
		if len(childMap) == 0 {
			return fmt.Sprintf("{} [%s]", origin), isMapOrSlice
		}
	}
	if childSlice, ok := item.([]interface{}); ok {
		isMapOrSlice = true
		if len(childSlice) == 0 {
			return fmt.Sprintf("[] [%s]", origin), isMapOrSlice
		}
	}
	return "", isMapOrSlice
}
