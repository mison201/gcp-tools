<template>
  <div id="app">
    <header class="header">
      <div class="header-content">
        <h1>🚀 GCP Port Forwarding Tool</h1>
        <p>
          Easily manage port forwarding for your Google Cloud Platform instances
        </p>
      </div>
    </header>

    <main class="main-content">
      <!-- Authentication Section -->
      <div class="card auth-section" v-if="!isAuthenticated">
        <h2>🔐 Authentication</h2>
        <p>You need to authenticate with Google Cloud Platform to continue.</p>
        <button @click="authenticate" class="btn btn-primary">
          Authenticate with GCP
        </button>
      </div>

      <!-- Main Interface -->
      <div v-else class="dashboard">
        <!-- Project Selection -->
        <div class="card">
          <h2>📁 Project Configuration</h2>
          <div class="form-group">
            <label for="project-select">Active Project:</label>
            <select
              id="project-select"
              v-model="selectedProject"
              @change="onProjectChange"
              class="form-control"
            >
              <option value="">Loading projects...</option>
              <option
                v-for="project in projects"
                :key="project.projectId"
                :value="project.projectId"
              >
                {{ project.projectName }} ({{ project.projectId }})
              </option>
            </select>
          </div>
          <div class="project-info" v-if="selectedProject">
            <p><strong>Active Project:</strong> {{ selectedProject }}</p>
          </div>
        </div>

        <!-- Instance Selection -->
        <div class="card" v-if="selectedProject">
          <h2>🖥️ Instance Configuration</h2>
          <div class="form-row">
            <div class="form-group">
              <label for="zone-select">Zone:</label>
              <select
                id="zone-select"
                v-model="selectedZone"
                @change="onZoneChange"
                class="form-control"
              >
                <option value="">Select zone...</option>
                <option v-for="zone in zones" :key="zone" :value="zone">
                  {{ zone }}
                </option>
              </select>
            </div>
            <div class="form-group">
              <label for="instance-select">Instance:</label>
              <select
                id="instance-select"
                v-model="selectedInstance"
                class="form-control"
              >
                <option value="">Select instance...</option>
                <option
                  v-for="instance in instances"
                  :key="instance"
                  :value="instance"
                >
                  {{ instance }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <!-- Port Forwarding Configuration -->
        <div class="card" v-if="selectedInstance">
          <h2>🔌 Port Forwarding</h2>
          <div class="form-row">
            <div class="form-group">
              <label for="local-port">Local Port:</label>
              <input
                id="local-port"
                v-model.number="localPort"
                type="number"
                min="1024"
                max="65535"
                class="form-control"
                placeholder="8080"
              />
            </div>
            <div class="form-group">
              <label for="remote-port">Remote Port:</label>
              <input
                id="remote-port"
                v-model.number="remotePort"
                type="number"
                min="1"
                max="65535"
                class="form-control"
                placeholder="80"
              />
            </div>
          </div>

          <div class="port-forward-actions">
            <button
              @click="startPortForward"
              class="btn btn-success"
              :disabled="!canStartPortForward"
            >
              ▶️ Start Port Forward
            </button>
            <button
              @click="stopPortForward"
              class="btn btn-danger"
              :disabled="!localPort"
            >
              ⏹️ Stop Port Forward
            </button>
            <button
              @click="testConnection"
              class="btn btn-info"
              :disabled="!localPort"
            >
              🔍 Test Connection
            </button>
          </div>
        </div>

        <!-- Active Port Forwards -->
        <div class="card">
          <h2>📊 Active Port Forwards</h2>
          <div class="port-forwards-list">
            <div v-if="activeForwards.length === 0" class="empty-state">
              <p>No active port forwards</p>
            </div>
            <div
              v-else
              class="forward-item"
              v-for="forward in activeForwards"
              :key="forward.localPort"
            >
              <div class="forward-info">
                <span class="forward-label"
                  >Local {{ forward.localPort }} → Remote
                  {{ forward.remotePort }}</span
                >
                <span class="forward-status" :class="forward.status">{{
                  forward.status
                }}</span>
              </div>
              <div class="forward-actions">
                <button
                  @click="stopSpecificForward(forward.localPort)"
                  class="btn btn-sm btn-danger"
                >
                  Stop
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Status Messages -->
    <div v-if="statusMessage" class="status-message" :class="statusType">
      {{ statusMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue"

// Reactive data
const isAuthenticated = ref(false)
const projects = ref([])
const selectedProject = ref("")
const zones = ref([])
const selectedZone = ref("")
const instances = ref([])
const selectedInstance = ref("")
const localPort = ref(8080)
const remotePort = ref(80)
const activeForwards = ref([])
const statusMessage = ref("")
const statusType = ref("info")

// Computed properties
const canStartPortForward = computed(() => {
  return (
    selectedProject.value &&
    selectedZone.value &&
    selectedInstance.value &&
    localPort.value &&
    remotePort.value
  )
})

// Methods
const showStatus = (message, type = "info") => {
  statusMessage.value = message
  statusType.value = type
  setTimeout(() => {
    statusMessage.value = ""
  }, 5000)
}

const checkAuth = async () => {
  try {
    const authResult = await window.go.main.App.CheckGCloudAuth()
    isAuthenticated.value = authResult
    if (!authResult) {
      showStatus("Please authenticate with GCP to continue", "warning")
    }
  } catch (error) {
    console.error("Auth check failed:", error)
    showStatus("Failed to check authentication status", "error")
  }
}

const authenticate = async () => {
  try {
    await window.go.main.App.AuthenticateGCloud()
    await checkAuth()
    showStatus("Authentication successful!", "success")
  } catch (error) {
    console.error("Authentication failed:", error)
    showStatus("Authentication failed. Please try again.", "error")
  }
}

const loadProjects = async () => {
  try {
    const projectList = await window.go.main.App.GetGCPProjects()
    projects.value = projectList
    if (projectList.length > 0) {
      selectedProject.value = projectList[0].projectId
    }
  } catch (error) {
    console.error("Failed to load projects:", error)
    showStatus("Failed to load GCP projects", "error")
  }
}

const onProjectChange = async () => {
  if (selectedProject.value) {
    try {
      await window.go.main.App.SetActiveProject(selectedProject.value)
      await loadZones()
      showStatus(`Switched to project: ${selectedProject.value}`, "success")
    } catch (error) {
      console.error("Failed to set project:", error)
      showStatus("Failed to set active project", "error")
    }
  }
}

const loadZones = async () => {
  if (!selectedProject.value) return

  try {
    const zoneList = await window.go.main.App.GetZones(selectedProject.value)
    zones.value = zoneList
    if (zoneList.length > 0) {
      selectedZone.value = zoneList[0]
    }
  } catch (error) {
    console.error("Failed to load zones:", error)
    showStatus("Failed to load zones", "error")
  }
}

const onZoneChange = async () => {
  if (selectedZone.value) {
    await loadInstances()
  }
}

const loadInstances = async () => {
  if (!selectedProject.value || !selectedZone.value) return

  try {
    const instanceList = await window.go.main.App.GetComputeInstances(
      selectedProject.value,
      selectedZone.value,
    )
    instances.value = instanceList
  } catch (error) {
    console.error("Failed to load instances:", error)
    showStatus("Failed to load compute instances", "error")
  }
}

const startPortForward = async () => {
  if (!canStartPortForward.value) return

  try {
    const config = {
      localPort: localPort.value,
      remotePort: remotePort.value,
      projectId: selectedProject.value,
      zone: selectedZone.value,
      instance: selectedInstance.value,
      status: "starting",
    }

    await window.go.main.App.StartPortForward(config)

    // Add to active forwards
    activeForwards.value.push(config)

    showStatus(
      `Port forward started: localhost:${localPort.value} → ${selectedInstance.value}:${remotePort.value}`,
      "success",
    )
  } catch (error) {
    console.error("Failed to start port forward:", error)
    showStatus("Failed to start port forward", "error")
  }
}

const stopPortForward = async () => {
  if (!localPort.value) return

  try {
    await window.go.main.App.StopPortForward(localPort.value)

    // Remove from active forwards
    activeForwards.value = activeForwards.value.filter(
      (f) => f.localPort !== localPort.value,
    )

    showStatus(`Port forward stopped for port ${localPort.value}`, "success")
  } catch (error) {
    console.error("Failed to stop port forward:", error)
    showStatus("Failed to stop port forward", "error")
  }
}

const stopSpecificForward = async (port) => {
  try {
    await window.go.main.App.StopPortForward(port)

    // Remove from active forwards
    activeForwards.value = activeForwards.value.filter(
      (f) => f.localPort !== port,
    )

    showStatus(`Port forward stopped for port ${port}`, "success")
  } catch (error) {
    console.error("Failed to stop port forward:", error)
    showStatus("Failed to stop port forward", "error")
  }
}

const testConnection = async () => {
  if (!localPort.value) return

  try {
    const isWorking = await window.go.main.App.TestConnection(localPort.value)
    if (isWorking) {
      showStatus(
        `Connection test successful for port ${localPort.value}`,
        "success",
      )
    } else {
      showStatus(`Connection test failed for port ${localPort.value}`, "error")
    }
  } catch (error) {
    console.error("Connection test failed:", error)
    showStatus("Connection test failed", "error")
  }
}

// Lifecycle
onMounted(async () => {
  await checkAuth()
  if (isAuthenticated.value) {
    await loadProjects()
  }
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.header {
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  padding: 2rem 0;
}

.header-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  text-align: center;
  color: white;
}

.header h1 {
  font-size: 2.5rem;
  margin-bottom: 0.5rem;
  font-weight: 700;
}

.header p {
  font-size: 1.1rem;
  opacity: 0.9;
}

.main-content {
  flex: 1;
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
  display: grid;
  gap: 2rem;
}

.card {
  background: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.card h2 {
  color: #333;
  margin-bottom: 1.5rem;
  font-size: 1.5rem;
  font-weight: 600;
}

.form-group {
  margin-bottom: 1rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #555;
  font-weight: 500;
}

.form-control {
  width: 100%;
  padding: 0.75rem;
  border: 2px solid #e1e5e9;
  border-radius: 8px;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.form-control:focus {
  outline: none;
  border-color: #667eea;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  text-decoration: none;
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-primary {
  background: #667eea;
  color: white;
}

.btn-primary:hover:not(:disabled) {
  background: #5a6fd8;
  transform: translateY(-2px);
}

.btn-success {
  background: #10b981;
  color: white;
}

.btn-success:hover:not(:disabled) {
  background: #059669;
  transform: translateY(-2px);
}

.btn-danger {
  background: #ef4444;
  color: white;
}

.btn-danger:hover:not(:disabled) {
  background: #dc2626;
  transform: translateY(-2px);
}

.btn-info {
  background: #3b82f6;
  color: white;
}

.btn-info:hover:not(:disabled) {
  background: #2563eb;
  transform: translateY(-2px);
}

.btn-sm {
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
}

.port-forward-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
  flex-wrap: wrap;
}

.project-info {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8fafc;
  border-radius: 8px;
  border-left: 4px solid #667eea;
}

.port-forwards-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.forward-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: #f8fafc;
  border-radius: 8px;
  border: 1px solid #e1e5e9;
}

.forward-info {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.forward-label {
  font-weight: 500;
  color: #333;
}

.forward-status {
  font-size: 0.875rem;
  padding: 0.25rem 0.5rem;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 500;
}

.forward-status.running {
  background: #dcfce7;
  color: #166534;
}

.forward-status.stopped {
  background: #fef2f2;
  color: #991b1b;
}

.forward-status.starting {
  background: #fef3c7;
  color: #92400e;
}

.empty-state {
  text-align: center;
  padding: 2rem;
  color: #6b7280;
}

.auth-section {
  text-align: center;
  max-width: 500px;
  margin: 0 auto;
}

.auth-section p {
  margin-bottom: 1.5rem;
  color: #6b7280;
}

.status-message {
  position: fixed;
  top: 2rem;
  right: 2rem;
  padding: 1rem 1.5rem;
  border-radius: 8px;
  color: white;
  font-weight: 500;
  z-index: 1000;
  animation: slideIn 0.3s ease;
}

.status-message.success {
  background: #10b981;
}

.status-message.error {
  background: #ef4444;
}

.status-message.warning {
  background: #f59e0b;
}

.status-message.info {
  background: #3b82f6;
}

@keyframes slideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@media (max-width: 768px) {
  .header h1 {
    font-size: 2rem;
  }

  .main-content {
    padding: 1rem;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .port-forward-actions {
    flex-direction: column;
  }

  .forward-item {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }
}
</style>
