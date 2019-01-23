@Library('jenkinsfile_stdlib') _

yproperties() // Sets releng approved global properties (SCM polling, build log rotation, etc)

emailResult(['operations@yelp.com']) {
    ystage('test') {
        node('xenial') {
            clone('packages/tplan')
            if (sh(script: 'make --dry-run test', returnStatus: true)) {
                echo 'Skipping `make test` as target does not exist.'
            } else {
                sh 'make test'
            }
        }
    }

    // Runs `make itest_${version}` and attempts to upload to apt server if not a automatically timed run
    debItestUpload('packages/tplan', ['xenial'])
}