/*
 * @Author: UerAx
 * @Date: 2022-06-13 21:30:50
 * @FilePath: \steam-concurrent-users\steam\api_test.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package steam

import "testing"

func Test_findConcurrentUsersByName(t *testing.T) {
	getGameNameCodeMap()
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{"Plants vs. Zombies"}, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			get, err := FindConcurrentUsersByName(tt.args.name)
			t.Log(get)
			if (err != nil) != tt.wantErr {
				t.Errorf("findConcurrentUsersByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
