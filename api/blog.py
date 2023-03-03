import requests
import urllib3
import xmltodict
import json

from http.server import BaseHTTPRequestHandler

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
    'Accept': '*/*',
    'Accept-Language': 'zh-CN,zh;q=0.9,en;q=0.8',
}


def blog():
    url = 'https://blog.code520.com.cn/search.xml'
    response = requests.get(
        url, headers=headers, verify=False)
    if response.status_code == 200:
        xp = xmltodict.parse(response.text)
        data = json.dumps(xp)
        return data


class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(blog())
        return
