pipeline {
	agent any
	
	stages {
		stage("Build") {
			steps {
				sh '''
					echo Build
				'''
			}
		}
		stage("Test") {
			steps {
				sh 'echo TEST'
			}
		}
		stage('Quality Control') {
            environment {
                scannerHome = tool 'sonar'
            }
            steps {
                withSonarQubeEnv('GCP Sonarqube') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
                timeout(time: 10, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }
		stage("PUSH") {
			steps {
				sh 'echo PUSH'
			}
		}
	}
}