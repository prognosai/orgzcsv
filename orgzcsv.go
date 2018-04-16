// orgzcsv hello
package orgzcsv

import (
	"compress/gzip"
	"encoding/csv"
	"os"
	"strings"

	"github.com/pkg/errors"
)

//  holds stuff
type Reader struct {
	f   *os.File
	gzr *gzip.Reader
	Csv *csv.Reader
}

func isGz(name string) bool {
	return strings.HasSuffix(name, ".gz")
}

func Open(name string) (*Reader, error) {

	r := new(Reader)
	var err error

	if r.f, err = os.Open(name); err != nil {
		return nil, errors.Wrap(err, "Opening of the csv (gz|not) failed.")
	}

	if isGz(name) {
		if r.gzr, err = gzip.NewReader(r.f); err != nil {
			r.f.Close()
			return nil, errors.Wrap(err, "Creating the gz reader for the csv failed.")
		}
	}

	if r.gzr != nil {
		r.Csv = csv.NewReader(r.gzr)
	} else {
		r.Csv = csv.NewReader(r.f)
	}

	return r, nil
}

func (r *Reader) Close() error {

	if r.gzr != nil {
		if err := r.gzr.Close(); err != nil {
			return errors.Wrap(err, "Closing of the gz reader for (gz|not) failed.")
		}
	}

	if err := r.f.Close(); err != nil {
		return errors.Wrap(err, "Closing of the file for (gz|not) failed.")
	}

	return nil
}
