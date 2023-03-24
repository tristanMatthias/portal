package huggingface

func (m *HuggingFace) ActionSearchModels(search string) ([]HFModel, error) {
	println("ActionSearchModels")
	println(search)
	return GetTransformers(APIQueryParams{
		Limit: 	 10,
		Search:  search,
		Sort:    "downloads",
		Direction: -1,
	})
}
