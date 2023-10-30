package taproot

import (
	"errors"
	"strings"

	"github.com/mleku/ec/chaincfg"
)

// AddressSegWit is the base address type for all SegWit addresses.
type AddressSegWit struct {
	hrp            string
	witnessVersion byte
	witnessProgram []byte
}

// AddressTaproot is an Address for a pay-to-taproot (P2TR) output. See BIP 341
// for further details.
type AddressTaproot struct {
	AddressSegWit
}

// NewAddressTaproot returns a new AddressTaproot.
func NewAddressTaproot(witnessProg []byte,
	net *chaincfg.Params) (*AddressTaproot, error) {

	return newAddressTaproot(net.Bech32HRPSegwit, witnessProg)
}

// newAddressWitnessScriptHash is an internal helper function to create an
// AddressWitnessScriptHash with a known human-readable part, rather than
// looking it up through its parameters.
func newAddressTaproot(hrp string,
	witnessProg []byte) (*AddressTaproot, error) {

	// Check for valid program length for witness version 1, which is 32
	// for P2TR.
	if len(witnessProg) != 32 {
		return nil, errors.New("witness program must be 32 bytes for " +
			"p2tr")
	}

	addr := &AddressTaproot{
		AddressSegWit{
			hrp:            strings.ToLower(hrp),
			witnessVersion: 0x01,
			witnessProgram: witnessProg,
		},
	}

	return addr, nil
}
