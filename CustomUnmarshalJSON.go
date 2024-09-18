
type response struct {
	Album     string `json:"album"`
	Title     string
	Year      string
	Downloads string
}

type responseWrapper struct {
	response
}

func (r *responseWrapper) UnmarshalJSON(b []byte) (err error) {
	// Parse the raw data into a map for custom handling
	var raw map[string]interface{}

	// First unmarshal into r.response (assuming it's for storing the raw response)
	if err := json.Unmarshal(b, &r.response); err != nil {
		return err // Return the error if unmarshalling fails
	}

	// Now unmarshal into the raw map
	if err := json.Unmarshal(b, &raw); err != nil {
		return err // Return the error if unmarshalling fails
	}

	fmt.Println("Hello")
	// Switch based on r.Item to parse different structures
	switch r.Album {
	case "song":
		if inner, ok := raw["song"].(map[string]interface{}); ok {
			if title, ok := inner["title"].(string); ok {
				r.Title = title
			}
		}
		if inner, ok := raw["song"].(map[string]interface{}); ok {
			if year, ok := inner["year"].(string); ok {
				r.Year = year
			}
		}
		if inner, ok := raw["song"].(map[string]interface{}); ok {
			if dw, ok := inner["downloads"].(string); ok {
				r.Downloads = dw
			}
		}
	}

	return nil // Return nil if everything is successful
}

func main() {

	var resp1 responseWrapper
	var err error

	data := `{
		"album":"song",
		"song":{ "title":"Happy trails","year":"2232", "downloads":"21222222"}
	 }`
	if err = json.Unmarshal([]byte(data), &resp1); err != nil {

		fmt.Printf("%#v\n", err)

	}
	fmt.Printf("%#v\n", resp1.response)

}
