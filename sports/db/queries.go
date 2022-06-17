package db

const (
	eventsList = "list"
)

func getEventQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				name, 
				sport, 
				visible, 
				advertised_start_time 
			FROM events
		`,
	}
}
