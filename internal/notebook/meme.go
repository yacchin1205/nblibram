package notebook

func GetMemeID(c Cell) string {
	memeMap := getMemeMap(c)
	if memeMap == nil {
		return ""
	}
	current, ok := memeMap["current"].(string)
	if !ok {
		return ""
	}
	return current
}

func GetMemePrevious(c Cell) string {
	memeMap := getMemeMap(c)
	if memeMap == nil {
		return ""
	}
	prev, ok := memeMap["previous"].(string)
	if !ok {
		return ""
	}
	return prev
}

func GetMemeNext(c Cell) string {
	memeMap := getMemeMap(c)
	if memeMap == nil {
		return ""
	}
	next, ok := memeMap["next"].(string)
	if !ok {
		return ""
	}
	return next
}

func getMemeMap(c Cell) map[string]interface{} {
	meme, ok := c.Metadata["lc_cell_meme"]
	if !ok {
		return nil
	}
	memeMap, ok := meme.(map[string]interface{})
	if !ok {
		return nil
	}
	return memeMap
}
