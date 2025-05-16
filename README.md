# Team ToDo Page – University Project
A collaborative ToDo web application that enables users to create and manage teams, assign tasks, and track progress via a team-based point scoreboard. Built as part of a university project by a team of 3 programmers and 2 designers, the application showcases modern full-stack development practices with scalable cloud infrastructure.

# 🚀 Features
### ✅ Team Creation & Management
Create teams, invite members, and assign tasks to specific team groups.

### 🧩 Task Assignment
Create, edit, and delete tasks with deadlines, descriptions, and priorities.

### 🗓️ Calendar
Track task deadlines with built-in calendar.

### 🏆 Point-Based Scoreboard
Automatically assigns points to teams based on task completions. View live rankings.

### 💾 Persistent Cloud Storage
All user and task data are stored in Google Cloud SQL for reliability and performance.

### 🔄 Data Pipelines
Utilizes Google Cloud Dataflow for real-time task and point analytics.

### ☁️ Kubernetes Deployment
Deployed on Kubernetes clusters for containerized scaling and efficient resource usage.

# 🛠️ Tech Stack
## Frontend
Vue 3 – Reactive UI framework

Tailwind CSS – Utility-first CSS for custom styling

Vite – Build tool

## Backend
Go (Golang) – REST API with high performance

## Cloud Infrastructure
Google Cloud SQL – PostgreSQL instance for relational data

Google Cloud Dataflow – Stream and batch processing

Google Kubernetes Engine (GKE) – For scalable app deployment

# 📦 Installation
### Prerequisites
Docker

Docker Compose

Google Cloud SDK (optional, for cloud deployment)

- Clone the repository
- Download a docker
- Enter your files location (e.g. in powershell) ```cd (path}```
- Build the project ```docker-compose up --build```
- Go to your browser and visit http://localhost:5173/
