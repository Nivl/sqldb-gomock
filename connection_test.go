package mocksqldb_test

import (
	"testing"

	"github.com/Nivl/sqldb-gomock"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestQueryableWithConnection(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocksqldb.NewMockConnection(mockCtrl)
	m.QEXPECT().GetNotFound(&testStruct{})

	output := &testStruct{}
	err := m.Get(output, "select * ...", "uuid")
	assert.Error(t, err, "m.Get() should have failed")
}
