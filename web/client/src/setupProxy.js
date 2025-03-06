const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    '/api',
    createProxyMiddleware({
      target: 'https://doit-backend-oktavarium.amvera.io',
      changeOrigin: true,
      secure: false,
      router: {
        'mdzen.nl.tuna.am': 'https://doit-backend-oktavarium.amvera.io',
      },
      pathRewrite: {
        '^/api': '/api'
      },
      onProxyReq: (proxyReq) => {
        console.log('Proxying to:', proxyReq.path);
      },
      onError: (err, req, res) => {
        console.error('Proxy Error:', err);
      }
    })
  );
}; 