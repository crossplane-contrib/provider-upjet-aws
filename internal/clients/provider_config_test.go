// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestFingerprintCredentials(t *testing.T) {
	base := aws.Credentials{AccessKeyID: "AKID", SecretAccessKey: "secret", SessionToken: "token"}

	t.Run("IsStableAndDoesNotLeakTheMaterial", func(t *testing.T) {
		got := fingerprintCredentials(base)
		if got != fingerprintCredentials(base) {
			t.Error("fingerprintCredentials(...) is not stable for the same credentials")
		}
		if len(got) != 2*sha256.Size {
			t.Errorf("fingerprintCredentials(...): want a hex sha256 digest, got %q", got)
		}
		for _, s := range []string{base.AccessKeyID, base.SecretAccessKey, base.SessionToken} {
			if strings.Contains(got, s) {
				t.Errorf("fingerprintCredentials(...) leaks the credential material %q", s)
			}
		}
	})

	t.Run("IsKeyedToThisProcess", func(t *testing.T) {
		// the digest must not be reproducible from the credential material
		// alone, so that it cannot be confirmed by someone holding a candidate
		// credential and reading the cache keys off the debug logs.
		h := sha256.New()
		for _, s := range []string{base.AccessKeyID, base.SecretAccessKey, base.SessionToken} {
			_ = binary.Write(h, binary.BigEndian, uint64(len(s)))
			_, _ = h.Write([]byte(s))
		}
		if unkeyed := fmt.Sprintf("%x", h.Sum(nil)); fingerprintCredentials(base) == unkeyed {
			t.Error("fingerprintCredentials(...) is an unkeyed digest of the credential material")
		}
	})

	t.Run("DiffersForEveryField", func(t *testing.T) {
		cases := map[string]aws.Credentials{
			"AccessKeyID":     {AccessKeyID: "other", SecretAccessKey: base.SecretAccessKey, SessionToken: base.SessionToken},
			"SecretAccessKey": {AccessKeyID: base.AccessKeyID, SecretAccessKey: "other", SessionToken: base.SessionToken},
			"SessionToken":    {AccessKeyID: base.AccessKeyID, SecretAccessKey: base.SecretAccessKey, SessionToken: "other"},
			// the fields are separated while hashing, so shifting a field
			// boundary must change the digest
			"ShiftedBoundary": {AccessKeyID: base.AccessKeyID + base.SecretAccessKey, SecretAccessKey: "", SessionToken: base.SessionToken},
		}
		want := fingerprintCredentials(base)
		for name, creds := range cases {
			t.Run(name, func(t *testing.T) {
				if got := fingerprintCredentials(creds); got == want {
					t.Errorf("fingerprintCredentials(...): want a digest different from %q, got the same", want)
				}
			})
		}
	})

	t.Run("FramingIsUnambiguous", func(t *testing.T) {
		// AWS credentials never contain a NUL byte, but the framing must not
		// depend on that: these two collided when the fields were merely
		// suffixed with a NUL separator instead of length-prefixed.
		cases := map[string][2]aws.Credentials{
			"NULInsideAField": {
				{AccessKeyID: "a\x00b", SecretAccessKey: "c"},
				{AccessKeyID: "a", SecretAccessKey: "b\x00c"},
			},
			"EmptyVersusAbsent": {
				{AccessKeyID: "a", SecretAccessKey: "", SessionToken: "b"},
				{AccessKeyID: "a", SecretAccessKey: "b", SessionToken: ""},
			},
		}
		for name, pair := range cases {
			t.Run(name, func(t *testing.T) {
				if a, b := fingerprintCredentials(pair[0]), fingerprintCredentials(pair[1]); a == b {
					t.Errorf("fingerprintCredentials(...): distinct credentials digested to the same value %q", a)
				}
			})
		}
	})
}