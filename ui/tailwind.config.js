module.exports = {
    darkMode: 'class', // Enables dark mode based on the class applied
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
        "fonts/**/*.ttf"
    ],
    theme: {
        extend: {
            fontFamily: {
                'logo': ['CarterOne']
            }
        },
    },
    plugins: []
}