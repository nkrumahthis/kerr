package core

type Principle struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func GetPrinciples() []Principle {
	return []Principle{
		{
			Description: "Communicate proactively: Overcommunicate with stakeholders.",
			Tags:        []string{"communication", "tasks", "stakeholders"},
		},
		{
			Description: "Be early: Deliver work before the deadline.",
			Tags:        []string{"tasks", "time-management"},
		},
		{
			Description: "Double-check your work: Ensure it's error-free.",
			Tags:        []string{"tasks", "quality"},
		},
		{
			Description: "Take notes: Record meeting outcomes and action items.",
			Tags:        []string{"meetings", "notes"},
		},
		{
			Description: "Ask questions: Clarify doubts immediately.",
			Tags:        []string{"meetings", "questions"},
		},
		{
			Description: "Follow up: Send recap emails to ensure alignment.",
			Tags:        []string{"communication", "tasks"},
		},
		{
			Description: "Own your mistakes: Fix issues and learn from them.",
			Tags:        []string{"responsibility", "growth"},
		},
		{
			Description: "Be a team player: Help others succeed.",
			Tags:        []string{"teamwork", "collaboration"},
		},
		{
			Description: "Keep learning: Improve your skills continuously.",
			Tags:        []string{"growth", "learning"},
		},
		{
			Description: "Show gratitude: Thank colleagues for their contributions.",
			Tags:        []string{"teamwork", "gratitude"},
		},
	}
}
