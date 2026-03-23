'use strict';

const { describe, it, before, after } = require('node:test');
const assert = require('node:assert');
const { createServer, HAPPY_MESSAGES, getRandomHappyMessage } = require('../src/server');

describe('uuhappy server', () => {
  let server;
  let port;

  before(() => new Promise((resolve) => {
    server = createServer();
    server.listen(0, () => {
      port = server.address().port;
      resolve();
    });
  }));

  after(() => new Promise((resolve) => {
    server.close(resolve);
  }));

  async function get(path) {
    const res = await fetch(`http://localhost:${port}${path}`);
    const body = await res.json();
    return { status: res.status, body };
  }

  describe('GET /', () => {
    it('returns 200 with happy status', async () => {
      const { status, body } = await get('/');
      assert.strictEqual(status, 200);
      assert.strictEqual(body.status, 'happy');
      assert.ok(body.message.includes('uuhappy'));
    });
  });

  describe('GET /happy', () => {
    it('returns 200 with a happy message', async () => {
      const { status, body } = await get('/happy');
      assert.strictEqual(status, 200);
      assert.strictEqual(body.status, 'happy');
      assert.ok(HAPPY_MESSAGES.includes(body.message));
    });
  });

  describe('GET /health', () => {
    it('returns 200 with ok status', async () => {
      const { status, body } = await get('/health');
      assert.strictEqual(status, 200);
      assert.strictEqual(body.status, 'ok');
    });
  });

  describe('GET /unknown', () => {
    it('returns 404', async () => {
      const { status, body } = await get('/unknown');
      assert.strictEqual(status, 404);
      assert.strictEqual(body.status, 'not found');
    });
  });

  describe('getRandomHappyMessage', () => {
    it('returns a message from the HAPPY_MESSAGES list', () => {
      const message = getRandomHappyMessage();
      assert.ok(HAPPY_MESSAGES.includes(message));
    });
  });
});
