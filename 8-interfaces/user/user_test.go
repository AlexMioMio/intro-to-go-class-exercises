package user

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Group struct {
	name Name
}

func (g Group) Name() Name {
	return g.name
}

func TestID(t *testing.T) {
	assert := assert.New(t)

	u, _ := NewUser("ringo", "apple")

	if !assert.Equal("Name", reflect.TypeOf(u.Name()).Name(), "Name() returns an Name variable") {
		t.Fatal("Cannot continue unless Name() returns an Name variable")
	}

	assert.Equal(u.username, getName(u), "User type implements HasName interface correctly")

	g := Group{name: "admins"}
	assert.Equal(g.name, getName(g), "HasName interface works for Group type")
}

func getName(v HasName) Name {
	return v.Name()
}
