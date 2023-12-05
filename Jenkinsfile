pipeline {
    agent any

    // environment {
    //     GOPATH = '/path/to/gopath'  // Set your GOPATH here
    // }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    sh 'go build -o myapp'
                }
            }
        }

        stage('Publish to Nexus') {
            steps {
                script {
                    def nexusUrl = 'http://192.168.186.200:8081/repository/test-golang'
                    def nexusUser = 'admin'
                    def nexusPassword = '1'
                    def artifactName = 'test-golang'

                    withCredentials([usernamePassword(credentialsId: 'nexus-credentials', usernameVariable: 'NEXUS_USERNAME', passwordVariable: 'NEXUS_PASSWORD')]) {
                        sh "curl -v -u $NEXUS_USERNAME:$NEXUS_PASSWORD --upload-file ${artifactName} ${nexusUrl}/${artifactName}"
                    }
                }
            }
        }
    }
}
