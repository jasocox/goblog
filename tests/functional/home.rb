require "./tests/selenium"

begin
  $driver.navigate.to "http://localhost:2002/"

  element = $driver.find_element(:class => "navbar")
  links = element.find_elements(:tag_name, "a")
  if "About Me" != links[0].text
    raise "About Me link missing"
  end
  if "Contact Me" != links[1].text
    raise "Contact Me link missing"
  end
rescue Exception => e
  puts "Tests failed!"
  puts e
end

$driver.quit
