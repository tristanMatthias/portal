package settings

import "portal/server/database"

func Setup() ESettings {
	var result []ESettings
	// Query the database to see if the settings table is empty
	res := database.DB.Find(&result)

	if res.Error != nil {
		panic(res.Error)
	}

	// If the settings table is empty, create the default settings
	if len(result) == 0 {
		settings := ESettings{
			Setup: true,
		}
		res := database.DB.Create(&settings)
		if res.Error != nil {
			panic(res.Error)
		}
		return settings
	}

	return result[0]
}
