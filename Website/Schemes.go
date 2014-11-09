package Website

type Basis struct {
	Name    string
	Version string
	Lang    string
}

type PagePassword struct {
	Basis        Basis
	TextPassword string
}

type PageStart struct {
	Basis           Basis
	TextWelcome     string
	TextStartButton string
}
