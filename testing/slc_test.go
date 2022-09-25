package testing

import (
	"reflect"
	"testing"

	"github.com/guoyangguo/go-unit-test/slc"
)

func TestRemoveItemWithSelf(t *testing.T) {
	inputs := []string{"a", "b", "c", "d", "d", "c"}
	outputs := slc.RemoveItemWithSelf(inputs, "c")
	expected := []string{"a", "b", "d", "d"}
	if !reflect.DeepEqual(outputs, expected) {
		t.Errorf("outputs:%v,expected:%v", outputs, expected)
	}
}

func TestRemoveItemByCopy(t *testing.T) {
	inputs := []string{"a", "b", "c", "d", "d", "c"}
	outputs := slc.RemoveItemByCopy(inputs, "c")
	expected := []string{"a", "b", "d", "d"}
	if !reflect.DeepEqual(outputs, expected) {
		t.Errorf("outputs:%v,expected:%v", outputs, expected)
	}

}

func TestRemoveItemBySub(t *testing.T) {
	inputs := []string{"a", "b", "c", "d", "d", "c"}
	outputs := slc.RemoveItemBySub(inputs, "c")
	expected := []string{"a", "b", "d", "d"}
	if !reflect.DeepEqual(outputs, expected) {
		t.Errorf("outputs:%v,expected:%v", outputs, expected)
	}
}

func TestRemoveItemWithShare(t *testing.T) {
	inputs := []string{"a", "b", "c", "d", "d", "c"}
	outputs := slc.RemoveItemWithShare(inputs, "c")
	expected := []string{"a", "b", "d", "d"}
	if !reflect.DeepEqual(outputs, expected) {
		t.Errorf("outputs:%v,expected:%v", outputs, expected)
	}
}
