package log

type Fields struct {
	Timestamp string
	Level     string
	Msg       string
	Func      string
	File      string
	Line      int

	Context
}
