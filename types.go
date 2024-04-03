package gedis

type Multiple struct {
	Key   string
	Value any
}

type TTL struct {
	Option TimeOptions
	Time   string
}
