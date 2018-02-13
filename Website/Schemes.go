package Website

// Basis contains data commonly used by all templates.
type Basis struct {
	Name                  string
	Version               string
	Lang                  string
	Logo                  string
	Session               string
	SiteVerificationToken string
}

//Add adds two integers. Utility method as templates do not provide arithmetic operations.
func (basis Basis) Add(a, b int) int {
	return a + b
}
