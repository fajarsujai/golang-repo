package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	// TODO: Initialize
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		// TODO: Initialize
		for i := 0; i < b.N; i++ {
			// TODO: Your Code Here
			HelloWorld("Eko")
		}
	})

	b.Run("Kurniawan", func(b *testing.B) {
		// TODO: Initialize
		for i := 0; i < b.N; i++ {
			// TODO: Your Code Here
			HelloWorld("Kurniawan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

//Menggagalkan Test
func TestHelloWorldSujai(t *testing.T) {
	result := HelloWorld("Sujai")

	if result != "Hello Eko" {
		//error
		t.Fatal("Result must be 'Hello Sujai'")
	}
	fmt.Println("TestHelloworldSujai Done")
}

// Test Assertion
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// Require == FailnNow Test
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Require Done")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("Must run on Mac OS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Eko)",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorld(Kurniawan)",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Hello must be"+test.request)
		})
	}
}
