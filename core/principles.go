package core

type Principle struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func GetPrinciples() []Principle {
	return []Principle{
		{
			Description: "Communicate proactively: Overcommunicate with stakeholders.",
			Tags:        []string{"communication", "slack"},
		},
		{
			Description: "Be early: Deliver work before the deadline.",
			Tags:        []string{"time", "early"},
		},
		{
			Description: "Double-check your work: Ensure it's error-free.",
			Tags:        []string{"quality", "perfect"},
		},
		{
			Description: "Take notes: Record meeting outcomes and action items.",
			Tags:        []string{"notes"},
		},
		{
			Description: "Ask questions: Clarify doubts immediately.",
			Tags:        []string{"questions"},
		},
		{
			Description: "Follow up: Send recap emails to ensure alignment.",
			Tags:        []string{"recap", "email", "emails"},
		},
		{
			Description: "Own your mistakes: Fix issues and learn from them.",
			Tags:        []string{"responsibility", "fix", "mistakes", "mybad"},
		},
		{
			Description: "Be a team player: Help others succeed.",
			Tags:        []string{"teamwork", "collaboration", "help"},
		},
		{
			Description: "Show gratitude: Thank colleagues for their contributions.",
			Tags:        []string{"thanks", "thank", "gratitude"},
		},
	}
}
