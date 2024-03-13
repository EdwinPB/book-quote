module.exports = {
  content: [
    'internal/components/**/*.templ',
    'internal/public/*.js',
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}