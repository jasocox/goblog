require "./tests/selenium"

begin
  $driver.navigate.to "http://localhost:2002/blogs"

  blog_list = $driver.find_element(:class, "blog_list")
  titles = blog_list.find_elements(:tag_name, "h1")
  if "Example 4" != titles[0].text
    raise "Didn't get Example 4"
  end
  if "Example 3" != titles[1].text
    raise "Didn't get Example 3"
  end
  if "Example 2" != titles[2].text
    raise "Didn't get Example 2"
  end

  blog_links = $driver.find_element(:class, "blog_links")
  links = blog_links.find_elements(:tag_name, "a")
  if "Example 1" != links[0].text
    raise "Didn't get link to example"
  end
rescue Exception => e
  puts "Tests fails!"
  puts e
end

$driver.quit
