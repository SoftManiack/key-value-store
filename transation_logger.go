package main

import "os"

type EventType byte

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type TransactionLogger interface {
	WriteDelete(key string)
	WritePut(key string)
}

type FileTransactionLogger struct {
	events       chan<- Event // Канал только для записи: для передачи событий
	errors       <-chan error // Канал только для чтения для рпиема ошибок
	lastSequense uint64       // Последний использованный порядковый номер
	file         *os.File     //местоположение файла журнала
}

func (l *FileTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *FileTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *FileTransactionLogger) Err() <-chan error {
	return l.errors
}
