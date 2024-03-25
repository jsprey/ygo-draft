import React, {useState} from "react";
import classNames from "classnames";
import {User, useUsers} from "../api/hooks/useUsers";
import {Alert} from "react-bootstrap";

function UserManagement() {
    const [currentPage, setCurrentPage] = useState<number>(0);
    const [currentPageSize, setCurrentPageSize] = useState<number>(15);
    const {data, isLoading, error} = useUsers(currentPage, currentPageSize);

    let tableEntries: JSX.Element[] = []
    if (error) {
        tableEntries.push(<tr><td colSpan={5}><Alert variant={"danger"}>Failed to load users!</Alert></td></tr>)
    } else if (isLoading) {
        tableEntries = createPlaceholderEntries()
    }else if (data) {
        tableEntries = createTableWithActualData(data.users)
    }

    function createTableWithActualData(users: User[]): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        users.forEach(user => {
            const entry = <tr key={`row-${user.id}`}>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1 p-2 text-sm")}>{user.id}</td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1 p-2 text-sm")}>{user.email}</td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1 p-2 text-sm")}>{user.display_name}</td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1 p-2 text-sm")}>{user.is_admin ? "Yes" : "No"}</td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1 p-2 text-sm")}></td>
            </tr>
            tableEntries.push(entry)
        })

        return tableEntries
    }

    function createPlaceholderEntries(): JSX.Element[] {
        const tableEntries: JSX.Element[] = []

        for (let i = 0; i < 15; i++) {
            const entry = <tr key={`row-placeholder-${i}`}>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1")}><span
                    className={"placeholder w-100"}>1</span></td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1")}><span
                    className={"placeholder w-100"}>1</span></td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1")}><span
                    className={"placeholder w-100"}>1</span></td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1")}><span
                    className={"placeholder w-100"}>1</span></td>
                <td className={classNames("bg-ygo-table dark:bg-ygo-table-dark", "border-1")}><span
                    className={"placeholder w-100"}>1</span></td>
            </tr>
            tableEntries.push(entry)
        }

        return tableEntries
    }

    function getTableHeader() {
        let allHeaderCN = classNames("border border-2", "p-2", "text-uppercase font-monospace fw-bold text-xs", "bg-gray-200 dark:bg-gray-600 text-ygo-table-header-text dark:text-ygo-table-header-text-dark")
        let headerTopLeftCN = classNames(allHeaderCN, "rounded-tl")
        let headerTopRightCN = classNames(allHeaderCN, "rounded-tr-1")
        let headerInnerCN = classNames(allHeaderCN, "")

        return <thead>
        <tr>
            <th className={headerTopLeftCN}>ID</th>
            <th className={headerInnerCN}>Email</th>
            <th className={headerInnerCN}>Display Name</th>
            <th className={headerInnerCN}>Administrator</th>
            <th className={headerTopRightCN}>Actions</th>
        </tr>
        </thead>;
    }


    function getTableControls() {
        return <div className={"flex justify-content-end bg-ygo-table dark:bg-ygo-table-dark border-1 "}>
            <div className={""}>
                asd
            </div>
        </div>
    }

    return <p className={"placeholder-glow"}>
        <table className="shadow w-full table-auto table">
            {getTableHeader()}
            <tbody>
            {tableEntries}
            </tbody>
        </table>
        {getTableControls()}
    </p>
}

export default UserManagement
