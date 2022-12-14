/*
Copyright 2020 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"github.com/openebs/api/v3/pkg/apis/types"
	"reflect"
	"testing"

	cstor "github.com/openebs/api/v3/pkg/apis/cstor/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsBlockDeviceReplacementCase(t *testing.T) {
	type args struct {
		newRaidGroup cstor.RaidGroup
		oldRaidGroup cstor.RaidGroup
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case#1:Not a blockdevice replacement case",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: false,
		},

		{
			name: "case#2:Not a blockdevice replacement case",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},
			},
			want: false,
		},

		{
			name: "case#3:A blockdevice replacement case",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-5"},
					},
				},
			},
			want: true,
		},

		{
			name: "case#4:A blockdevice replacement case",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-3"},
					},
				},
			},
			want: true,
		},

		// Following cases are not replacement case although it should be rejected.

		{
			name: "case#5:Not a blockdevice replacement case: SHOULD BE REJECTED",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-1"},
					},
				},
			},
			want: false,
		},

		// Following cases are still replacement case although INVALID and will be rejected.

		{
			name: "case#6:A blockdevice replacement case: Althoguh INVALID, should be REJECTED",
			args: args{
				newRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},

				oldRaidGroup: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-6"},
						{BlockDeviceName: "bd-5"},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt // pin it
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlockDeviceReplacementCase(&tt.args.newRaidGroup, &tt.args.oldRaidGroup); got != tt.want {
				t.Errorf("IsBlockDeviceReplacementCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumberOfDiskReplaced(t *testing.T) {
	type args struct {
		newRG cstor.RaidGroup
		oldRG cstor.RaidGroup
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case#1:No block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: 0,
		},

		{
			name: "case#2:No block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-1"},
					},
				},
			},
			want: 0,
		},

		{
			name: "case#3:1 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-3"},
					},
				},
			},
			want: 1,
		},

		{
			name: "case#4:1 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-5"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},
			},
			want: 1,
		},

		// Following test case is a invalid type of bd replacement and hence will be rejected finally by validations.
		{
			name: "case#5:2 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-4"},
						{BlockDeviceName: "bd-3"},
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt // pin it
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberOfDiskReplaced(&tt.args.newRG, &tt.args.oldRG); got != tt.want {
				t.Errorf("GetNumberOfDiskReplaced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMoreThanOneDiskReplaced(t *testing.T) {
	type args struct {
		newRG cstor.RaidGroup
		oldRG cstor.RaidGroup
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case#1:No block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: false,
		},

		{
			name: "case#2:No block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-1"},
					},
				},
			},
			want: false,
		},

		{
			name: "case#3:1 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-3"},
					},
				},
			},
			want: false,
		},

		{
			name: "case#4:1 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-5"},
						{BlockDeviceName: "bd-3"},
						{BlockDeviceName: "bd-4"},
					},
				},
			},
			want: false,
		},

		// Following test case is a invalid type of bd replacement and hence will be rejected finally by validations.
		{
			name: "case#5:2 block device replaced",
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-4"},
						{BlockDeviceName: "bd-3"},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt // pin it
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMoreThanOneDiskReplaced(&tt.args.newRG, &tt.args.oldRG); got != tt.want {
				t.Errorf("IsMoreThanOneDiskReplaced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockDeviceReplacement_IsNewBDPresentOnCurrentCSPC(t *testing.T) {
	type fields struct {
		OldCSPC *cstor.CStorPoolCluster
		NewCSPC *cstor.CStorPoolCluster
	}
	type args struct {
		newRG cstor.RaidGroup
		oldRG cstor.RaidGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Case#1: New BD present on current CSPC",
			fields: fields{
				OldCSPC: &cstor.CStorPoolCluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cspc-mirror",
						Namespace: "openebs",
					},
					Spec: cstor.CStorPoolClusterSpec{
						Pools: []cstor.PoolSpec{
							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-1"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-1"},
											{BlockDeviceName: "bd-2"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-3"},
											{BlockDeviceName: "bd-4"},
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-3"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: true,
		},

		{
			name: "Case#2: New BD  present on current CSPC",
			fields: fields{
				OldCSPC: &cstor.CStorPoolCluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cspc-mirror",
						Namespace: "openebs",
					},
					Spec: cstor.CStorPoolClusterSpec{
						Pools: []cstor.PoolSpec{
							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-1"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-1"},
											{BlockDeviceName: "bd-2"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-3"},
											{BlockDeviceName: "bd-4"},
										},
									},
								},
							},

							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-2"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-5"},
											{BlockDeviceName: "bd-6"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-7"},
											{BlockDeviceName: "bd-8"},
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-7"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: true,
		},

		{
			name: "Case#3: New BD not present on current CSPC",
			fields: fields{
				OldCSPC: &cstor.CStorPoolCluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cspc-mirror",
						Namespace: "openebs",
					},
					Spec: cstor.CStorPoolClusterSpec{
						Pools: []cstor.PoolSpec{
							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-1"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-1"},
											{BlockDeviceName: "bd-2"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-3"},
											{BlockDeviceName: "bd-4"},
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-8"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: false,
		},

		{
			name: "Case#4: New BD not present on current CSPC",
			fields: fields{
				OldCSPC: &cstor.CStorPoolCluster{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "cspc-mirror",
						Namespace: "openebs",
					},
					Spec: cstor.CStorPoolClusterSpec{
						Pools: []cstor.PoolSpec{
							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-1"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-1"},
											{BlockDeviceName: "bd-2"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-3"},
											{BlockDeviceName: "bd-4"},
										},
									},
								},
							},

							{
								NodeSelector: map[string]string{"kubernetes.io/hostname": "node-2"},
								DataRaidGroups: []cstor.RaidGroup{
									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-5"},
											{BlockDeviceName: "bd-6"},
										},
									},

									{
										CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
											{BlockDeviceName: "bd-7"},
											{BlockDeviceName: "bd-8"},
										},
									},
								},
							},
						},
					},
				},
			},
			args: args{
				newRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-9"},
					},
				},

				oldRG: cstor.RaidGroup{
					CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
						{BlockDeviceName: "bd-1"},
						{BlockDeviceName: "bd-2"},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt // pin it
		t.Run(tt.name, func(t *testing.T) {
			bdr := &PoolOperations{
				OldCSPC: tt.fields.OldCSPC,
				NewCSPC: tt.fields.NewCSPC,
			}
			if got := bdr.IsNewBDPresentOnCurrentCSPC(&tt.args.newRG, &tt.args.oldRG); got != tt.want {
				t.Errorf("BlockDeviceReplacement.IsNewBDPresentOnCurrentCSPC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateRaidGroupChanges(t *testing.T) {
	tests := map[string]struct {
		oldRG         *cstor.RaidGroup
		newRG         *cstor.RaidGroup
		oldRgType     string
		expectedError bool
	}{
		"removing block devices": {
			oldRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
					{BlockDeviceName: "bd-2"},
				},
			},
			newRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
				},
			},
			expectedError: true,
		},
		"adding block devices for raid groups": {
			oldRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
					{BlockDeviceName: "bd-2"},
				},
			},
			newRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
					{BlockDeviceName: "bd-2"},
					{BlockDeviceName: "bd-3"},
				},
			},
			oldRgType:     "raidz",
			expectedError: true,
		},
		"adding block devices for stripe raid groups": {
			oldRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
					{BlockDeviceName: "bd-2"},
				},
			},
			newRG: &cstor.RaidGroup{
				CStorPoolInstanceBlockDevices: []cstor.CStorPoolInstanceBlockDevice{
					{BlockDeviceName: "bd-1"},
					{BlockDeviceName: "bd-2"},
					{BlockDeviceName: "bd-3"},
				},
			},
			oldRgType:     "stripe",
			expectedError: false,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			err := validateRaidGroupChanges(test.oldRG, test.newRG, test.oldRgType)
			if test.expectedError && err == nil {
				t.Errorf("test %s failed expectedError to be error but got nil", name)
			}
			if !test.expectedError && err != nil {
				t.Errorf("test %s failed expectedError not to be error but got error %v", name, err)
			}
		})
	}
}

func Test_getAllowedTagMap(t *testing.T) {
	type args struct {
		cspcAnnotation map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "Test case #1",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: "fast,slow"},
			},
			want: map[string]bool{"fast": true, "slow": true},
		},

		{
			name: "Test case #2",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: "fast,slow"},
			},
			want: map[string]bool{"slow": true, "fast": true},
		},

		{
			name: "Test case #3 -- Nil Annotations",
			args: args{
				cspcAnnotation: nil,
			},
			want: map[string]bool{},
		},

		{
			name: "Test case #4 -- No BD tag Annotations",
			args: args{
				cspcAnnotation: map[string]string{"some-other-annotation-key": "awesome-openebs"},
			},
			want: map[string]bool{},
		},

		{
			name: "Test case #5 -- Improper format 1",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: ",fast,slow,,"},
			},
			want: map[string]bool{"fast": true, "slow": true},
		},

		{
			name: "Test case #6 -- Improper format 2",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: ",fast,slow"},
			},
			want: map[string]bool{"fast": true, "slow": true},
		},

		{
			name: "Test case #7 -- Improper format 2",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: ",fast,,slow"},
			},
			want: map[string]bool{"fast": true, "slow": true},
		},

		{
			name: "Test case #7 -- Improper format 2",
			args: args{
				cspcAnnotation: map[string]string{types.OpenEBSAllowedBDTagKey: "this is improper"},
			},
			want: map[string]bool{"this is improper": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllowedTagMap(tt.args.cspcAnnotation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllowedTagMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
