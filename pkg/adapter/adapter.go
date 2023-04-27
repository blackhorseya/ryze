package adapter

// Restful restful adapter
type Restful interface {
	InitRouting() error
}

// Cronjob is a cronjob adapter
type Cronjob interface {
	// Start to run
	Start() error

	// Stop to end
	Stop() error
}

// CLI is a command line interface adapter
type CLI interface {
	// Execute serve caller to execute command
	Execute() error
}
