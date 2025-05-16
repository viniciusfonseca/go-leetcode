package main

import (
	"math/rand"
	"testing"
)

type RandomizedSet struct {
	keys  []int
	state map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		state: make(map[int]int),
		keys:  make([]int, 0), // stores the keys from state
	}
}

func (set *RandomizedSet) Insert(val int) bool {
	if _, ok := set.state[val]; ok {
		return false
	}
	set.keys = append(set.keys, val)
	set.state[val] = len(set.keys) - 1 // store the last index from keys
	return true
}

func (set *RandomizedSet) Remove(val int) bool {
	i, ok := set.state[val]
	if !ok {
		return false
	}
	set.keys[i], set.keys[len(set.keys)-1] = set.keys[len(set.keys)-1], set.keys[i] // swap the removed with the last
	set.state[set.keys[i]] = i

	delete(set.state, val)
	set.keys = set.keys[:len(set.keys)-1]
	return true
}

func (set *RandomizedSet) GetRandom() int {
	return set.keys[rand.Intn(len(set.keys))]
}

func TestRandomizedSet(t *testing.T) {
	set := Constructor()

	ret := set.Insert(1)
	if !ret {
		t.Errorf("set.Insert(1) = %t; want true", ret)
	}
	ret = set.Remove(2)
	if ret {
		t.Errorf("set.Remove(2) = %t; want false", ret)
	}
	ret = set.Insert(2)
	if !ret {
		t.Errorf("set.Insert(2) = %t; want true", ret)
	}
	iret := set.GetRandom()
	if iret == 0 {
		t.Errorf("set.GetRandom() = %d; want 1 or 2", iret)
	}
	ret = set.Remove(1)
	if !ret {
		t.Errorf("set.Remove(1) = %t; want true", ret)
	}
	ret = set.Insert(2)
	if ret {
		t.Errorf("set.Insert(2) = %t; want false", ret)
	}
	iret = set.GetRandom()
	if iret != 2 {
		t.Errorf("set.GetRandom() = %d; want 2", iret)
	}

}
