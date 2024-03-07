module.exports = {
  content: [
    'internal/**/*.txt',
    'internal/**/**/*.txt',
    "internal/public/*.js",  
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}