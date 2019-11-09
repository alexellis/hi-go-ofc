# OpenFaaS Cloud example

This example deploys to both Kubernetes and AWS Lambda at the same time.

Simply send a PR and when merged, both will be redeployed.

PRs must be signed off with `git commit -s` or a `Signed-off-by: Name <email>` message as the last line of your commit message.

After the merge, view the status of your build here:

* https://github.com/alexellis/hi-go-ofc/commits/master 

Invoke URLs:

* K8s Pod -> https://alexellis.ofc.plonk.dev/run-on-k8s
* AWS Lambda function -> https://alexellis.ofc.plonk.dev/run-on-lambda
