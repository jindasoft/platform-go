init:
	@echo "Initializing..."
	@pre-commit install
	@go mod tidy

vuln:
	@echo "Checking for vulnerabilities..."
	@govulncheck ./...

vuln-v:
	@echo "Checking for vulnerabilities (verbose)..."
	@govulncheck -show=verbose ./...

sec:
	@echo "Checking for security issues..."
	@gosec ./...

check:
	@echo "Pre-commit check..."
	@pre-commit run --all-files

trivy:
	@echo "Running Trivy scan..."
	@trivy fs \
		--db-repository ghcr.io/aquasecurity/trivy-db \
		--timeout 15m \
		--exit-code 1 \
		--severity HIGH,CRITICAL \
		--skip-files configs/secret.json .
