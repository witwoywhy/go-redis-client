package gedis

type RESPType string

const (
	BulkStrings     RESPType = "$"
	BulkStringsNull RESPType = "$-1"
	Arrays          RESPType = "*"
	Integers        RESPType = ":"
)

type TimeOptions string

const (
	EX      TimeOptions = "EX"
	PX      TimeOptions = "PX"
	EXAT    TimeOptions = "EXAT"
	PXAT    TimeOptions = "PXAT"
	PERSIST TimeOptions = "PERSIST"
)

const (
	AUTH = "AUTH"
)

const (
	SELECT = "SELECT"
)

const (
	APPEND      = "APPEND"
	DECR        = "DECR"
	DECRBY      = "DECRBY"
	GET         = "GET"
	GETDEL      = "GETDEL"
	GETEX       = "GETEX"
	GETRANGE    = "GETRANGE"
	GETSET      = "GETSET"
	INCR        = "INCR"
	INCRBY      = "INCRBY"
	INCRBYFLOAT = "INCRBYFLOAT"

	SET  = "SET"
	MGET = "MGET"
	MSET = "MSET"
)
