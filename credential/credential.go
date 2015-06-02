package credential

type Credential struct {
	BetfairUsername string `yaml:"betfairUsername"`
	BetfairPassword string `yaml:"betfairPassword"`
	ApplicationKey string `yaml:"applicationKey"`
}