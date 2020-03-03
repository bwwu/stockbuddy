import urllib.request
import re
import json
from bs4 import BeautifulSoup

html_doc = urllib.request.urlopen('https://www.macrotrends.net/stocks/charts/INTC/intel/balance-sheet?freq=A').read()
soup = BeautifulSoup(html_doc, 'html.parser')

inContentArea = False
jsonStr = ''
for row in soup.prettify().splitlines():
    if 'originalData' in row:
        inContentArea = True
        row = row.replace('var originalData = ', '')
    if inContentArea:
        bail = False
        if ';' in row:
            bail = True
            row = row.replace(';', '')
        jsonStr = jsonStr + row
        if bail:
            break;

obj = json.loads(jsonStr)
for row in obj:
    #print(row['field_name'])
    print(re.sub('<[^<]+?>', '', row['field_name']))
    for k in row:
        if k != 'field_name' or k != 'popup_icon':
            print(k + ',' + row[k])
