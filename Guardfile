guard 'shell' do
  watch(/(\.*)\.go$/) do |m|
    puts "Starting Tests"
    `rake`
  end
end
