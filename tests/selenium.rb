require "selenium-webdriver"

$driver = Selenium::WebDriver.for :firefox
$wait = Selenium::WebDriver::Wait.new(:timeout => 3)
