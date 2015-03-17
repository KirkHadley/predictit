import urllib2
import json 
def get_data(contract_id, days=90):
    base = 'https://www.predictit.org/Home/GetChartPriceData?contractId='
    suffix = '&days='
    data = urllib2.urlopen(base + str(contract_id) + suffix + str(days)).read()
    d = json.loads(data)
    return d

print get_data(439)


