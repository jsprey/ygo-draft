import React, {useState} from "react";
import {Draft} from "../../api/Draft";
import DraftSettingsDetails from "./DraftSettingsDetails";
import classNames from "classnames";
import DraftBattleBadge from "./DraftBattleBadge";
import {User} from "../../api/hooks/users/useUsers";
import {Friend} from "../../api/UserModel";
import YgoIcon from "../../core/YgoIcon";

const CollapsedIcon = <YgoIcon icon={"collapse"} size={18}
                               classNames={"fill-gray-700 hover:fill-gray-600 active:fill-gray-500 dark:fill-gray-400 hover:dark:fill-gray-300 active:dark:fill-gray-200"}/>
const ExpandIcon = <YgoIcon icon={"expand"} size={18}
                            classNames={"fill-gray-700 hover:fill-gray-600 active:fill-gray-500 dark:fill-gray-400 hover:dark:fill-gray-300 active:dark:fill-gray-200"}/>

type DraftHeaderProps = {
    draft: Draft;
    player: User,
    enemy: Friend
};

function DraftHeader(props: DraftHeaderProps) {
    const [settingCollapsed, setSettingCollapsed] = useState<boolean>(true)

    const settingsHeaderCN = classNames("flex-grow-1 pl-2", "rounded-tl rounded-tr", settingCollapsed ? "rounded-bl rounded-br" : "", "border dark:border-white", "bg-gray-200 dark:bg-gray-600")
    const settingsBodyCN = classNames("p-2", "rounded-bl rounded-br", "border-start border-bottom border-end dark:border-white", "bg-gray-200 dark:bg-gray-600")

    return <div className={"w-100 text-gray-400 dark:text-gray-100"}>
        <DraftBattleBadge playerNameOne={props.player.display_name} playerNameTwo={props.enemy.name}/>

        <div className={classNames("flex justify-content-center", "mt-2", settingsHeaderCN)}>
            <div
                className={classNames("align-self-center mr-1 uppercase fw-bold")}>
                Draft Settings
            </div>
            <div className={"align-self-center"} onClick={() => setSettingCollapsed(!settingCollapsed)}>
                {settingCollapsed ? CollapsedIcon : ExpandIcon}
            </div>
        </div>
        {settingCollapsed ? <></> :
            <DraftSettingsDetails settings={props.draft.settings} containerClasses={settingsBodyCN}/>}
    </div>
}

export default DraftHeader
