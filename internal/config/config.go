package config

type Config struct {
	StatusOnly     bool
	JSONOutput     bool
	FollowRedirect bool
	TimeoutSeconds int
	Targets        []string
}
