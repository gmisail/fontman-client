package api

/*
	Get all fonts defined in the registry.
*/
func GetAllFonts() []string { // TODO: change to a Font data model
	return []string{"arial", "comic sans", "impact"}
}

/*
	Get a specific font by name.
*/
func GetFontByName(name string) string { // TODO: change to a Font data model
	return name
}

/*
	Get a specific font by ID.
*/
func GetFontById(id string) string { // TODO: change to a Font data model
	return id
}
