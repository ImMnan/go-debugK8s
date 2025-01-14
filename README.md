# go-debugK8s

### This is a Go based image, to record cluster nodes capacity and usage for troubleshooting.

- Use the yaml file in /kubernetes directory
- Extract logs once the testing has been done
```sh
kubectl logs go-debug -n <namespace_name>
```
- Redirect the logs to a file

Use it for further troubleshooting cluster nodes. 

