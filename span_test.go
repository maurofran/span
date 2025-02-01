package span_test

import (
	"github.com/maurofran/span"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func intSpan(t *testing.T, start, end int) span.Span[int] {
	t.Helper()
	fixture, err := span.Between(start, end)
	require.NoError(t, err)
	return fixture
}

func TestSpan_Contains(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.False(t, fixture.Contains(4))
	assert.True(t, fixture.Contains(5))
	assert.True(t, fixture.Contains(6))
	assert.True(t, fixture.Contains(9))
	assert.False(t, fixture.Contains(10))
}

func TestSpan_ContainsSpan(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.False(t, fixture.ContainsSpan(intSpan(t, 4, 10)))
	assert.True(t, fixture.ContainsSpan(intSpan(t, 5, 9)))
	assert.True(t, fixture.ContainsSpan(intSpan(t, 6, 9)))
}

func TestSpan_Fit(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.Equal(t, 5, fixture.Fit(4))
	assert.Equal(t, 5, fixture.Fit(5))
	assert.Equal(t, 6, fixture.Fit(6))
	assert.Equal(t, 9, fixture.Fit(9))
	assert.Equal(t, 9, fixture.Fit(10))
}

func TestSpan_Intersect(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.Equal(t, intSpan(t, 5, 9), fixture.Intersect(intSpan(t, 5, 9)))
	assert.Equal(t, intSpan(t, 5, 5), fixture.Intersect(intSpan(t, 5, 5)))
	assert.Equal(t, intSpan(t, 5, 6), fixture.Intersect(intSpan(t, 3, 6)))
	assert.Equal(t, intSpan(t, 7, 9), fixture.Intersect(intSpan(t, 7, 12)))
}

func TestSpan_After(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.After(4))
	assert.False(t, fixture.After(5))
}

func TestSpan_AfterSpan(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.AfterSpan(intSpan(t, 2, 3)))
	assert.False(t, fixture.AfterSpan(intSpan(t, 3, 8)))
}

func TestSpan_Before(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.Before(10))
	assert.False(t, fixture.Before(7))
}

func TestSpan_BeforeSpan(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.BeforeSpan(intSpan(t, 10, 12)))
	assert.False(t, fixture.BeforeSpan(intSpan(t, 6, 8)))
}

func TestSpan_StartedBy(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.StartedBy(5))
	assert.False(t, fixture.StartedBy(4))
}

func TestSpan_EndedBy(t *testing.T) {
	fixture := intSpan(t, 5, 9)
	assert.True(t, fixture.EndedBy(9))
	assert.False(t, fixture.EndedBy(10))
}

func timeSpan(t *testing.T, start, end time.Time) span.Span[time.Time] {
	t.Helper()
	fixture, err := span.BetweenTimes(start, end)
	require.NoError(t, err)
	return fixture
}
