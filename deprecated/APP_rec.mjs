import { createServer } from 'node:http';
import { appendFile } from 'node:fs/promises';
import { format } from 'node:util';
import { createRequire } from 'node:module';
const require = createRequire(import.meta.url);
const Math = require("C:/Users/Renat/AppData/Roaming/npm/node_modules/mathjs")
function getRandom(min, max) { return Math.random() * (max - min) + min; }

// @todo get from ENV
const hostname = '127.0.0.1'; // '192.168.8.242';
const port = 3001;
const dirname = 'app1/'; // '/home/user/'; 

const delay = 3000; // getRandom(500, 5000);

function loglog(d) {
  console.log(d)
  appendFile(dirname + 'debug.log', format(d) + '\n');
};


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
