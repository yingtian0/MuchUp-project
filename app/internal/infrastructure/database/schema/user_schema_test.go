package schema

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	gormschema "gorm.io/gorm/schema"
)

func TestUserSchemaMessagesRelation(t *testing.T) {
	parsed, err := gormschema.Parse(&UserSchema{}, &sync.Map{}, gormschema.NamingStrategy{})
	require.NoError(t, err)

	relation := parsed.Relationships.Relations["Messages"]
	require.NotNil(t, relation)
	require.Len(t, relation.References, 1)
	require.Equal(t, "ID", relation.References[0].PrimaryKey.Name)
	require.Equal(t, "SenderID", relation.References[0].ForeignKey.Name)
}
