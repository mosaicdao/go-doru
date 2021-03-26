package core

type Options struct {
	ThreadsBadgerRepoPath string
}

type Option func(*Options)

func WithBadgerThreadsPersistance(repositoryPath string) Option {
	return func(options *Options) {
		options.ThreadsBadgerRepoPath = repositoryPath
	}
}
