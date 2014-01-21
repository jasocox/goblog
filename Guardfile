guard 'shell' do
  watch(/(.*)_test.go/) {|m|
    `go test #{m}`
  }
  watch(/(.*).go/) {
    `rake`
  }
end
