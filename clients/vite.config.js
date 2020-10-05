module.exports = {
  port: 5000,
  proxy: {
    '/api': {
      target: 'http://localhost:3000',
      changeOrigin: true,
      rewrite: path => path.replace(/^\/api/, ''),
    },
  },
}
