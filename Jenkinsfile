#!/usr/bin/env groovy
// The above line is used to trigger correct syntax highlighting.

pipeline {
    // Lets Jenkins use Docker for us later.
    agent any    

    tools {
        go 'go 1.19'
    }
     environment {
        GO111MODULE = 'on'
    }
    stages {
//        stage ('verify tooling'){
            
//             steps {
//                 sh '''
//                 docker version
//                 docker info 
//                 docker compose version
//                 curl --version
//             '''
//                 }
//        }
        
       stage ('prune docker data'){
            
            steps {
                sh 'docker system prune -a --volumes -f'
                }
       }
        
//        stage ('migrate database'){
            
//             steps {
//                 sh 'docker-compose up -d db'
//                 }
//        }
        
       stage('docker run') {
            steps {
                sh 'docker-compose up -d'
                 sh 'docker ps'
            }
        }
    }
    
    // If anything fails, the whole Pipeline stops.
//     stages {
//         stage('Build & Test') {   
//             // Use golang.
//             agent { docker { image 'golang' } }

//             steps {                                           
//                 // Create our project directory.
//                 sh 'cd ${GOPATH}/src'
//                 sh 'mkdir -p ${GOPATH}/src/MY_PROJECT_DIRECTORY'

//                 // Copy all files in our Jenkins workspace to our project directory.                
//                 sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/MY_PROJECT_DIRECTORY'

//                 // Copy all files in our "vendor" folder to our "src" folder.
//                 sh 'cp -r ${WORKSPACE}/vendor/* ${GOPATH}/src'

//                 // Build the app.
//                 sh 'go build'               
//             }            
//         }

//         stage('Test') {
//             // Use golang.
//             agent { docker { image 'golang' } }

//             steps {                 
//                 // Create our project directory.
//                 sh 'cd ${GOPATH}/src'
//                 sh 'mkdir -p ${GOPATH}/src/MY_PROJECT_DIRECTORY'

//                 // Copy all files in our Jenkins workspace to our project directory.                
//                 sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/MY_PROJECT_DIRECTORY'

//                 // Copy all files in our "vendor" folder to our "src" folder.
//                 sh 'cp -r ${WORKSPACE}/vendor/* ${GOPATH}/src'

//                 // Remove cached test results.
//                 sh 'go clean -cache'

//                 // Run Unit Tests.
//                 sh 'go test ./... -v -short'            
//             }
//         }      

//         stage('Docker') {         
//             environment {
//                 // Extract the username and password of our credentials into "DOCKER_CREDENTIALS_USR" and "DOCKER_CREDENTIALS_PSW".
//                 // (NOTE 1: DOCKER_CREDENTIALS will be set to "your_username:your_password".)
//                 // The new variables will always be YOUR_VARIABLE_NAME + _USR and _PSW.
//                 // (NOTE 2: You can't print credentials in the pipeline for security reasons.)
//                 DOCKER_CREDENTIALS = credentials('my-docker-credentials-id')
//             }

//             steps {                           
//                 // Use a scripted pipeline.
//                 script {
//                     node {
//                         def app

//                         stage('Clone repository') {
//                             checkout scm
//                         }

//                         stage('Build image') {                            
//                             app = docker.build("${env.DOCKER_CREDENTIALS_USR}/my-project-img")
//                         }

//                         stage('Push image') {  
//                             // Use the Credential ID of the Docker Hub Credentials we added to Jenkins.
//                             docker.withRegistry('https://registry.hub.docker.com', 'my-docker-credentials-id') {                                
//                                 // Push image and tag it with our build number for versioning purposes.
//                                 app.push("${env.BUILD_NUMBER}")                      

//                                 // Push the same image and tag it as the latest version (appears at the top of our version list).
//                                 app.push("latest")
//                             }
//                         }              
//                     }                 
//                 }
//             }
//         }
//     }

//     post {
//         always {
//             // Clean up our workspace.
//             deleteDir()
//         }
//     }
}   
