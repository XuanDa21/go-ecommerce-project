pipeline {
    agent any
    
    tools {
        go '1.21.5'
        maven '3.9.6'
    }
    
    environment {
        DOCKER_IMAGE = 'myapp:latest'
        DOCKERHUB_CREDENTIALS = credentials('dockerhub')
        registryCredentials =  'login-nexus-test'
        NEXUS_REGISTRY = 'http://192.168.86.129:8123'
    }
    
    stages {
        stage("Env Variables") {
            steps {
                sh "printenv"
            }
        }
        
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], userRemoteConfigs: [[url: 'https://github.com/XuanDa21/go-ecommerce-project.git']]])
            }
        }

        stage('Preparation') {
            steps {
                sh 'go version'
            }
        }
        
        stage('Login to dockerhub') {
            steps {
                script {
                    sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                }   
            }
        }
        
        
        stage('Build an image and push it to dockerhub') {
            steps {
                script {
                    docker.build("datrinh/ecommerce:${env.BUILD_NUMBER}")
                    
                    sh "docker push datrinh/ecommerce:${env.BUILD_NUMBER}"
                    sh "docker push datrinh/ecommerce:latest"
                    
                }   
            }
        }
        
        stage('Login to Nexus') {
            steps {
                script {
                    sh "echo 1 | docker login -u admin --password-stdin  http://192.168.86.129:8123"
                }
            }
        }
        
        stage('Push an image to Nexus') {
            steps {
                script {
                    // Tag the Docker image for Nexus
                    sh " docker tag datrinh/ecommerce  192.168.86.129:8123/datrinh/ecommerce:${env.BUILD_NUMBER}"

                    // Push the Docker image to Nexus
                    sh " docker push 192.168.86.129:8123/datrinh/ecommerce:${env.BUILD_NUMBER}"
                    sh " docker push 192.168.86.129:8123/datrinh/ecommerce:latest"
                }
            }
        }
        
        stage('Publish artifact to Nexus') {
            steps {
                script {
                    def nexusUrl = '192.168.86.129:8081'
                    def artifactName = 'ecommerce-app'
                    

                    sh 'go build -o myapp'
                    def artifactPath = 'myapp'
                    
                    nexusArtifactUploader(
                        nexusVersion: 'nexus3',
                        protocol: 'http',
                        nexusUrl: nexusUrl,
                        version: '1.0',
                        repository: 'ecommerce-app',
                        credentialsId: 'nexus-credentials',
                        packaging: 'exe',
                        artifacts: [
                            [artifactId: 'ecommerce-app', file: artifactPath]
                        ]
                    )
                }
            }
        }
    }
}

