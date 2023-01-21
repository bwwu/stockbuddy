import json
import re
import sqlite3
import sys
import urllib.request

from bs4 import BeautifulSoup

def getFinancials(conn, symbol):
    print('Starting on (' + symbol + ')')
    sql = '''INSERT INTO BalanceSheet(symbol, date, field, val) VALUES(?,?,?,?);'''
    cur = conn.cursor()

    html_doc = urllib.request.urlopen('https://www.macrotrends.net/stocks/charts/' + symbol + '/intel/balance-sheet?freq=A').read()
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
        field = re.sub('<[^<]+?>', '', row['field_name'])
        for k in row:
            if k != 'field_name' and k != 'popup_icon':
                val = row[k]
                try:
                    val = float(row[k])
                except:
                    pass
                tup = (symbol, k, field, val)
                try:
                    cur.execute(sql,tup)
                except sqlite3.OperationalError as e:
                    print(e)
    conn.commit()
    cur.close()
    print('Done with (' + symbol + ')')


try:
    f = open('sp500.txt', 'r')
    conn = sqlite3.connect('/home/bwu/stock/financials.db')

except:
    print(sys.exec_info()[0])
    quit()

for symbol in f.readlines():
    getFinancials(conn, symbol.rstrip())

conn.close()
f.close()
