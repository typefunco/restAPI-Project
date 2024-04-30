package events

type Events struct {
	ID             int
	TotalPeople    int    `binding:"required"`
	Theme          string `binding:"required"`
	MinuteDuration int
}

var eventsList = []Events{
	{ID: 1,
		TotalPeople:    20,
		Theme:          "Crypto",
		MinuteDuration: 90,
	},
	{
		ID:             2,
		TotalPeople:    30,
		Theme:          "IT",
		MinuteDuration: 240,
	},
}

func GetEvents() []Events {
	return eventsList
}

func (e Events) Save() {
	eventsList = append(eventsList, e)
}
