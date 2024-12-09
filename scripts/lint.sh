lint_service() {
    local service_path=$1
    echo "Linting service: $service_path"
    cd "$service_path" || exit
    golangci-lint run ./... --fix
    cd - || exit
}

lint_service "services/auth-ms"
