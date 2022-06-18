package gopnet

// func ToQuery(model interface{}) (string, error) {
// 	var params map[string]interface{}
// 	q := url.Values{}

// 	bytes, err := json.Marshal(model)
// 	if err != nil {
// 		return "", err
// 	}

// 	json.Unmarshal(bytes, &params)

// 	for field, val := range params {
// 		value := fmt.Sprint(val)

// 		if len(value) > 0 {
// 			q.Add(field, value)
// 		}
// 	}

// 	return EncodeQuery(q), nil
// }

// func EncodeQuery(v url.Values) string {
// 	if v == nil {
// 		return ""
// 	}
// 	var buf strings.Builder
// 	keys := make([]string, 0, len(v))
// 	for k := range v {
// 		keys = append(keys, k)
// 	}

// 	for _, k := range keys {
// 		vs := v[k]
// 		keyEscaped := url.QueryEscape(k)
// 		for _, v := range vs {
// 			if buf.Len() > 0 {
// 				buf.WriteByte('&')
// 			}
// 			buf.WriteString(keyEscaped)
// 			buf.WriteByte('=')
// 			buf.WriteString(url.QueryEscape(v))
// 		}
// 	}
// 	return buf.String()
// }
