package release

import (
	"strconv"
	"testing"

	"github.com/giantswarm/apiextensions/v2/pkg/apis/release/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_Release_mustFindPatch(t *testing.T) {
	testCases := []struct {
		name            string
		version         string
		release         v1alpha1.Release
		releases        []v1alpha1.Release
		expectedRelease v1alpha1.Release
	}{
		{
			name:    "case 0",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
		},
		{
			name:    "case 1",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.1"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.2"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.2"},
			},
		},
		{
			name:    "case 2",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.2.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"},
			},
		},
		{
			name:    "case 3",
			version: "v12.0.0",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"},
			},
		},
		{
			name:    "case 4",
			version: "v12.0.9",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"},
			},
		},
		{
			name:    "case 5",
			version: "v12.0.0-dev",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"},
			},
		},
		{
			name:    "case 6",
			version: "v12.0.0-dev",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.3-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v12.0.3-dev"},
			},
		},
		{
			name:    "case 7",
			version: "v100.0.0-xh3b4sd",
			releases: []v1alpha1.Release{
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.0-dev"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0-xh3b4sd"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v100.0.3"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.0.5"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "v12.1.1"}},
			},
			expectedRelease: v1alpha1.Release{
				ObjectMeta: metav1.ObjectMeta{Name: "v100.0.0-xh3b4sd"},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			release := mustFindPatch(tc.version, tc.releases)

			if !cmp.Equal(release, tc.expectedRelease) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedRelease, release))
			}
		})
	}
}
