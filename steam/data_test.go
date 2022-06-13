/*
 * @Author: UerAx
 * @Date: 2022-06-13 20:45:09
 * @FilePath: \steam-concurrent-users\steam\data_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */

package steam

import "testing"

func Test_getGameNameCodeMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getGameNameCodeMap()
			t.Logf("%v", hash)
		})
	}
}
