pipeline {
	agent any 
	stages {
		stage('build') {
			steps {
				sh 'go version'
			}
		}
	}
	post {
		always {
			echo 'This will always run'
		}
		success {
			echo 'Build succeeded'
		}
	}
}
