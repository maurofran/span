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

func timeSpan(t *testing.T, start, end time.Time) span.Span[time.Time] {
	t.Helper()
	fixture, err := span.BetweenTimes(start, end)
	require.NoError(t, err)
	return fixture
}
