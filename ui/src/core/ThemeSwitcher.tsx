import React from 'react';
import {useTheme} from "./context/ColorThemeProvider";

const ThemeSwitcher = () => {
    const {isDarkMode, toggleDarkMode} = useTheme()

    return (
        <button
            onClick={toggleDarkMode}
            className="m-1 px-2 py-2 dark:text-white"
        >
            {isDarkMode ? 'Light Mode' : 'Dark Mode'}
        </button>
    );
};

export default ThemeSwitcher;