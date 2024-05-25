import { createServer } from 'node:http';
import { appendFile } from 'node:fs/promises';
import { format } from 'node:util';

const hostname = '127.0.0.2';
const port = 3000;
const delay = 3000;

const dirname = '/home/user'; // @todo get from ENV
function loglog(d) {
  console.log(d)
  appendFile(dirname + '/debug.log', format(d) + '\n');
};

// function getRandom(min, max) { return random() * (max - min) + min; }

const server = createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  const timeoutScheduled = Date.now();
  setTimeout(() => {
    const res_delay = Date.now() - timeoutScheduled;

    res.end('Hello World!\n');
    loglog(`${res_delay} ms processing request`);
  }, delay);
});


server.listen(port, hostname, () => {
  console.log('Listening on ' + hostname + ':' + port);
});
