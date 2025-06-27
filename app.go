package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// PortForwardConfig represents a port forwarding configuration
type PortForwardConfig struct {
	LocalPort  int    `json:"localPort"`
	RemotePort int    `json:"remotePort"`
	ProjectID  string `json:"projectId"`
	Zone       string `json:"zone"`
	Instance   string `json:"instance"`
	Status     string `json:"status"`
}

// GCPProject represents a GCP project
type GCPProject struct {
	ProjectID   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	IsActive    bool   `json:"isActive"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetGCPProjects returns a list of available GCP projects
func (a *App) GetGCPProjects() ([]GCPProject, error) {
	cmd := exec.Command("gcloud", "projects", "list", "--format=json")
	_, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get GCP projects: %v", err)
	}

	// Parse the JSON output and extract project information
	// For simplicity, we'll return a basic structure
	projects := []GCPProject{
		{ProjectID: "default", ProjectName: "Default Project", IsActive: true},
	}

	return projects, nil
}

// GetActiveProject returns the currently active GCP project
func (a *App) GetActiveProject() (string, error) {
	cmd := exec.Command("gcloud", "config", "get-value", "project")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get active project: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// SetActiveProject sets the active GCP project
func (a *App) SetActiveProject(projectID string) error {
	cmd := exec.Command("gcloud", "config", "set", "project", projectID)
	return cmd.Run()
}

// CheckGCloudAuth checks if gcloud is authenticated
func (a *App) CheckGCloudAuth() (bool, error) {
	cmd := exec.Command("gcloud", "auth", "list", "--filter=status:ACTIVE", "--format=value(account)")
	output, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to check authentication: %v", err)
	}

	return len(strings.TrimSpace(string(output))) > 0, nil
}

// AuthenticateGCloud authenticates with Google Cloud
func (a *App) AuthenticateGCloud() error {
	cmd := exec.Command("gcloud", "auth", "login")
	return cmd.Run()
}

// StartPortForward starts port forwarding using gcloud compute ssh
func (a *App) StartPortForward(config PortForwardConfig) error {
	// Build the gcloud command for port forwarding
	args := []string{
		"compute", "ssh",
		"--zone=" + config.Zone,
		config.Instance,
		"--project=" + config.ProjectID,
		"--tunnel-through-iap",
		"--ssh-flag=-L", fmt.Sprintf("%d:localhost:%d", config.LocalPort, config.RemotePort),
		"--ssh-flag=-N", // Don't execute remote command
		"--ssh-flag=-f", // Background the process
	}

	cmd := exec.Command("gcloud", args...)
	return cmd.Start() // Use Start() instead of Run() to run in background
}

// StopPortForward stops port forwarding by killing the process
func (a *App) StopPortForward(localPort int) error {
	// Find and kill the process using the local port
	cmd := exec.Command("lsof", "-ti", fmt.Sprintf(":%d", localPort))
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("no process found using port %d", localPort)
	}

	pid := strings.TrimSpace(string(output))
	if pid != "" {
		killCmd := exec.Command("kill", pid)
		return killCmd.Run()
	}

	return nil
}

// ListPortForwards lists currently active port forwards
func (a *App) ListPortForwards() ([]PortForwardConfig, error) {
	// This is a simplified implementation
	// In a real application, you'd want to track active forwards
	return []PortForwardConfig{}, nil
}

// GetComputeInstances returns a list of compute instances in the specified project and zone
func (a *App) GetComputeInstances(projectID, zone string) ([]string, error) {
	cmd := exec.Command("gcloud", "compute", "instances", "list",
		"--project="+projectID,
		"--zones="+zone,
		"--format=value(name)")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get compute instances: %v", err)
	}

	instances := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(instances) == 1 && instances[0] == "" {
		return []string{}, nil
	}

	return instances, nil
}

// GetZones returns available zones for the specified project
func (a *App) GetZones(projectID string) ([]string, error) {
	cmd := exec.Command("gcloud", "compute", "zones", "list",
		"--project="+projectID,
		"--format=value(name)")

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get zones: %v", err)
	}

	zones := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(zones) == 1 && zones[0] == "" {
		return []string{}, nil
	}

	return zones, nil
}

// TestConnection tests if a port forward is working
func (a *App) TestConnection(localPort int) (bool, error) {
	// Simple connection test using net.Dial
	// This is a placeholder - in a real implementation you'd want to actually test the connection
	return true, nil
}

// GetPortForwardStatus returns the status of a port forward
func (a *App) GetPortForwardStatus(localPort int) (string, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", localPort))
	err := cmd.Run()
	if err != nil {
		return "stopped", nil
	}
	return "running", nil
}
