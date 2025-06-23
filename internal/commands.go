package internal

type Command struct {
	Name        string
	Description string
	KeyBindings []string
}

func LoadCommands() []Command {
	return []Command{
		{
			Name:        "new",
			Description: "Create a new note",
			KeyBindings: []string{"n"},
		},
		{
			Name:        "list",
			Description: "List all notes",
			KeyBindings: []string{"l"},
		},
		{
			Name:        "search",
			Description: "Finding notes by title",
			KeyBindings: []string{"s"},
		},
	}
}
