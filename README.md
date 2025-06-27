# 🚀 GCP Port Forwarding Tool

A modern desktop application for easily managing port forwarding to Google Cloud Platform instances. Built with Wails, Go, and Vue.js.

## ✨ Features

- **🔐 GCP Authentication**: Seamless authentication with Google Cloud Platform
- **📁 Project Management**: Switch between different GCP projects
- **🖥️ Instance Selection**: Browse and select compute instances by zone
- **🔌 Port Forwarding**: Start/stop port forwarding with a simple interface
- **📊 Real-time Monitoring**: Track active port forwards and their status
- **🔍 Connection Testing**: Test if port forwards are working correctly
- **🎨 Modern UI**: Beautiful, responsive interface with real-time feedback

## 🛠️ Prerequisites

Before using this tool, make sure you have:

1. **Google Cloud SDK (gcloud)** installed and configured
2. **Go 1.23+** installed
3. **Node.js 16+** and **npm** installed
4. **Wails CLI** installed (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

## 🚀 Installation

1. **Clone the repository**:

   ```bash
   git clone <repository-url>
   cd gcp-tools
   ```

2. **Install dependencies**:

   ```bash
   # Install Go dependencies
   go mod tidy

   # Install frontend dependencies
   cd frontend
   npm install
   cd ..
   ```

3. **Build the application**:
   ```bash
   wails build
   ```

## 🎯 Usage

### First Time Setup

1. **Authenticate with GCP**:

   - Launch the application
   - Click "Authenticate with GCP"
   - Follow the browser authentication flow

2. **Select Project**:

   - Choose your active GCP project from the dropdown
   - The tool will automatically load available zones

3. **Configure Port Forwarding**:
   - Select a zone and compute instance
   - Set local and remote ports
   - Click "Start Port Forward"

### Port Forwarding

The tool uses `gcloud compute ssh` with IAP tunneling to establish secure port forwards:

```bash
gcloud compute ssh --zone=<zone> <instance> --project=<project> --tunnel-through-iap --ssh-flag=-L <local-port>:localhost:<remote-port> --ssh-flag=-N --ssh-flag=-f
```

### Features

- **Start Port Forward**: Creates a secure tunnel from your local machine to the GCP instance
- **Stop Port Forward**: Terminates the port forwarding connection
- **Test Connection**: Verifies if the port forward is working
- **Monitor Status**: Real-time status of all active port forwards

## 🏗️ Development

### Live Development

```bash
wails dev
```

This starts the development server with hot reload for both frontend and backend changes.

### Project Structure

```
gcp-tools/
├── app.go              # Main Go application logic
├── main.go             # Wails application entry point
├── frontend/           # Vue.js frontend
│   ├── src/
│   │   ├── App.vue     # Main application component
│   │   └── ...
│   └── ...
├── wails.json          # Wails configuration
└── go.mod              # Go module dependencies
```

### Key Components

#### Backend (Go)

- **Authentication**: GCP authentication management
- **Project Management**: GCP project switching and listing
- **Instance Discovery**: Automatic zone and instance discovery
- **Port Forwarding**: SSH tunnel management using gcloud CLI

#### Frontend (Vue.js)

- **Modern UI**: Responsive design with real-time updates
- **State Management**: Reactive data management with Vue 3 Composition API
- **Error Handling**: Comprehensive error handling and user feedback
- **Status Monitoring**: Real-time status updates and notifications

## 🔧 Configuration

### Environment Variables

The tool uses the following environment variables (optional):

- `GOOGLE_APPLICATION_CREDENTIALS`: Path to service account key file
- `GCP_PROJECT`: Default GCP project ID

### gcloud Configuration

Make sure your gcloud CLI is properly configured:

```bash
# Check current configuration
gcloud config list

# Set default project
gcloud config set project <your-project-id>

# Authenticate (if needed)
gcloud auth login
```

## 🚨 Troubleshooting

### Common Issues

1. **Authentication Failed**:

   - Ensure gcloud CLI is installed and authenticated
   - Run `gcloud auth login` in terminal
   - Check if you have proper GCP permissions

2. **No Instances Found**:

   - Verify the selected project has compute instances
   - Check if the zone contains running instances
   - Ensure you have compute.instance.list permission

3. **Port Forward Fails**:

   - Check if the local port is already in use
   - Verify the remote port is open on the instance
   - Ensure the instance allows SSH connections

4. **Permission Denied**:
   - Verify your GCP account has necessary IAM roles:
     - `compute.instance.list`
     - `compute.instances.get`
     - `compute.instances.osLogin`

### Debug Mode

Enable debug logging by setting the environment variable:

```bash
export DEBUG=true
wails dev
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Wails](https://wails.io/) - For the amazing desktop app framework
- [Vue.js](https://vuejs.org/) - For the reactive frontend framework
- [Google Cloud Platform](https://cloud.google.com/) - For the cloud infrastructure
