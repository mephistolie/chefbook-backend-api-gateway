package response

func NonNilStringMap(response map[string]string) map[string]string {
	if response != nil {
		return response
	}
	return make(map[string]string)
}
