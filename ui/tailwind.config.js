module.exports = {
    darkMode: 'class', // Enables dark mode based on the class applied
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
        "./fonts/**/*.ttf",
        "./fonts/**/*.ttf"
    ],
    theme: {
        extend: {
            fontFamily: {
                'logo': ['CarterOne'],
                'comic': ['Best Friends']
            },
            dropShadow: {
                'comic': [
                    '-2px -2px 0 black', '2px -2px 0 black', '-2px 2px 0 black', '2px 2px 0 black',
                    '-2px 0px 0 black', '2px 0px 0 black', '0px 2px 0 black', '0px -2px 0 black'
                ],
                'comic2': [
                    '1px 0px 1px #CCCCCC',
                    '0px 1px 1px #EEEEEE',
                    '2px 1px 1px #CCCCCC',
                    '1px 2px 1px #EEEEEE',
                    '3px 2px 1px #CCCCCC',
                    '2px 3px 1px #EEEEEE',
                    '4px 3px 1px #CCCCCC',
                    '3px 4px 1px #EEEEEE',
                    '5px 4px 1px #CCCCCC',
                    '4px 5px 1px #EEEEEE',
                    '6px 5px 1px #CCCCCC',
                    '5px 6px 1px #EEEEEE',
                    '7px 6px 1px #CCCCCC'
                ]
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