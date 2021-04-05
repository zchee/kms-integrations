package fakekms

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"

	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

func TestAllowlistScalars(t *testing.T) {
	var cases = []struct {
		Message proto.GeneratedMessage
		Path    string
	}{
		{
			Message: &kmspb.CreateKeyRingRequest{Parent: "foo"},
			Path:    "parent",
		},
		{
			Message: &kmspb.CreateCryptoKeyRequest{SkipInitialVersionCreation: true},
			Path:    "skip_initial_version_creation",
		},
		{
			Message: &kmspb.ListCryptoKeysRequest{PageSize: 1337},
			Path:    "page_size",
		},
		{
			Message: &kmspb.CreateKeyRingRequest{
				KeyRing: &kmspb.KeyRing{Name: "foo"},
			},
			Path: "key_ring.name",
		},
		{
			Message: &kmspb.CreateCryptoKeyVersionRequest{
				CryptoKeyVersion: &kmspb.CryptoKeyVersion{
					Algorithm: kmspb.CryptoKeyVersion_EC_SIGN_P256_SHA256,
				},
			},
			Path: "crypto_key_version.algorithm",
		},
		{
			Message: &kmspb.CreateCryptoKeyRequest{
				CryptoKey: &kmspb.CryptoKey{
					VersionTemplate: &kmspb.CryptoKeyVersionTemplate{
						ProtectionLevel: kmspb.ProtectionLevel_HSM,
					},
				},
			},
			Path: "crypto_key.version_template.protection_level",
		},
	}

	for _, c := range cases {
		t.Run(c.Path, func(t *testing.T) {
			if err := allowlist(c.Path).check(c.Message); err != nil {
				t.Errorf("unexpected allowlist failure: %v", err)
			}

			err := allowlist().check(c.Message)
			st, _ := status.FromError(err)
			if st.Code() != codes.Unimplemented {
				t.Errorf("expected UNIMPLEMENTED for unallowed field %s, got %v", c.Path, err)
			}
		})
	}
}

func TestAllowlistSuccessMultipleFields(t *testing.T) {
	req := &kmspb.CreateCryptoKeyRequest{
		Parent: "bar",
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_DECRYPT,
		},
	}

	var cases = []struct {
		Name      string
		Allowlist protoAllowlister
	}{
		{
			Name:      "ExactFields",
			Allowlist: allowlist("parent", "crypto_key.purpose"),
		},
		{
			Name:      "ExtraField",
			Allowlist: allowlist("parent", "crypto_key_id", "crypto_key.purpose"),
		},
		{
			Name:      "NestedMessage",
			Allowlist: allowlist("parent", "crypto_key"),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if err := c.Allowlist.check(req); err != nil {
				t.Fatalf("expected success, got %v", err)
			}
		})
	}
}

func TestAllowlistFailureMultipleFields(t *testing.T) {
	req := &kmspb.CreateCryptoKeyRequest{
		Parent: "bar",
		CryptoKey: &kmspb.CryptoKey{
			Purpose: kmspb.CryptoKey_ASYMMETRIC_DECRYPT,
		},
	}

	var cases = []struct {
		Name      string
		Allowlist protoAllowlister
	}{
		{
			Name:      "MissingPurpose",
			Allowlist: allowlist("parent"),
		},
		{
			Name:      "MissingParent",
			Allowlist: allowlist("crypto_key.purpose"),
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			err := c.Allowlist.check(req)
			st, _ := status.FromError(err)
			if st.Code() != codes.Unimplemented {
				t.Errorf("expected UNIMPLEMENTED, got %v", err)
			}
		})
	}
}

func TestAllowlistFailsUnknownFields(t *testing.T) {
	kr := proto.MessageV2(&kmspb.KeyRing{Name: "foo"})
	kr.ProtoReflect().SetUnknown(protoreflect.RawFields{0xF0, 0x01})

	err := allowlist("name").check(kr)
	st, _ := status.FromError(err)
	if st.Code() != codes.Unimplemented {
		t.Errorf("expected UNIMPLEMENTED for unknown fields, got %v", err)
	}
}