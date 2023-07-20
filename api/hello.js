import axios from 'axios';

export default function handler(request, response) {
  // api/[name].js -> /api/lee
  // request.query.name -> "lee"
  const { name } = request.query;

  axios.get('https://jsonplaceholder.typicode.com/todos/1')
    .then(res => {
      console.log(res.data);
      return response.end(`Hello ${name}!`);
    })
    .catch(err => {
      console.error(err);
      return response.end(`Hello ${name}!`);
    });
}
