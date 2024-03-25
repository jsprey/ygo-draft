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
            colors: {
                'ygo-table-header-text': '#64748b',
                'ygo-table-header-text-dark': '#ccd9ee'
            },
            backgroundColor: {
                'ygo-light': '#f8f9fa',
                'ygo-dark': '#212529',
                'ygo-card-viewer': '#3f3f46',
                'ygo-success': '#198754',
                'ygo-success-hover': '#167e4e',
                'ygo-success-active': '#117346',
                'ygo-success-disabled': '#506258',
                'ygo-danger': '#dc3545',
                'ygo-table-header': '#f8fafc',
                'ygo-table-header-dark': '#77797a',
                'ygo-table': '#f8f9fa',
                'ygo-table-dark': '#212529',
            }
        },
    },
    plugins: []
}