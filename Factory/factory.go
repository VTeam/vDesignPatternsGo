package factory

import "fmt"

type LogFactory interface {
	Create(name string) Handler
	createHandler(name string) Handler
	registerHandler(h Handler)
}

type Handler interface {
	Sink()
	GetName() string
}

type LogFileHandler struct {
	name string
}

func (lf *LogFileHandler) GetName() string {

	return lf.name
}
func (lf *LogFileHandler) Sink() {
	fmt.Printf("Use %s Log Handler handle Log!\n", lf.name)
}

type LogFileFactory struct {
	handlers []Handler
}

func (t *LogFileFactory) Create(name string) Handler {
	handler := t.createHandler(name)
	t.registerHandler(handler)
	return handler
}

func (t *LogFileFactory) createHandler(name string) Handler {
	return &LogFileHandler{name: name}
}

func (t *LogFileFactory) registerHandler(h Handler) {
	t.handlers = append(t.handlers, h)
}

type LogStdOutHandler struct {
	name string
}

func (ls *LogStdOutHandler) GetName() string {

	return ls.name
}
func (ls *LogStdOutHandler) Sink() {
	fmt.Printf("Use %s Log Handler handle Log!\n", ls.name)
}

type LogStdOutFactory struct {
	handlers []Handler
}

func (t *LogStdOutFactory) Create(name string) Handler {
	handler := t.createHandler(name)
	t.registerHandler(handler)
	return handler
}

func (t *LogStdOutFactory) createHandler(name string) Handler {
	return &LogFileHandler{name: name}
}

func (t *LogStdOutFactory) registerHandler(h Handler) {
	t.handlers = append(t.handlers, h)
}
