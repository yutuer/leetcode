package mybolt

import (
	"github.com/pkg/errors"
	"hash/fnv"
	"os"
	"time"
	"unsafe"
)

// The data file format version.
const version = 2

// Represents a marker value to indicate that a file is a Bolt DB.
const magic uint32 = 0xED0CDAED

var (
	// ErrInvalid is returned when both meta pages on a database are invalid.
	// This typically occurs when a file is not a bolt database.
	ErrInvalid = errors.New("ErrInvalid")

	// ErrVersionMismatch is returned when the data file was created with a
	// different version of Bolt.
	ErrVersionMismatch = errors.New("version mismatch")

	// ErrChecksum is returned when either meta page checksum does not match.
	ErrChecksum = errors.New("checksum error")
)

var defaultPageSize = os.Getpagesize()

// DefaultOptions represent the options used if nil options are passed into Open().
// No timeout is used which will cause Bolt to wait indefinitely for a lock.
var DefaultOptions = &Options{
	Timeout:    0,
	NoGrowSync: false,
}

type DB struct {
	// When true, skips the truncate call when growing the database.
	//	// Setting this to true is only safe on non-ext3/ext4 systems.
	//	// Skipping truncation avoids preallocation of hard drive space and
	//	// bypasses a truncate() and fsync() syscall on remapping.
	//	//
	//	// https://github.com/boltdb/bolt/issues/284
	NoGrowSync bool

	// If you want to read the entire database fast, you can set MmapFlag to
	// syscall.MAP_POPULATE on Linux 2.6.23+ for sequential read-ahead.
	MmapFlags int

	opened bool

	path string

	file *os.File

	ops struct {
		writeAt func(b []byte, off int64) (n int, err error)
	}
	// 页大小
	pageSize int
}

// Open creates and opens a database at the given path.
// If the file does not exist then it will be created automatically.
// Passing in nil options will cause Bolt to open the database with the default options.
func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	db := &DB{opened: true, pageSize: os.Getpagesize()}

	// Set default options if no options are provided.
	if options == nil {
		options = DefaultOptions
	}
	db.NoGrowSync = options.NoGrowSync
	db.MmapFlags = options.MmapFlags

	// Open ptr file and separate sync handler for metadata writes.
	db.path = path
	var err error
	db.file, err = os.OpenFile(db.path, os.O_RDWR|os.O_CREATE, mode)
	if err != nil {
		_ = db.Close()
		return nil, err
	}

	// Default values for test hooks
	db.ops.writeAt = db.file.WriteAt

	if info, err := db.file.Stat(); err != nil {
		return nil, err
	} else if info.Size() == 0 {
		if err := db.init(); err != nil {
			return nil, err
		}
	} else {
		// Read the first meta page to determine the page size.
		var buf [0x1000]byte //4k
		if _, err := db.file.ReadAt(buf[:], 0); err == nil {
			// 检验前4k个字节
			m := db.pageInBuffer(buf[:], 0).meta()
			if err := m.validate(); err != nil {

			}

		} else {

		}
	}

	return db, err
}

func (db *DB) String() string {
	return ""
}

func (db *DB) Close() error {
	db.file.Close()
}

func (db *DB) init() error {
	var data [os.Getpagesize() * 4]byte

	var pages [4]page
	for i, p := range pages {
		if i < 2 {
			//2 meta pages
		} else if i == 2 {
			//freelist
		} else {
			//branch
		}
	}

}

// 找到buf内的第i个页
// pageInBuffer retrieves a page reference from a given byte array based on the current page size.
func (db *DB) pageInBuffer(buf []byte, id pgid) *page {
	return (*page)(unsafe.Pointer(&buf[id*pgid(db.pageSize)]))
}

// Options represents the options that can be set when opening a database.
type Options struct {
	// Timeout is the amount of time to wait to obtain a file lock.
	// When set to zero it will wait indefinitely. This option is only
	// available on Darwin and Linux.
	Timeout time.Duration

	// Sets the DB.NoGrowSync flag before memory mapping the file.
	NoGrowSync bool

	// Sets the DB.MmapFlags flag before memory mapping the file.
	MmapFlags int
}

type meta struct {
	magic    uint32
	version  uint32
	pageSize uint32

	checksum uint64
}

func (m *meta) magicFunc() {
	m.magic = 0xFFFFFFFF

}

func (m *meta) sum64() uint64 {
	var h = fnv.New64a()
	_, _ = h.Write(  [unsafe.Offsetof(meta{}.checksum)]byte)  unsafe.Pointer(m))
	return h.Sum64()
}

func (m *meta) validate() error {
	if m.magic != magic {
		return ErrInvalid
	} else if m.version != version {
		return ErrVersionMismatch
	} else if m.checksum != 0 && m.checksum != m.sum64() {
		return ErrChecksum
	}

	return nil
}
