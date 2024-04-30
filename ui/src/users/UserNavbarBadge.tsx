import {useCurrentUser} from "../api/hooks/users/useUser";
import {Link} from "react-router-dom";
import Spinner from "../core/Spinner";
import classNames from "classnames";
import {UserPath} from "../routes/AppRouter";
import YgoIcon from "../core/YgoIcon";

export type UserNavbarBadgeProps = {
    className?: string
    focused: boolean
}

function UserNavbarBadge(props: UserNavbarBadgeProps) {
    const {data, isLoading, error} = useCurrentUser()

    let iconCN = classNames("fill-dark dark:fill-light")
    if (props.focused) {
        iconCN = classNames("fill-light")
    }
    const userIcon = <YgoIcon icon={"user"} size={18} classNames={iconCN}/>

    let content = <></>
    if (isLoading) {
        content = <div className={"flex align-content-center"}>
            {userIcon}
            <Spinner/>
        </div>
    } else if (error) {
        content = <div className={"flex align-content-center"}>
            {userIcon}
            <div className={"bg-danger text-white pl-1 pr-1"}>Failed to get user!</div>
        </div>
    } else if (data) {
        content = <div className={"flex justify-content-center"}>
            {userIcon}
            <span className={"ml-2"}>
                {data.display_name}
            </span>
        </div>
    }

    return <Link to={UserPath} className={classNames(props.className ? props.className : "")}>
            {content}
        </Link>
}

export default UserNavbarBadge