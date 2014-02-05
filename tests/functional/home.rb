require "./tests/selenium"
$driver.navigate.to "http://localhost:2002/"

sleep 1
$driver.quit
