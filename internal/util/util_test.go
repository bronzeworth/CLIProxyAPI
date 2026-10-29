package util

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestResolveAuthDir(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	absolute := t.TempDir()

	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty defaults to standard dir under home",
			in:   "",
			want: filepath.Join(home, ".cli-proxy-api"),
		},
		{
			name: "tilde alone resolves to home",
			in:   "~",
			want: home,
		},
		{
			name: "tilde slash resolves under home",
			in:   "~/auths",
			want: filepath.Join(home, "auths"),
		},
		{
			name: "tilde backslash resolves under home",
			in:   `~\auths`,
			want: filepath.Join(home, "auths"),
		},
		{
			name: "tilde slash nested path resolves under home",
			in:   "~/a/b",
			want: filepath.Join(home, "a", "b"),
		},
		{
			name: "tilde with trailing separator resolves to home",
			in:   "~/",
			want: home,
		},
		{
			name: "absolute path is returned unchanged",
			in:   absolute,
			want: absolute,
		},
		{
			name: "relative path is returned relative",
			in:   "auths",
			want: "auths",
		},
		{
			name: "relative path is cleaned",
			in:   "auths/./sub/../keys",
			want: filepath.Join("auths", "keys"),
		},
		{
			name: "absolute path is cleaned",
			in:   absolute + "/sub/../keys/",
			want: filepath.Join(absolute, "keys"),
		},
		{
			name: "tilde path is cleaned",
			in:   "~/a/../b",
			want: filepath.Join(home, "b"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveAuthDir(tt.in)
			if err != nil {
				t.Fatalf("ResolveAuthDir(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Errorf("ResolveAuthDir(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestResolveAuthDirRejectsTildeUser(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	tests := []struct {
		name string
		in   string
	}{
		{name: "tilde user with path", in: "~bob/auths"},
		{name: "tilde user alone", in: "~bob"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveAuthDir(tt.in)
			if err == nil {
				t.Fatalf("ResolveAuthDir(%q) = %q, want error", tt.in, got)
			}
			if !strings.Contains(err.Error(), tt.in) {
				t.Errorf("ResolveAuthDir(%q) error %q does not name the value", tt.in, err)
			}
		})
	}
}
