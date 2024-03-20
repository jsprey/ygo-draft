import {Context, createContext, useContext, useMemo, useState} from "react";

const ThemeContext: Context<ThemeContextType> = createContext({} as ThemeContextType);
export type ThemeContextType = {
    isDarkMode: boolean,
    toggleDarkMode: () => void
}

// Contains the identifier to save the current state of the dark mode
const DarkModeStorageKey: string = "darkMode"

export type ColorThemeProviderProps = {
    children: JSX.Element[] | JSX.Element;
}

const ColorThemeProvider = (props: ColorThemeProviderProps) => {
    // State to hold the current theme
    const [isDarkMode, setDarkMode] = useState<boolean>(() => {
        const isDarkMode = localStorage.getItem(DarkModeStorageKey) === 'true';
        document.documentElement.classList.toggle('dark', isDarkMode);
        return isDarkMode ? isDarkMode : false;
    });

    // Function to set the current theme
    const toggleDarkMode = (): void => {
        const newMode = !isDarkMode

        setDarkMode(newMode)
        document.documentElement.classList.toggle('dark', newMode);
        localStorage.setItem(DarkModeStorageKey, String(newMode));
    };

    // Memoized value of the authentication context
    const contextValue = useMemo(
        () => ({
            isDarkMode,
            toggleDarkMode,
        }),
        [isDarkMode]
    );

    // Provide the authentication context to the children components
    return (
        <ThemeContext.Provider value={contextValue}>{props.children}</ThemeContext.Provider>
    );
};

export const useTheme = () => {
    return useContext(ThemeContext);
};

export default ColorThemeProvider;