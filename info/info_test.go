package info

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

type Test struct {
	BuildNumber string `imprint:"BUILD_NUMBER"`
}

func TestEncoding(t *testing.T) {
	t.Run("Meta data tag handling", func(t *testing.T) {
		err := os.Setenv("BUILD_NUMBER", "foo")
		assert.NoError(t, err)

		info := InfoBuild{}
		ti := reflect.TypeOf(info)

		assert.Equal(t, "github.com/grahambrooks/imprint/info", ti.PkgPath())

		for i := 0; i < ti.NumField(); i++ {
			field := ti.Field(i)

			if field.Type.Kind() == reflect.Struct {

			}
			imprintTag := field.Tag.Get("imprint")
			fmt.Printf("%s %s\n", field.Name, imprintTag)
			of := reflect.ValueOf(&info)
			s := of.Elem()
			s.Field(i).SetString(os.Getenv(imprintTag))
		}

		assert.Equal(t, "foo", info.BuildNumber)
	})

	t.Run("Imprinting", func(t *testing.T) {
		s := Test{}
		BuildNumber = "wow"

		err := Imprint(&s)
		assert.NoError(t, err)
		assert.Equal(t, "wow", s.BuildNumber)
	})
	t.Run("Imprinting Info Block", func(t *testing.T) {
		s := InfoBlock{}
		BuildNumber = "wow"

		err := Imprint(&s)
		assert.NoError(t, err)
		assert.Equal(t, "wow", s.Build.BuildNumber)
	})
}
