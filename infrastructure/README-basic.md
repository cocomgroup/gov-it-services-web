### Installation
## Step 3: Deploy AWS Infrastructure

### 3.1 Validate CloudFormation Template

```bash
aws cloudformation validate-template \
  --template-body file://webapp-simple.yaml
```

### 3.2 Deploy Stack

```
cd  infrastructure

.\deploy-stack-simple.ps1
EnvironmentName -> dev
InstanceType -> t3.micro
KeyName -> my-key
SSHLocation -> 0.0.0.0/0
region -> us-east-1

```

### 3.3 Wait for Completion

```bash
aws cloudformation wait stack-create-complete --stack-name my-webapp
```

This takes 10-15 minutes. Monitor progress:
```bash
aws cloudformation describe-stack-events --stack-name my-webapp
```

### 3.4 Get Stack Outputs

```bash
aws cloudformation describe-stacks \
  --stack-name my-webapp \
  --query 'Stacks[0].Outputs'
```

#### Backend Setup
```bash
cd backend
go mod download
go run main.go
```
Backend runs on `http://localhost:8080`

#### Frontend Setup
```bash
cd frontend
npm install
npm run dev
```
Frontend runs on `http://localhost:3000`