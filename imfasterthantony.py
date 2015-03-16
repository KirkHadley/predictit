import urllib2
import json 
def get_data(contract_id):
    base = 'https://www.predictit.org/Home/GetChartPriceData?contractId='
    data = urllib2.urlopen(base + str(contract_id)).read()
    d = json.loads(data)
    return d

print get_data(439)


