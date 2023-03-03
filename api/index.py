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

blog_tree = {}


class handler(BaseHTTPRequestHandler):

    def do_GET(self):
        url = 'https://blog.code520.com.cn/search.xml'
        response = requests.get(url)
        if response.status_code == 200:
            xp = xmltodict.parse(response.text)
            data = json.dumps(xp)
            obj = json.loads(data)
            search = obj.get("search")
            for key, value in search.items():
                for item in value:
                    for k, v in list(item.items()):
                        if k == "title" or k == "url":
                            title = item.get("title")
                            path = item.get("url")
                            blog_tree[title] = path

            self.send_response(200)
            self.send_header(headers)
            self.end_headers()
            self.wfile.write(data.encode())
        else:
            self.send_response(500)
            self.end_headers()
