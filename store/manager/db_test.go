package manager_test

import (
	"testing"

	"github.com/sonofbytes/gocrawl/store/manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNew(t *testing.T) {
	db := manager.New()
	assert.IsType(t, &manager.DB{}, db)
}

func TestDB_Add(t *testing.T) {
	db := manager.New()

	err := db.Add("", []string{})
	assert.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, s.Code())

	err = db.Add("someurl", []string{})
	assert.NoError(t, err)
}

func TestDB_Get(t *testing.T) {
	db := manager.New()

	urls, err := db.Get("")
	assert.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, s.Code())
	assert.Equal(t, []string{}, urls)

	urls, err = db.Get("unknown")
	assert.Error(t, err)
	s, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, s.Code())
	assert.Equal(t, []string{}, urls)

	db.Add("someurl", []string{})
	urls, err = db.Get("someurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{}, urls)

	db.Add("anotherurl", []string{"one"})
	urls, err = db.Get("anotherurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{"one"}, urls)

	db.Add("manyurl", []string{"one", "two", "three"})
	urls, err = db.Get("manyurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{"one", "two", "three"}, urls)

	// ensure previous unchanged
	urls, err = db.Get("someurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{}, urls)

	urls, err = db.Get("anotherurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{"one"}, urls)

	// overwrite
	db.Add("someurl", []string{"overwrite"})
	urls, err = db.Get("someurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{"overwrite"}, urls)

	db.Add("manyurl", []string{"1", "2", "3"})
	urls, err = db.Get("manyurl")
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "2", "3"}, urls)
}
