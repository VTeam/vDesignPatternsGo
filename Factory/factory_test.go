package factory

import "testing"

func TestLogFactory(t *testing.T) {

	var factory LogFactory

	factory = &LogFileFactory{
		handlers: make([]Handler, 0),
	}

	fileHandler := factory.Create("file")
	fileHandler.Sink()

	factory = &LogStdOutFactory{
		handlers: make([]Handler, 0),
	}

	stdoutHandler := factory.Create("stdout")
	stdoutHandler.Sink()

}
