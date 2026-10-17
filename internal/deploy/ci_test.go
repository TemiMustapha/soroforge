package deploy

import (
	"os"
	"testing"
)

func TestCIRunURLGitHubActions(t *testing.T) {
	// Save original values
	orig := map[string]string{
		"GITHUB_ACTIONS":    os.Getenv("GITHUB_ACTIONS"),
		"GITHUB_SERVER_URL": os.Getenv("GITHUB_SERVER_URL"),
		"GITHUB_REPOSITORY": os.Getenv("GITHUB_REPOSITORY"),
		"GITHUB_RUN_ID":     os.Getenv("GITHUB_RUN_ID"),
	}
	defer func() {
		for k, v := range orig {
			if v == "" {
				os.Unsetenv(k)
			} else {
				os.Setenv(k, v)
			}
		}
	}()

	// Set GitHub Actions environment
	os.Setenv("GITHUB_ACTIONS", "true")
	os.Setenv("GITHUB_SERVER_URL", "https://github.com")
	os.Setenv("GITHUB_REPOSITORY", "soroworks/soroforge")
	os.Setenv("GITHUB_RUN_ID", "12345")

	url := ciRunURL()
	expected := "https://github.com/soroworks/soroforge/actions/runs/12345"
	if url != expected {
		t.Errorf("ciRunURL() = %q, want %q", url, expected)
	}
}

func TestCIRunURLGitLabCI(t *testing.T) {
	// Save original values
	orig := map[string]string{
		"GITLAB_CI":       os.Getenv("GITLAB_CI"),
		"CI_SERVER_URL":   os.Getenv("CI_SERVER_URL"),
		"CI_PROJECT_PATH": os.Getenv("CI_PROJECT_PATH"),
		"CI_PIPELINE_ID":  os.Getenv("CI_PIPELINE_ID"),
	}
	defer func() {
		for k, v := range orig {
			if v == "" {
				os.Unsetenv(k)
			} else {
				os.Setenv(k, v)
			}
		}
	}()

	// Clear GitHub Actions env to avoid interference
	os.Unsetenv("GITHUB_ACTIONS")

	// Set GitLab CI environment
	os.Setenv("GITLAB_CI", "true")
	os.Setenv("CI_SERVER_URL", "https://gitlab.com")
	os.Setenv("CI_PROJECT_PATH", "soroworks/soroforge")
	os.Setenv("CI_PIPELINE_ID", "67890")

	url := ciRunURL()
	expected := "https://gitlab.com/soroworks/soroforge/-/pipelines/67890"
	if url != expected {
		t.Errorf("ciRunURL() = %q, want %q", url, expected)
	}
}

func TestCIRunURLOutsideCI(t *testing.T) {
	// Save original values
	orig := map[string]string{
		"GITHUB_ACTIONS":    os.Getenv("GITHUB_ACTIONS"),
		"GITLAB_CI":         os.Getenv("GITLAB_CI"),
		"GITHUB_SERVER_URL": os.Getenv("GITHUB_SERVER_URL"),
		"GITHUB_REPOSITORY": os.Getenv("GITHUB_REPOSITORY"),
		"GITHUB_RUN_ID":     os.Getenv("GITHUB_RUN_ID"),
		"CI_SERVER_URL":     os.Getenv("CI_SERVER_URL"),
		"CI_PROJECT_PATH":   os.Getenv("CI_PROJECT_PATH"),
		"CI_PIPELINE_ID":    os.Getenv("CI_PIPELINE_ID"),
	}
	defer func() {
		for k, v := range orig {
			if v == "" {
				os.Unsetenv(k)
			} else {
				os.Setenv(k, v)
			}
		}
	}()

	// Clear all CI environment variables
	os.Unsetenv("GITHUB_ACTIONS")
	os.Unsetenv("GITLAB_CI")
	os.Unsetenv("GITHUB_SERVER_URL")
	os.Unsetenv("GITHUB_REPOSITORY")
	os.Unsetenv("GITHUB_RUN_ID")
	os.Unsetenv("CI_SERVER_URL")
	os.Unsetenv("CI_PROJECT_PATH")
	os.Unsetenv("CI_PIPELINE_ID")

	url := ciRunURL()
	if url != "" {
		t.Errorf("ciRunURL() = %q, want empty string outside CI", url)
	}
}

func TestCIRunURLIncompleteGitHubEnv(t *testing.T) {
	// Save original values
	orig := map[string]string{
		"GITHUB_ACTIONS":    os.Getenv("GITHUB_ACTIONS"),
		"GITHUB_SERVER_URL": os.Getenv("GITHUB_SERVER_URL"),
		"GITHUB_REPOSITORY": os.Getenv("GITHUB_REPOSITORY"),
		"GITHUB_RUN_ID":     os.Getenv("GITHUB_RUN_ID"),
	}
	defer func() {
		for k, v := range orig {
			if v == "" {
				os.Unsetenv(k)
			} else {
				os.Setenv(k, v)
			}
		}
	}()

	// Set partial GitHub Actions environment (missing RUN_ID)
	os.Setenv("GITHUB_ACTIONS", "true")
	os.Setenv("GITHUB_SERVER_URL", "https://github.com")
	os.Setenv("GITHUB_REPOSITORY", "soroworks/soroforge")
	os.Unsetenv("GITHUB_RUN_ID")

	url := ciRunURL()
	if url != "" {
		t.Errorf("ciRunURL() = %q with incomplete env, want empty string", url)
	}
}
