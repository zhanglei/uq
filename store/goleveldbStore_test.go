package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewLevelStore(t *testing.T) {
	Convey("Test New Level Store", t, func() {
		ldb, err := NewLevelStore("/tmp/ldb")
		So(err, ShouldBeNil)
		So(ldb, ShouldNotBeNil)
		ldb.Close()
	})
}

func TestSetLevel(t *testing.T) {
	Convey("Test Level Store Set", t, func() {
		ldb, err := NewLevelStore("/tmp/ldb")
		So(err, ShouldBeNil)
		So(ldb, ShouldNotBeNil)
		defer ldb.Close()

		err = ldb.Set("foo", []byte("bar"))
		So(err, ShouldBeNil)
	})
}

func TestGetLevel(t *testing.T) {
	Convey("Test Level Store Get", t, func() {
		ldb, err := NewLevelStore("/tmp/ldb")
		So(err, ShouldBeNil)
		So(ldb, ShouldNotBeNil)
		defer ldb.Close()

		err = ldb.Set("foo", []byte("bar"))
		So(err, ShouldBeNil)

		data, err := ldb.Get("foo")
		So(err, ShouldBeNil)
		So(string(data), ShouldEqual, "bar")
	})
}

func TestDelLevel(t *testing.T) {
	Convey("Test Level Store Del", t, func() {
		ldb, err := NewLevelStore("/tmp/ldb")
		So(err, ShouldBeNil)
		So(ldb, ShouldNotBeNil)
		defer ldb.Close()

		err = ldb.Set("foo", []byte("bar"))
		So(err, ShouldBeNil)

		err = ldb.Del("foo")
		So(err, ShouldBeNil)
	})
}