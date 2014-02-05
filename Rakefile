task :default => [:test, :build]

task :build do
  puts `go build && echo "Build complete"`
end

task :install do
  puts `go install`
end

task :test do
  puts `go test reader/*.go`
end

task :selenium do
  pid = spawn("go run main.go -b blogs -p 2002")
  puts `ruby tests/functional/*.rb`
  Process.kill 0, pid
end

task :fmt do
  puts `go fmt`
  puts `go fmt reader/*.go`
  puts `go fmt view/*.go`
end
