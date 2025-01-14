# go-debugK8s

### This is a Go based image, to record cluster nodes capacity and usage for troubleshooting.

- Use the yaml file in /kubernetes directory
```sh
kubectl apply -f kubernetes/go-debug.yaml
```
- Extract logs once the testing has been done
```sh
kubectl logs go-debug -n <namespace_name>
```
- You can also redirect the logs to a file 


