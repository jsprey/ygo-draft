import React, {useEffect, useState} from "react";

export default function useLocalState<T>(key:string, defaultValue: T): [T, React.Dispatch<React.SetStateAction<T>>, () => void] {
    const initialize = (): T => {
        const valueRaw = localStorage.getItem(key)
        if (!valueRaw) {
            return defaultValue
        }

        try {
            return JSON.parse(valueRaw) as T;
        } catch (e) {
            localStorage.removeItem(key)
            return defaultValue
        }
    };

    const [value, setValue] = useState(initialize)
    useEffect(() => {
        localStorage.setItem(key, JSON.stringify(value))
    }, [value]);

    const resetStorage = (): void => {
        localStorage.removeItem(key)
        setValue(defaultValue)
        localStorage.setItem(key, JSON.stringify(defaultValue))
    }

    return [value, setValue, resetStorage]
}

export function removeStorageKeys(prefix: string) {
    const deleteKeys = []
    for (let localStorageKey in localStorage) {
        if (localStorageKey.startsWith(prefix)){
            deleteKeys.push(localStorageKey)
        }
    }

    deleteKeys.forEach(key => {
        localStorage.removeItem(key)
    })
}