// Copyright 2015 Diffeo, Inc.
// This software is released under an MIT/X11 open source license.

package postgres_test

import (
	"github.com/benbjohnson/clock"
	"github.com/diffeo/go-coordinate/coordinate/coordinatetest"
	"github.com/diffeo/go-coordinate/postgres"
	"gopkg.in/check.v1"
	"testing"
)

// Test is the top-level entry point to run tests.
//
// This creates a PostgreSQL Coordinate backend using an empty string
// as the connection string.  This means that, when you run "go test",
// you must set environment variables as describe in
// http://www.postgresql.org/docs/current/static/libpq-envars.html
func Test(t *testing.T) { check.TestingT(t) }

func init() {
	clk := clock.NewMock()
	c, err := postgres.NewWithClock("", clk)
	if err != nil {
		panic(err)
	}
	check.Suite(&coordinatetest.Suite{
		Coordinate: c,
		Clock:      clk,
	})
}
