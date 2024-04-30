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
            colors: {
                'primary': {
                    'light': '#f6d5bc',
                    DEFAULT: '#e1741e',
                    'hover': '#cb681b',
                    'active': '#b45d18',
                    'dark': '#9e5115'
                },
                'secondary': {
                    'light': '#bcdcf6',
                    DEFAULT: '#1E8BE1',
                    'hover': '#1b7dcb',
                    'active': '#186fb4',
                    'dark': '#15619e'
                },
                'danger': {
                    'light': '#ffc2c2',
                    DEFAULT: '#ff3333',
                    'hover': '#e62e2e',
                    'active': '#cc2929',
                    'dark': '#b32424'
                },
                'success': {
                    'light': '#bdebc2',
                    DEFAULT: '#22BB33',
                    'hover': '#1fa82e',
                    'active': '#1b9629',
                    'dark': '#188324'
                },
                'warning': {
                    'light': '#fff0b3',
                    DEFAULT: '#FFCC00',
                    'hover': '#e6b800',
                    'active': '#cca300',
                    'dark': '#b38f00'
                },
                light: {
                    DEFAULT: '#f8f8f8',
                    1: '#dfdfdf',
                    2: '#c6c6c6',
                    3: '#aeaeae',
                },
                dark: {
                    DEFAULT: '#0d1117',
                    1: '#25292e',
                    2: '#3d4145',
                    3: '#56585d',
                },
                cardcolors: {
                    effect: '#bb6034',
                    spell: '#04947f',
                    trap: '#a31270'
                }
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