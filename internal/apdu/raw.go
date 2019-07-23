package apdu

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Header represents the fixed-size header for the APDU structure
type Header struct {
	CLA, INS, P1, P2 byte
}

// Read is a dummy function allowing this to be used with io.Copy
func (Header) Read(b []byte) (n int, err error) { panic("this value cannot be read. use WriteTo instead") }

// ReadFrom implements the io.ReaderFrom interface.
func (h *Header) ReadFrom(rd io.Reader) (n int64, err error) {
	return int64(binary.Size(h)), binary.Read(rd, endianness, h)
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (h *Header) UnmarshalBinary(buf []byte) (err error) {
	_, err = h.ReadFrom(bytes.NewReader(buf))
	return
}

// WriteTo implements the io.WriterTo interface.
func (h Header) WriteTo(w io.Writer) (n int64, err error) {
	return int64(binary.Size(h)), binary.Write(w, endianness, h)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (h Header) MarshalBinary() (b []byte, err error) {
	var buf bytes.Buffer
	if _, err = h.WriteTo(&buf); err != nil {
		return
	}

	return buf.Bytes(), nil
}


// Command represents the body of an APDU command
type Command struct {
	Lc      []byte // 0, 1, or 3 bytes
	Command []byte
	Le      []byte // 0, 1, 2, or 3 bytes
}

// Read is a dummy function allowing this to be used with io.Copy
func (Command) Read(b []byte) (n int, err error) { panic("this value cannot be read. use WriteTo() instead") }

// WriteTo implements the io.WriterTo interface.
func (c Command) WriteTo(w io.Writer) (n int64, err error) {
	for _, bts := range [][]byte{c.Lc, c.Command, c.Le} {
		var written int
		written, err = w.Write(bts)
		if err != nil {
			return
		}

		n += int64(written)
	}

	return
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (c Command) MarshalBinary() (b []byte, err error) {
	var buf bytes.Buffer
	if _, err = c.WriteTo(&buf); err != nil {
		return
	}

	return buf.Bytes(), nil
}

// Raw represents the raw binary of an APDU structure
type Raw struct {
	Header
	Command
}


// WriteTo implements the io.WriterTo interface.
func (r Raw) WriteTo(w io.Writer) (n int64, err error) {
	return io.Copy(w, io.MultiReader(r.Header, r.Command))
}

func (r Raw) Read(b []byte) (n int64, err error) { panic("this value cannot be read. use WriteTo() instead") }