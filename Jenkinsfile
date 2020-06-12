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
		stage("PUSH") {
			steps {
				sh 'echo PUSH'
			}
		}
	}
}