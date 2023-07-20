import axios from 'axios';

export default function handler(request, response) {
  // api/[name].js -> /api/lee
  // request.query.name -> "lee"
  const { name } = request.query;

  axios.get('https://jsonplaceholder.typicode.com/todos/1')
    .then(res => {
        return response.status(200).end(`Hello ${name}, ${res.data}!`);
    })
    .catch(err => {
      console.error(err);
        return response.status(200).end(`Hello ${name}!, ${err.message}`);
    });
}
