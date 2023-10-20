package query_test

import (
	"testing"
	"time"

	op_cmd "github.com/angrypufferfish/goodm/src/operators"
	"github.com/angrypufferfish/goodm/src/query"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestInc(t *testing.T) {

	value1 := bson.E{"field", 12}
	value2 := []bson.E{{"field", 12}, {"field", "fieldValue"}}

	expression1 := query.Inc(value1)
	expression2 := query.Inc(value2...)

	assert.Equal(t, bson.D{{op_cmd.INC, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.INC, value2}}, expression2)
}

func TestMin(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", 12}, {"field", 234}}

	expression1 := query.Min(value1)
	expression2 := query.Min(value2...)

	assert.Equal(t, bson.D{{op_cmd.MIN, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.MIN, value2}}, expression2)
}

func TestMax(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", 12}, {"field", 234}}

	expression1 := query.Max(value1)
	expression2 := query.Max(value2...)

	assert.Equal(t, bson.D{{op_cmd.MAX, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.MAX, value2}}, expression2)
}

func TestMul(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", 12}, {"field", 234}}

	expression1 := query.Mul(value1)
	expression2 := query.Mul(value2...)

	assert.Equal(t, bson.D{{op_cmd.MUL, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.MUL, value2}}, expression2)
}

func TestRename(t *testing.T) {

	value1 := bson.E{"field", "newFieldName"}
	value2 := []bson.E{{"field", "newFieldName"}, {"field", "newFieldName"}}

	expression1 := query.Rename(value1)
	expression2 := query.Rename(value2...)

	assert.Equal(t, bson.D{{op_cmd.RENAME, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.RENAME, value2}}, expression2)
}

func TestSet(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "fieldValue"}, {"field", 234}}

	expression1 := query.Set(value1)
	expression2 := query.Set(value2)

	assert.Equal(t, bson.D{{op_cmd.SET, value1}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.SET, value2}}, expression2)
}

func TestSetOnInsert(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.SetOnInsert(value1)
	expression2 := query.SetOnInsert(value2...)

	assert.Equal(t, bson.D{{op_cmd.SET_ON_INSERT, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.SET_ON_INSERT, value2}}, expression2)
}

func TestUnset(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.Unset(value1)
	expression2 := query.Unset(value2...)

	assert.Equal(t, bson.D{{op_cmd.UNSET, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.UNSET, value2}}, expression2)
}

func TestCurrentDate(t *testing.T) {

	value1 := bson.E{"field", time.Now().UTC()}
	value2 := []bson.E{{"field", time.Now().UTC()}, {"field", time.Now().UTC()}}

	expression1 := query.CurrentDate(value1)
	expression2 := query.CurrentDate(value2...)

	assert.Equal(t, bson.D{{op_cmd.CURRENT_DATE, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.CURRENT_DATE, value2}}, expression2)
}

func TestAddToSet(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.AddToSet(value1)
	expression2 := query.AddToSet(value2...)

	assert.Equal(t, bson.D{{op_cmd.ADD_TO_SET, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.ADD_TO_SET, value2}}, expression2)
}

func TestPop(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.Pop(value1)
	expression2 := query.Pop(value2...)

	assert.Equal(t, bson.D{{op_cmd.POP, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.POP, value2}}, expression2)
}

func TestPull(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.Pull(value1)
	expression2 := query.Pull(value2...)

	assert.Equal(t, bson.D{{op_cmd.PULL, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.PULL, value2}}, expression2)
}

func TestPush(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.Push(value1)
	expression2 := query.Push(value2...)

	assert.Equal(t, bson.D{{op_cmd.PUSH, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.PUSH, value2}}, expression2)
}

func TestPullAll(t *testing.T) {

	value1 := bson.E{"field", -12}
	value2 := []bson.E{{"field", "FieldValue"}, {"field", true}}

	expression1 := query.PullAll(value1)
	expression2 := query.PullAll(value2...)

	assert.Equal(t, bson.D{{op_cmd.PULL_ALL, []bson.E{value1}}}, expression1)
	assert.Equal(t, bson.D{{op_cmd.PULL_ALL, value2}}, expression2)
}
