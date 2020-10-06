const path = require('path')
module.exports = {
  port: 5000,
  alias: {
    '/@/': path.resolve(__dirname, 'src'),
  },
  proxy: {
    '/api': {
      target: 'http://localhost:3000',
      changeOrigin: true,
      // rewrite: path => path.replace(/^\/api/, ''),
    },
  },
}
