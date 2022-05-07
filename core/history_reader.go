package core

import "io"

type HistoryReader struct {
	io.Reader
	historyBuffer []byte
	reader        io.Reader
}

func NewHistoryReader(r io.Reader) *HistoryReader {
	return &HistoryReader{
		reader: r,
	}
}

func (hr *HistoryReader) Read(p []byte) (n int, err error) {
	n, err = hr.reader.Read(p)
	if n > 0 {
		hr.historyBuffer = append(hr.historyBuffer, p[:n]...)
	}
	return n, err
}

func (hr *HistoryReader) HistoryBuffer() []byte {
	return hr.historyBuffer
}
