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
