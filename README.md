# Go Web Application

This is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.

## Running the server

To run the server, execute the following command:

```bash
go run main.go
```

The server will start on port 8080. You can access it by navigating to `http://localhost:8080/courses` in your web browser.

The root URL `/` redirects to `/home`. The `/courses`, `/about`, and `/contact` pages are also available.

## Access through Kubernetes Ingress

The [Ingress](k8s/manifests/ingress.yaml) routes `go-web-app.local` to the `go-web-app` Service on port 80, which forwards to the application on port 8080. It preserves request paths; do not add a rewrite to `/`.

An NGINX ingress controller serving the `nginx` IngressClass must be installed and reachable from your machine.

Rebuild and push the application with a new image tag so the root redirect is included:

```bash
docker build -t mkurukuri/go-web-app:v2 .
docker push mkurukuri/go-web-app:v2
```

Update the image in [deployment.yaml](k8s/manifests/deployment.yaml) to `mkurukuri/go-web-app:v2` (or your own registry and new tag), then apply the manifests:

```bash
kubectl apply -f k8s/manifests/
kubectl rollout status deployment/go-web-app
kubectl get ingress go-web-app
```

Map `go-web-app.local` to the reachable ingress controller IP in your machine's `/etc/hosts` file:

```text
<INGRESS_CONTROLLER_IP> go-web-app.local
```

Use the ingress controller's IP, not the application's ClusterIP. For a local controller publishing port 80 on localhost, use `127.0.0.1`. The Ingress host rule does not create a DNS record.

Open `http://go-web-app.local`; it redirects to `/home`. Verify the redirect and another page with:

```bash
curl -I http://go-web-app.local/
curl -I http://go-web-app.local/courses
```

Expect `302` with `Location: /home` for `/`, and `200` for `/courses`.

## Looks like this

![Website](static/images/golang-website.png)

