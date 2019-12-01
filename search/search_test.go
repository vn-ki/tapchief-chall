package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexAndSearch1(t *testing.T) {
	err := Index(`
	Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Magna ac placerat vestibulum lectus. Elit duis tristique sollicitudin nibh sit amet commodo. Senectus et netus et malesuada fames. Fermentum iaculis eu non diam phasellus vestibulum lorem sed. Dictumst quisque sagittis purus sit amet volutpat consequat mauris. Aliquam ut porttitor leo a diam sollicitudin tempor. Consectetur a erat nam at lectus urna duis convallis. Sed viverra ipsum nunc aliquet bibendum enim facilisis gravida neque. 



Maecenas volutpat blandit aliquam etiam erat velit scelerisque. Lectus sit amet est placerat in egestas erat imperdiet. Ante in nibh mauris cursus mattis. Tellus rutrum tellus pellentesque eu tincidunt. Euismod quis viverra nibh cras pulvinar mattis. Proin nibh nisl condimentum id venenatis a. Quam elementum pulvinar etiam non quam. Arcu dictum varius duis at consectetur lorem donec. Aliquet porttitor lacus luctus accumsan tortor. Duis ut diam quam nulla porttitor massa id.
	`)
	assert.Nil(t, err)

	searched, err := Search("lorem")
	assert.Nil(t, err)
	assert.Equal(t, len(searched), 2, "Both paragraphs should be returned")

	searched, err = Search("maecenas")
	assert.Nil(t, err)
	assert.Equal(t, len(searched), 1, "1 paragraph should be returned")
}

func TestClear(t *testing.T) {
	err := Index(`blah blah`)
	assert.Nil(t, err)

	searched, err := Search("blah")
	assert.Nil(t, err)
	assert.Equal(t, len(searched), 1, "It should be present before clearing")

	ClearIndex()

	searched, err = Search("blah")
	assert.Nil(t, err)
	assert.Equal(t, len(searched), 0, "no results should be found")
}
