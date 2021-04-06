// Code generated by entc, DO NOT EDIT.

package entgen

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	Channel        []ent.Hook
	EmbedMessage   []ent.Hook
	Emote          []ent.Hook
	EmotePack      []ent.Hook
	File           []ent.Hook
	FileHash       []ent.Hook
	FileMessage    []ent.Hook
	ForeignUser    []ent.Hook
	Guild          []ent.Hook
	GuildListEntry []ent.Hook
	Invite         []ent.Hook
	LocalUser      []ent.Hook
	Message        []ent.Hook
	Permission     []ent.Hook
	Profile        []ent.Hook
	Role           []ent.Hook
	Session        []ent.Hook
	TextMessage    []ent.Hook
	User           []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
