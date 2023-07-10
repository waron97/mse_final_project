/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['./cmd/**/*.{js,jsx,ts,tsx}'],
    theme: {
        extend: {},
    },
    plugins: [require('@tailwindcss/line-clamp')],
}
