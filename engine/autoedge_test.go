// Mgmt
// Copyright (C) 2013-2023+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//go:build !root

package engine

import (
	"testing"
)

func TestIFF1(t *testing.T) {
	uid := &BaseUID{Name: "/tmp/unit-test"}
	same := &BaseUID{Name: "/tmp/unit-test"}
	diff := &BaseUID{Name: "/tmp/other-file"}

	if !uid.IFF(same) {
		t.Errorf("basic resource UIDs with the same name should satisfy each other's IFF condition")
	}

	if uid.IFF(diff) {
		t.Errorf("basic resource UIDs with different names should NOT satisfy each other's IFF condition")
	}
}
