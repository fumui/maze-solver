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
            timeout(time: 1, unit: 'HOURS') { // Just in case something goes wrong, pipeline will be killed after a timeout
                def qg = waitForQualityGate() // Reuse taskId previously collected by withSonarQubeEnv
                if (qg.status != 'OK') {
                    error "Pipeline aborted due to quality gate failure: ${qg.status}"
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