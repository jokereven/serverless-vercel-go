const http = require('http');

function handler(request, response) {
  const { name } = request.query;

  const options = {
    hostname: 'jsonplaceholder.typicode.com',
    path: '/todos/1',
    method: 'GET'
  };

  const req = http.request(options, res => {
    let data = '';

    res.on('data', chunk => {
      data += chunk;
    });

    res.on('end', () => {
      return response.status(200).end(`Hello ${name}, ${data}!`);
    });
  });

  req.on('error', error => {
    console.error(error);
    return response.status(500).end('Internal Server Error');
  });

  req.end();
}

module.exports = handler;
