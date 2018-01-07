//  Copyright 2016 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package intern

import (
	"testing"
)

func TestSetsAddIndex(t *testing.T) {
	s := NewSetOfPatterns()
	zanys := []*Pattern{newZAny()}
	want := s.Add(zanys)
	got := s.Index(zanys)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}
	if s.Get(want).NullIndex != 0 {
		t.Fatal("nullindex != 0")
	}
	if !s.Get(want).Accept {
		t.Fatal("not accept")
	}
	notzanys := []*Pattern{newNotZAny(), newNotZAny()}
	if s.Index(notzanys) != -1 {
		t.Fatal("not found")
	}
	state := s.Add(notzanys)
	if s.Get(state).NullIndex != 1 {
		t.Fatal("nullindex != 1")
	}
	if s.SetOfBits.Index(s.Get(state).Nullables) != s.Get(state).NullIndex {
		t.Fatal("wrong nullindex")
	}
	if s.Get(state).Accept {
		t.Fatal("accept")
	}
}