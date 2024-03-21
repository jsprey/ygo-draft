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
            },
            backgroundColor: {
                'ygo-light':'#f8f9fa',
                'ygo-dark':'#212529',
                'ygo-card-viewer':'#3f3f46',
            }
        },
    },
    plugins: []
}