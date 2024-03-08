package db

import (
	"testing"
	"time"
)


func TestRefreshDate(t *testing.T) {
    staleDate := time.Date(2021,01,01,0,0,0,0,time.Local)
    want := time.Date(time.Now().Year(),01,01,0,0,0,0,time.Local)
    res  := RefreshDate(&staleDate)
    if !want.Equal(res)  {
        t.Fatalf(`arg1: %q arg2:%q`, res, want)
    }
}

