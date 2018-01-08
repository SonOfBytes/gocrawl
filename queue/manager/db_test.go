package manager_test

import (
	"testing"

	"time"

	"github.com/sonofbytes/gocrawl/queue/manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNew(t *testing.T) {
	db := manager.New()
	assert.IsType(t, &manager.DB{}, db)
}

func TestDB_Expire(t *testing.T) {
	db := manager.New()

	// default set
	x := db.GetExpire()
	assert.Equal(t, time.Minute*15, x)

	// override works
	db.SetExpire(time.Second)
	x = db.GetExpire()
	assert.Equal(t, time.Second, x)
}

func TestDB_Add(t *testing.T) {
	db := manager.New()

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	err := db.Add("", 0, "job")
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	err = db.Add("someurl", -1, "job")
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	err = db.Add("someurl", 0, "")
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, e.Code())

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	err = db.Add("someurl", 0, "job")
	assert.NoError(t, err)

	assert.Equal(t, 1, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())
}

func TestDB_Get(t *testing.T) {
	db := manager.New()

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	url, depth, job, err := db.Get()
	assert.Error(t, err)
	e, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, e.Code())
	assert.Equal(t, "", url)
	assert.Equal(t, int32(0), depth)
	assert.Equal(t, "", job)

	err = db.Add("someurl", 0, "job")
	assert.NoError(t, err)

	assert.Equal(t, 1, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	url, depth, job, err = db.Get()
	assert.NoError(t, err)
	assert.Equal(t, codes.NotFound, e.Code())
	assert.Equal(t, "someurl", url)
	assert.Equal(t, int32(0), depth)
	assert.Equal(t, "job", job)

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 1, db.DoneCount())

	url, depth, job, err = db.Get()
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.NotFound, e.Code())

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 1, db.DoneCount())

	// duplicate processed url
	err = db.Add("someurl", 0, "job")
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.AlreadyExists, e.Code())
	assert.Contains(t, e.Message(), "processed recently")

	// duplicate new url
	err = db.Add("newurl", 0, "job")
	assert.NoError(t, err)

	err = db.Add("newurl", 0, "job")
	assert.Error(t, err)
	e, ok = status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, codes.AlreadyExists, e.Code())
	assert.Contains(t, e.Message(), "submitted recently")
}

func TestDB_Purge(t *testing.T) {
	db := manager.New()

	db.SetExpire(time.Millisecond)

	err := db.Add("someurl", 0, "job")
	assert.NoError(t, err)

	assert.Equal(t, 1, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())

	url, depth, job, err := db.Get()
	assert.NoError(t, err)
	assert.Equal(t, "someurl", url)
	assert.Equal(t, int32(0), depth)
	assert.Equal(t, "job", job)

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 1, db.DoneCount())

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 0, db.NewCount())
	assert.Equal(t, 0, db.DoneCount())
}
