package manager

import (
	"math/rand"
	"sync"
	"time"

	"log"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const expireDuration = time.Minute * 15

type DB struct {
	expireDuration time.Duration
	new            sync.Map
	newLen         int64
	done           sync.Map
	doneLen        int64
	lock           sync.Mutex
}

type Record struct {
	job     string
	depth   int32
	start   time.Time
	expires time.Time
}

func init() {
	// ensure our job ids are seeded randomly-ish
	rand.Seed(time.Now().UTC().UnixNano())
}

func New() *DB {
	db := &DB{
		expireDuration: expireDuration,
	}
	go db.purge()
	go db.log()

	return db
}

func (db *DB) SetExpire(duration time.Duration) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.expireDuration = duration
}

func (db *DB) GetExpire() time.Duration {
	db.lock.Lock()
	defer db.lock.Unlock()
	return db.expireDuration
}

func (db *DB) NewCount() int {
	return int(atomic.LoadInt64(&db.newLen))
}

func (db *DB) DoneCount() int {
	return int(atomic.LoadInt64(&db.doneLen))
}

func (db *DB) Get() (url string, depth int32, job string, err error) {
	db.new.Range(func(k, v interface{}) bool {
		// There is a race condition here between get and done store where another get can operate - but it's PoC
		u := k.(string)
		r := v.(*Record)
		r.start = time.Now()
		r.expires = r.start.Add(db.GetExpire())

		// whether we store or it exists we delete from new
		db.Delete(u)
		if _, ok := db.done.LoadOrStore(u, r); ok {
			return true
		}
		atomic.AddInt64(&db.doneLen, 1)

		url = u
		depth = r.depth
		job = r.job
		return false
	})

	if url != "" {
		return url, depth, job, nil
	}
	return "", 0, "", status.Error(codes.NotFound, "Queue empty")
}

func (db *DB) Delete(url string) {
	db.new.Delete(url)
	atomic.AddInt64(&db.newLen, -1)
}

func (db *DB) Add(url string, depth int32, job string) error {
	if url == "" || depth < 0 || job == "" {
		return status.Error(codes.InvalidArgument, "url and job needs to be set and depth should be zero or positive")
	}

	if _, ok := db.done.Load(url); ok {
		return status.Error(codes.AlreadyExists, "submitted URL has already been processed recently")
	}

	_, l := db.new.LoadOrStore(url, &Record{
		job:     job,
		depth:   depth,
		expires: time.Now().Add(db.GetExpire()),
	})
	if l {
		return status.Error(codes.AlreadyExists, "submitted URL has already been submitted recently")
	}

	atomic.AddInt64(&db.newLen, 1)

	return nil
}

func (db *DB) log() {
	for {
		log.Printf("Queue lengths - New: %d  Done: %d", db.NewCount(), db.DoneCount())
		time.Sleep(10 * time.Second)
	}
}

func (db *DB) purge() {
	for {
		db.done.Range(func(k, v interface{}) bool {
			if v.(*Record).expires.Before(time.Now()) {
				db.done.Delete(k.(string))
				atomic.AddInt64(&db.doneLen, -1)
			}
			time.Sleep(min(10*time.Millisecond, db.GetExpire()/10))
			return true
		})
		time.Sleep(min(10*time.Millisecond, db.GetExpire()/10))
	}
}

func min(d ...time.Duration) time.Duration {
	min := d[0]
	for _, i := range d[1:] {
		if i < min {
			min = i
		}
	}
	return min
}
