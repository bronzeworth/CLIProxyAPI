package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v8/internal/config"
)

func writeServiceAccountKey(t *testing.T, dir string) string {
	t.Helper()
	key, errGen := rsa.GenerateKey(rand.Reader, 2048)
	if errGen != nil {
		t.Fatalf("generate rsa key: %v", errGen)
	}
	der, errMarshal := x509.MarshalPKCS8PrivateKey(key)
	if errMarshal != nil {
		t.Fatalf("marshal private key: %v", errMarshal)
	}
	pemKey := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})
	payload, errJSON := json.Marshal(map[string]any{
		"type":         "service_account",
		"project_id":   "demo-project",
		"client_email": "svc@demo-project.iam.gserviceaccount.com",
		"private_key":  string(pemKey),
	})
	if errJSON != nil {
		t.Fatalf("marshal service account: %v", errJSON)
	}
	keyPath := filepath.Join(dir, "key.json")
	if errWrite := os.WriteFile(keyPath, payload, 0o600); errWrite != nil {
		t.Fatalf("write key file: %v", errWrite)
	}
	return keyPath
}

func TestDoVertexImportWritesNothingForUnresolvableAuthDir(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	workDir := t.TempDir()
	t.Chdir(workDir)
	keyPath := writeServiceAccountKey(t, t.TempDir())

	DoVertexImport(&config.Config{AuthDir: "~bob/auths"}, keyPath, "")

	if _, errStat := os.Stat(filepath.Join(workDir, "~bob")); !os.IsNotExist(errStat) {
		t.Errorf("literal ./~bob directory was created (stat err: %v)", errStat)
	}
	for _, root := range []string{home, workDir} {
		entries, errRead := os.ReadDir(root)
		if errRead != nil {
			t.Fatalf("read dir %s: %v", root, errRead)
		}
		if len(entries) != 0 {
			t.Errorf("expected %s to stay empty, found %d entries (first: %s)", root, len(entries), entries[0].Name())
		}
	}
}
