# 📚 GCP Port Forwarding Tool Examples

This document provides examples of common use cases for the GCP Port Forwarding Tool.

## 🔧 Common Scenarios

### 1. Web Application Development

**Scenario**: You have a web application running on a GCP compute instance and want to access it locally for development.

**Configuration**:

- **Local Port**: 8080
- **Remote Port**: 80 (or 3000 for Node.js apps)
- **Instance**: `web-app-instance`
- **Zone**: `us-central1-a`

**Usage**:

1. Select your GCP project
2. Choose zone `us-central1-a`
3. Select instance `web-app-instance`
4. Set local port to 8080, remote port to 80
5. Click "Start Port Forward"
6. Access your app at `http://localhost:8080`

### 2. Database Access

**Scenario**: You need to connect to a database running on a GCP instance using a local database client.

**Configuration**:

- **Local Port**: 5432
- **Remote Port**: 5432 (PostgreSQL default)
- **Instance**: `database-instance`
- **Zone**: `us-west1-b`

**Usage**:

1. Configure the port forward as above
2. Use your local database client to connect to `localhost:5432`
3. The connection will be tunneled to your GCP database instance

### 3. SSH Access to Private Instances

**Scenario**: You have instances in a private VPC and need SSH access through IAP.

**Configuration**:

- **Local Port**: 2222
- **Remote Port**: 22 (SSH)
- **Instance**: `private-instance`
- **Zone**: `europe-west1-c`

**Usage**:

1. Set up the port forward
2. Connect via SSH: `ssh -p 2222 user@localhost`
3. The connection is automatically tunneled through IAP

### 4. Multiple Services on Same Instance

**Scenario**: You have multiple services running on different ports on the same instance.

**Configuration**:

- **Web App**: Local 8080 → Remote 80
- **API**: Local 3000 → Remote 3000
- **Database**: Local 5432 → Remote 5432
- **Instance**: `multi-service-instance`

**Usage**:

1. Create multiple port forwards for the same instance
2. Each service will be accessible on its local port
3. Monitor all forwards in the "Active Port Forwards" section

## 🛠️ Advanced Examples

### 5. Load Balancer Testing

**Scenario**: Test a load balancer configuration by forwarding to backend instances.

**Configuration**:

- **Instance 1**: Local 8081 → Remote 80
- **Instance 2**: Local 8082 → Remote 80
- **Instance 3**: Local 8083 → Remote 80

**Usage**:

1. Set up port forwards to each backend instance
2. Test load balancer behavior by accessing different local ports
3. Verify traffic distribution and health checks

### 6. Microservices Development

**Scenario**: Develop a microservices architecture with services on different instances.

**Configuration**:

- **User Service**: Local 3001 → Remote 3000
- **Order Service**: Local 3002 → Remote 3000
- **Payment Service**: Local 3003 → Remote 3000
- **Database**: Local 5432 → Remote 5432

**Usage**:

1. Set up port forwards for each service
2. Configure your local development environment to use localhost ports
3. Test service-to-service communication locally

## 🔍 Troubleshooting Examples

### Port Already in Use

**Error**: "Port 8080 is already in use"

**Solution**:

1. Check what's using the port: `lsof -i :8080`
2. Stop the existing process or choose a different local port
3. Use the "Stop Port Forward" button to clean up existing forwards

### Authentication Issues

**Error**: "Authentication failed"

**Solution**:

1. Run `gcloud auth login` in terminal
2. Verify your account has proper GCP permissions
3. Check if you're using the correct project

### Instance Not Found

**Error**: "No instances found in zone"

**Solution**:

1. Verify the instance is running in the selected zone
2. Check if you have `compute.instance.list` permission
3. Ensure the instance name is correct

## 📋 Best Practices

### Security

- Always use IAP tunneling (enabled by default)
- Don't expose sensitive services on public ports
- Regularly rotate SSH keys and credentials

### Performance

- Use appropriate local ports (avoid conflicts)
- Monitor resource usage on instances
- Close unused port forwards to free up resources

### Development Workflow

1. Start port forwards for all needed services
2. Test connections before starting development
3. Use the "Test Connection" feature regularly
4. Keep track of active forwards in the UI

## 🎯 Quick Reference

### Common Port Mappings

- **Web Servers**: 80, 443, 8080, 3000
- **Databases**: 5432 (PostgreSQL), 3306 (MySQL), 27017 (MongoDB)
- **SSH**: 22
- **Redis**: 6379
- **Elasticsearch**: 9200

### Useful Commands

```bash
# Check active port forwards
lsof -i -P | grep LISTEN

# Test connection to forwarded port
curl http://localhost:8080

# Check gcloud authentication
gcloud auth list

# List compute instances
gcloud compute instances list
```

### Environment Variables

```bash
# Set default project
export GCP_PROJECT=your-project-id

# Use service account (optional)
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/key.json
```
