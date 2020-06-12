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
		stage("Quality Control") {
            steps {
                withSonarQubeEnv('GCP Sonarqube') { // If you have configured more than one global server connection, you can specify its name
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