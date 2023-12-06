pipeline {
    agent any
    
    tools {
        go '1.21.5'
        maven '3.9.6'
    }

    environment {
        MYAPP_PATH = '' 
    }


    stages {
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
   

        stage('Build and Publish to Nexus') {
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
