package mybolt

import "unsafe"

type pgid uint64

const (
	branchPageFlag   = 0x01
	leafPageFlag     = 0x02
	metaPageFlag     = 0x04
	freelistPageFlag = 0x10
)

type page struct {
	id    pgid
	flags uint16
	ptr   bool
}

func (p *page) meta() *meta {
	return (*meta)(unsafe.Pointer(&p.ptr))
}

func (p *page) freelist() *freelist {

}

type freelist struct {
}
