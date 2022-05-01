/*
Copyright SecureKey Technologies Inc. All Rights Reserved.
Copyright Boran Car <boran.car@gmail.com>. All Rights Reserved.
Copyright Christian Nuss <christian@scaffold.ly>, Founder, Scaffoldly LLC. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dynamodb_test

import (
	"testing"

	. "github.com/hyperledger/aries-framework-go-ext/component/storage/dynamodb"
	dynamock "github.com/gusaul/go-dynamock"
)

func TestCreate(t *testing.T) {
	dbapi, _ := dynamock.New()

	p, err := NewProvider(WithMockDBAPI(dbapi))
	if err != nil || p == nil {
		t.Errorf("Provider not created")
	}
}
