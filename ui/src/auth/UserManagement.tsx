import React, {useState} from "react";
import {User, useUsers} from "../api/hooks/useUsers";
import {Alert} from "react-bootstrap";
import classNames from "classnames";

function UserManagement() {
    const [currentPage, setCurrentPage] = useState<number>(0);
    const [currentPageSize] = useState<number>(15);
    const {data, isLoading, error} = useUsers(currentPage, currentPageSize);

    let tableEntries: JSX.Element[] = []
    if (error) {
        tableEntries.push(<tr>
            <td colSpan={5}><Alert variant={"danger"}>Failed to load users!</Alert></td>
        </tr>)
    } else if (isLoading) {
        tableEntries = createPlaceholderEntries()
    } else if (data) {
        tableEntries = createTableWithActualData(data.users)
    }

    function createTableWithActualData(users: User[]): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        users.forEach(user => {
            const entry = <tr key={`row-${user.id}`}
                              className={"odd:bg-gray-50 odd:dark:bg-gray-900 even:bg-gray-100 even:dark:bg-gray-800 border-b dark:border-gray-700"}>
                <td className="px-6 py-2">{user.id}</td>
                <td className="px-6 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">{user.email}</td>
                <td className="px-6 py-2">{user.display_name}</td>
                <td className="px-6 py-2">{user.is_admin ? "Admin" : "User"}</td>
                <td className="px-6 py-2"></td>
            </tr>
            tableEntries.push(entry)
        })

        return tableEntries
    }

    function createPlaceholderEntries(): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        for (let i = 0; i < 15; i++) {
            const entry = <tr key={`row-placeholder-${i}`}
                              className={"odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700"}>
                <td><span className={"placeholder w-100 px-6 py-3"}>1</span></td>
                <td><span className={"placeholder w-100 px-6 py-3"}>1</span></td>
                <td><span className={"placeholder w-100 px-6 py-3"}>1</span></td>
                <td><span className={"placeholder w-100 px-6 py-3"}>1</span></td>
                <td><span className={"placeholder w-100 px-6 py-3"}>1</span></td>
            </tr>
            tableEntries.push(entry)
        }

        return tableEntries
    }

    function getTableHeader() {
        return <thead className={"text-xs text-gray-700 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400"}>
        <tr>
            <th className="px-6 py-3 rounded-tl">ID</th>
            <th className="px-6 py-3">Email</th>
            <th className="px-6 py-3">Display Name</th>
            <th className="px-6 py-3">Administrator</th>
            <th className="px-6 py-3 rounded-tr">Actions</th>
        </tr>
        </thead>;
    }

    function getTableControls() {
        if (isLoading || error || !data) {
            return <></>
        }

        let minUser = (currentPage * currentPageSize)
        let maxUser = ((currentPage + 1) * currentPageSize)
        if (data.numberOfUsers < maxUser) {
            maxUser = data.numberOfUsers
        }

        const allButtonCN = classNames("py-2 px-4", "text-sm font-semibold", "bg-gray-300 hover:bg-gray-400 active:bg-gray-500 disabled:bg-gray-200", "text-gray-800 disabled:text-gray-300")
        return <div
            className="flex justify-content-between xs:flex-row items-center xs:justify-between">
            <button
                disabled={currentPage === 0}
                className={classNames(allButtonCN, "rounded-bl")}
                onClick={event => {
                    let newPage = currentPage - 1;
                    setCurrentPage(newPage < 0 ? 0 : newPage)
                }
                }>
                Prev
            </button>
            <span className="text-xs xs:text-sm text-gray-700 bg-gray-200 dark:bg-gray-700 dark:text-gray-400 flex-grow-1 py-2 text-center place-self-stretch">
                            Showing {minUser} to {maxUser} of {data.numberOfUsers} Users
                        </span>
            <button
                disabled={data.numberOfUsers === maxUser}
                className={classNames(allButtonCN, "rounded-br")}
                onClick={event => {
                    let newPage = currentPage + 1;
                    setCurrentPage(newPage)
                }
                }>
                Next
            </button>
        </div>
    }

    return <p className={"placeholder-glow"}>
        <table className={"w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400"}>
            {getTableHeader()}
            <tbody>
            {tableEntries}
            </tbody>
        </table>
        {getTableControls()}
    </p>
}

export default UserManagement
