package sampletest

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		y    int
		want int
	}{
		{
			name: "a",
			x:    1,
			y:    2,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := add(tt.x, tt.y)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
		fmt.Println(jwt.UnsafeAllowNoneSignatureType)
	}
}
