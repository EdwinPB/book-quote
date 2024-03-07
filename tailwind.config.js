module.exports = {
  content: [
    'internal/**/*.templ',
    'internal/**/**/*.templ',
    "internal/public/*.js",  
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}