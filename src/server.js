'use strict';

const http = require('node:http');

const HAPPY_MESSAGES = [
  'uuhappy! 😊',
  'Everything is wonderful! 🌟',
  'Stay positive! ✨',
  'You are awesome! 🎉',
  'Happiness is contagious! 💛',
];

function getRandomHappyMessage() {
  return HAPPY_MESSAGES[Math.floor(Math.random() * HAPPY_MESSAGES.length)];
}

function createServer() {
  const server = http.createServer((req, res) => {
    const url = new URL(req.url, `http://${req.headers.host}`);

    if (req.method === 'GET' && url.pathname === '/') {
      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'happy', message: 'uuhappy server is running! 🎊' }));
      return;
    }

    if (req.method === 'GET' && url.pathname === '/happy') {
      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'happy', message: getRandomHappyMessage() }));
      return;
    }

    if (req.method === 'GET' && url.pathname === '/health') {
      res.writeHead(200, { 'Content-Type': 'application/json' });
      res.end(JSON.stringify({ status: 'ok' }));
      return;
    }

    res.writeHead(404, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ status: 'not found', message: 'Route not found' }));
  });

  return server;
}

module.exports = { createServer, getRandomHappyMessage, HAPPY_MESSAGES };
