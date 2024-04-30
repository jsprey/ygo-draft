import React from "react";

function Footer() {
    return <footer className="bg-light-1 dark:bg-dark-1 mt-2 text-center text-dark dark:text-light">
        {/* <!--Copyright section--> */}
        <div
            className="bg-secondary-200 p-4 text-center text-secondary-700 dark:bg-secondary-700 dark:text-secondary-200">
            © 2024 Joshua Sprey:
            <span
                className="text-secondary-800 dark:text-secondary-400"
            >YGO-Draft</span>
        </div>
    </footer>
}

export default Footer
