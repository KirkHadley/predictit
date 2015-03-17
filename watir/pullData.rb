require 'watir-webdriver'

def driver
  get_data

  return true
end

def get_data
  # open firefox and navigate to page
  b = Watir::Browser.new

  b.goto('https://www.predictit.org/Home/SingleOption?contractId=439#sthash.pwqtuqwo.x7PE3O2Y.dpbs')

  b.div(:id => 'buyChart').wait_until_present

  # get table data
  puts "30 days\n" + b.div(:id =>'buyChart').div.div.div.div.inner_html

  b.close
end

driver
