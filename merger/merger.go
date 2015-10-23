package merger

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
