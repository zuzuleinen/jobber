package sources

//Get all available sources
func All() []Source {
	s := make([]Source, 0)
	s = append(s, BerlinStartupJobs{"http://berlinstartupjobs.com"})
	s = append(s, StackOverflow{"http://stackoverflow.com"})
	return s
}
