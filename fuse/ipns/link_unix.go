package ipns

import (
	"os"

	"github.com/jbenet/go-ipfs/Godeps/_workspace/src/bazil.org/fuse"
	"github.com/jbenet/go-ipfs/Godeps/_workspace/src/golang.org/x/net/context"
)

type Link struct {
	Target string
}

func (l *Link) Attr() fuse.Attr {
	log.Debug("Link attr.")
	return fuse.Attr{
		Mode: os.ModeSymlink | 0555,
	}
}

func (l *Link) Readlink(req *fuse.ReadlinkRequest, ctx context.Context) (string, error) {
	log.Debugf("ReadLink: %s", l.Target)
	return l.Target, nil
}
