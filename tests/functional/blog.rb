require "./tests/selenium"

begin
  $driver.navigate.to "http://localhost:2002/blogs/example_1"

  title = $driver.find_element(:tag_name, "h1")
  if "Example 1" != title.text
    raise "Got #{title.text} instead of Example 1"
  end

  intro = $driver.find_element(:tag_name, "p")
  if "This is my awesome intro" != intro.text
    raise "Got #{intro.text} instead of This is my awesome intro"
  end
rescue Exception => e
  puts "Tests failed!"
  puts e
end

$driver.quit
