task :default => [:test, :build]

task :build do
  puts `go build && echo "Build complete"`
end

task :install do
  puts `go install`
end

task :test do
  puts `go test reader/*.go`
  puts `go test blog/*.go`
end

task :selenium do
  pid = spawn("go run main.go -b blogs -p 2002")
  Dir.glob("tests/functional/*.rb").each do |file|
    puts `ruby #{file}`
  end
  Process.kill 0, pid
end

task :fmt do
  puts `go fmt`
  puts `go fmt blog/*.go`
  puts `go fmt reader/*.go`
  puts `go fmt view/*.go`
  puts `go fmt router/*.go`
end
