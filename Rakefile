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
  puts `ruby tests/functional/*.rb`
end

task :fmt do
  puts `go fmt`
  puts `go fmt reader/*.go`
  puts `go fmt view/*.go`
end
