// Package segment provides abstraction for splitting files into segments.
package segment

import "io"

// Segment holds a reference to a file
// along with range defined as offset and length.
type Segment struct {
	file   io.ReadSeeker
	Offset int64
	Length int64
}

// newSegment constructor.
func newSegment(file io.ReadSeeker, offset int64, length int64) Segment {
	return Segment{
		file:   file,
		Offset: offset,
		Length: length,
	}
}

// Split divides file segment into pieces of a given length.
func (f *Segment) Split(length int64) (parts []Segment) {
	for position := int64(0); position < f.Length; position += length {
		part := newSegment(f.file, f.Offset+position, length)
		if position+length > f.Length {
			part.Length = f.Length % length
		}
		parts = append(parts, part)
	}

	return
}

// Bytes allows to read contents of a segment.
func (f *Segment) Bytes() (data []byte, err error) {
	buffer := make([]byte, f.Length)
	_, err = f.file.Seek(f.Offset, io.SeekStart)
	if err != nil {
		return
	}
	n, err := f.file.Read(buffer)
	if err != nil {
		return
	}
	data = buffer[:n]

	return
}

// MakeSegments splits file into segments of a given size.
func MakeSegments(
	file io.ReadSeeker,
	partSize int64,
) (parts []Segment, err error) {

	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return
	}

	part := newSegment(file, 0, fileSize)
	parts = part.Split(partSize)

	return
}
