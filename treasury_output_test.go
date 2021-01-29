package iota_test

import (
	"errors"
	"testing"

	"github.com/iotaledger/iota.go/v2"
	"github.com/stretchr/testify/assert"
)

func TestTreasuryOutput_Deserialize(t *testing.T) {
	type test struct {
		name   string
		source []byte
		target *iota.TreasuryOutput
		err    error
	}
	tests := []test{
		func() test {
			treasuryOutput, treasuryOutputData := randTreasuryOutput()
			return test{"ok- w/o treasuryOutput", treasuryOutputData, treasuryOutput, nil}
		}(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			treasuryOutput := &iota.TreasuryOutput{}
			bytesRead, err := treasuryOutput.Deserialize(tt.source, iota.DeSeriModePerformValidation)
			if tt.err != nil {
				assert.True(t, errors.Is(err, tt.err))
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, len(tt.source), bytesRead)
			assert.EqualValues(t, tt.target, treasuryOutput)
		})
	}
}

func TestTreasuryOutput_Serialize(t *testing.T) {
	type test struct {
		name   string
		source *iota.TreasuryOutput
		target []byte
	}
	tests := []test{
		func() test {
			treasuryOutput, treasuryOutputData := randTreasuryOutput()
			return test{"ok- w/o treasuryOutput", treasuryOutput, treasuryOutputData}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edData, err := tt.source.Serialize(iota.DeSeriModePerformValidation)
			assert.NoError(t, err)
			assert.Equal(t, tt.target, edData)
		})
	}
}
