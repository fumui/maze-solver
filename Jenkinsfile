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
        stage("Quality Gate"){
            steps {
                timeout(time: 1, unit: 'HOURS') {
                    // Parameter indicates whether to set pipeline to UNSTABLE if Quality Gate fails
                    // true = set pipeline to UNSTABLE, false = don't
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